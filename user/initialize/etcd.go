package initialize

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
	"user/global"

	clientv3 "go.etcd.io/etcd/client/v3"
)

type EtcdServer struct {
	Endpoints   []string
	DialTimeout int32

	etcdTTL     int64
	cli         *clientv3.Client
	leaseID     clientv3.LeaseID
	keepAliveCh <-chan *clientv3.LeaseKeepAliveResponse
	closeCh     <-chan struct{}

	serverInfo global.ServerInfo
}

func NewEtcdServer() *EtcdServer {
	return &EtcdServer{
		Endpoints:   []string{fmt.Sprintf("%s:%d", global.EtcdConfig.Host, global.EtcdConfig.Port)},
		DialTimeout: 5,
		serverInfo:  global.Server,
	}
}

func (etcd *EtcdServer) Register(ttl int64) {
	var err error
	if etcd.cli, err = clientv3.New(clientv3.Config{
		Endpoints:   etcd.Endpoints,
		DialTimeout: time.Duration(etcd.DialTimeout) * time.Second,
	}); err != nil {
		fmt.Println("etcd 初始化失败: ", err)
		return
	}
	etcd.etcdTTL = ttl
	if err := etcd.register(); err != nil {
		fmt.Println("etcd 注册失败: ", err)
		return
	}
	etcd.closeCh = make(chan struct{})
	go etcd.keepAlive()
}

func (etcd *EtcdServer) register() error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(etcd.etcdTTL)*time.Second)
	defer cancel()
	lease, err := etcd.cli.Grant(ctx, etcd.etcdTTL)
	if err != nil {
		return err
	}
	etcd.leaseID = lease.ID

	if etcd.keepAliveCh, err = etcd.cli.KeepAlive(ctx, etcd.leaseID); err != nil {
		return err
	}

	serverInfo, err := json.Marshal(etcd.serverInfo)
	if err != nil {
		return err
	}
	_, err = etcd.cli.Put(ctx, global.EtcdServerKey+etcd.serverInfo.Name, string(serverInfo), clientv3.WithLease(etcd.leaseID))
	return err
}

func (etcd *EtcdServer) keepAlive() error {
	ticker := time.NewTicker(time.Duration(etcd.etcdTTL) * time.Second)
	for {
		select {
		case <-etcd.closeCh:
			if err := etcd.unregister(); err != nil {
				fmt.Println("etcd 注销失败: ", err)
			}
		case res := <-etcd.keepAliveCh:
			if res == nil {
				if err := etcd.register(); err != nil {
					fmt.Println("etcd 保活失败: ", err)
				}
			}
		case <-ticker.C:
			if etcd.keepAliveCh == nil {
				if err := etcd.register(); err != nil {
					fmt.Println("etcd 保活失败: ", err)
				}
			}
		}
	}
}

func (etcd *EtcdServer) unregister() error {
	if _, err := etcd.cli.Delete(context.Background(), global.EtcdServerKey+etcd.serverInfo.Name); err != nil {
		return err
	}
	if _, err := etcd.cli.Revoke(context.Background(), etcd.leaseID); err != nil {
		return err
	}
	return nil
}

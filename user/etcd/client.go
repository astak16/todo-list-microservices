package etcd

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc/resolver"
)

var (
	// ErrNoKey indicates a client method needs a key but receives none.
	ErrNoKey = errors.New("no key provided")

	// ErrNoValue indicates a client method needs a value but receives none.
	ErrNoValue = errors.New("no value provided")
)

type Service struct {
	Key   string
	Value string
}

type client struct {
	ctx        context.Context
	cli        *clientv3.Client
	kv         clientv3.KV
	leaseID    clientv3.LeaseID
	leaser     clientv3.Lease
	watcher    clientv3.Watcher
	hbch       <-chan *clientv3.LeaseKeepAliveResponse
	clientConn resolver.ClientConn
}

type ClientOptions struct {
	DialTimeout time.Duration
}

func NewClient(ctx context.Context, machine []string, options ClientOptions) (client, error) {
	if options.DialTimeout == 0 {
		options.DialTimeout = 3 * time.Second
	}

	cli, err := clientv3.New(clientv3.Config{
		Context:     ctx,
		Endpoints:   machine,
		DialTimeout: options.DialTimeout,
	})

	if err != nil {
		return client{}, err
	}
	return client{
		cli: cli,
		ctx: ctx,
		kv:  clientv3.NewKV(cli),
	}, nil
}

func (c *client) Register(s Service) error {
	if s.Key == "" {
		return ErrNoKey
	}

	if s.Value == "" {
		return ErrNoValue
	}

	if c.leaser != nil {
		c.leaser.Close()
	}
	c.leaser = clientv3.NewLease(c.cli)

	if c.watcher != nil {
		c.watcher.Close()
	}
	c.watcher = clientv3.NewWatcher(c.cli)

	grantResp, err := c.leaser.Grant(c.ctx, 10)
	if err != nil {
		return err
	}

	c.leaseID = grantResp.ID
	_, err = c.kv.Put(c.ctx, s.Key, s.Value, clientv3.WithLease(c.leaseID))
	if err != nil {
		return err
	}

	c.hbch, err = c.leaser.KeepAlive(c.ctx, c.leaseID)
	if err != nil {
		return err
	}

	go func() {
		for {
			select {
			case r := <-c.hbch:
				if r != nil {
					return
				}
			case <-c.ctx.Done():
				return
			}
		}
	}()
	return nil
}

func (c *client) Deregister(s Service) error {
	defer c.close()
	if s.Key == "" {
		return ErrNoKey
	}
	if _, err := c.kv.Delete(c.ctx, s.Key, clientv3.WithIgnoreLease()); err != nil {
		return err
	}
	return nil
}

func (c *client) Watch(cc resolver.ClientConn, s Service) {
	if s.Key == "" {
		fmt.Println(ErrNoKey)
		return
	}

	c.clientConn = cc
	addrM := make(map[string]resolver.Address)
	state := resolver.State{}

	update := func() {
		addrList := make([]resolver.Address, 0, len(addrM))
		for _, address := range addrM {
			addrList = append(addrList, address)
		}
		state.Addresses = addrList
		err := c.clientConn.UpdateState(state)
		if err != nil {
			fmt.Println("更新地址出错：", err)
		}
	}

	resp, err := c.kv.Get(c.ctx, s.Key, clientv3.WithPrefix())
	if err != nil {
		fmt.Printf("获取地址出错：%v\n", err)
		return
	} else {
		for _, kv := range resp.Kvs {
			addr, err := splitPath(string(kv.Key))
			if err != nil {
				fmt.Printf("解析key报错：%v\n", err)
				return
			}
			addrM[string(kv.Value)] = resolver.Address{
				Addr:       addr,
				ServerName: s.Key,
			}
		}
	}

	update()

	dch := c.cli.Watch(context.Background(), s.Key, clientv3.WithPrefix(), clientv3.WithPrevKV())
	for response := range dch {
		for _, event := range response.Events {
			switch event.Type {
			case clientv3.EventTypePut:
				addr, err := splitPath(string(event.Kv.Key))
				if err != nil {
					fmt.Printf("解析key报错：%v\n", err)
					return
				}
				addrM[string(event.Kv.Value)] = resolver.Address{
					Addr:       addr,
					ServerName: s.Key,
				}
			case clientv3.EventTypeDelete:
				delete(addrM, string(event.Kv.Key))
			}
		}
		update()
	}

}

func (c *client) close() {
	if c.leaser != nil {
		c.leaser.Close()
	}
	if c.watcher != nil {
		c.watcher.Close()
	}
}

func splitPath(path string) (string, error) {
	strs := strings.Split(path, "/")
	if len(strs) == 0 {
		return "", errors.New("invalid path")
	}
	return strs[len(strs)-1], nil
}

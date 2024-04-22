package etcd

import (
	"google.golang.org/grpc/resolver"
)

type Discovery struct {
	cli     client
	service Service
}

func NewDiscovery(cli client, service Service) *Discovery {
	return &Discovery{
		cli:     cli,
		service: service,
	}
}

func (d *Discovery) Scheme() string {
	return "etcd"
}

func (d *Discovery) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
	go d.cli.Watch(cc, d.service)
	return d, nil
}

func (d *Discovery) ResolveNow(rn resolver.ResolveNowOptions) {}

func (d *Discovery) Close() {}

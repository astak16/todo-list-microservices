package etcd

import "fmt"

type Register struct {
	cli     client
	service Service
}

func NewRegistrar(cli client, service Service) *Register {
	return &Register{
		cli:     cli,
		service: service,
	}
}

func (r *Register) Register() {
	if err := r.cli.Register(r.service); err != nil {
		fmt.Println(err)
		return
	}
}

func (r *Register) Deregister() {
	if err := r.cli.Deregister(r.service); err != nil {
		fmt.Println(err)
	}
}

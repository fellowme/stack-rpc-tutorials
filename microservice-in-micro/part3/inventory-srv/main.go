package main

import (
	"fmt"

	"github.com/micro-in-cn/tutorials/microservice-in-micro/part3/basic"
	"github.com/micro-in-cn/tutorials/microservice-in-micro/part3/basic/config"
	"github.com/micro-in-cn/tutorials/microservice-in-micro/part3/inventory-srv/handler"
	"github.com/micro-in-cn/tutorials/microservice-in-micro/part3/inventory-srv/model"
	proto "github.com/micro-in-cn/tutorials/microservice-in-micro/part3/inventory-srv/proto/inventory"
	"github.com/micro/cli/v2"
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
)

func main() {
	// 初始化配置、数据库等信息
	basic.Init()

	// 使用etcd注册
	micReg := etcd.NewRegistry(registryOptions)

	// 新建服务
	service := micro.NewService(
		micro.Name("mu.micro.book.srv.inventory"),
		micro.Registry(micReg),
		micro.Version("latest"),
	)

	// 服务初始化
	service.Init(
		micro.Action(func(c *cli.Context) error {
			// 初始化模型层
			model.Init()
			// 初始化handler
			handler.Init()

			return nil
		}),
	)

	// 注册服务
	proto.RegisterInventoryHandler(service.Server(), new(handler.Service))

	// 启动服务
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

func registryOptions(ops *registry.Options) {
	etcdCfg := config.GetEtcdConfig()
	ops.Addrs = []string{fmt.Sprintf("%s:%d", etcdCfg.GetHost(), etcdCfg.GetPort())}
}

package main

import (
	"context"
	"fmt"
	"github.com/afex/hystrix-go/hystrix"
	consul2 "github.com/go-micro/plugins/v4/registry/consul"
	"github.com/go-micro/plugins/v4/wrapper/select/roundrobin"
	opentracingFn "github.com/go-micro/plugins/v4/wrapper/trace/opentracing"
	"github.com/opentracing/opentracing-go"
	service_cart "github.com/xiaohu913255/MicroTest/cart/proto/cart"
	"github.com/xiaohu913255/MicroTest/cartApi/common"
	"github.com/xiaohu913255/MicroTest/cartApi/handler"
	"go-micro.dev/v4"
	"go-micro.dev/v4/client"
	log "go-micro.dev/v4/logger"
	"go-micro.dev/v4/registry"
	"net"
	"net/http"

	cartApi "github.com/xiaohu913255/MicroTest/cartApi/proto/cartApi"
)

func main() {
	//注册中心
	consul := consul2.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			"127.0.0.1:8500",
		}
	})

	//链路追踪
	t, io, err := common.NewTracer("go.micro.api.cartApi", "localhost:6831")
	if err != nil {
		log.Error(err)
	}
	defer io.Close()
	opentracing.SetGlobalTracer(t)

	//熔断器
	hystrixStreamHandler := hystrix.NewStreamHandler()
	hystrixStreamHandler.Start()
	//启动端口
	go func() {
		err = http.ListenAndServe(net.JoinHostPort("0.0.0.0", " "), hystrixStreamHandler)
		if err != nil {
			log.Error(err)
		}
	}()

	// New Service
	service := micro.NewService(
		micro.Name("go.micro.api.cartApi"),
		micro.Version("latest"),
		micro.Address("0.0.0.0:8086"),
		//添加 consul 注册中心
		micro.Registry(consul),
		//添加链路追踪
		micro.WrapClient(opentracingFn.NewClientWrapper(opentracing.GlobalTracer())),
		//添加熔断
		micro.WrapClient(NewClientHystrixWrapper()),
		//添加负载均衡
		micro.WrapClient(roundrobin.NewClientWrapper()),
	)

	// Initialise service
	service.Init()

	cartService := service_cart.NewCartService("go.micro.service.cart", service.Client())

	cartService.AddCart(context.TODO(), &service_cart.CartInfo{

		UserId:    3,
		ProductId: 4,
		SizeId:    5,
		Num:       5,
	})

	// Register Handler
	if err := cartApi.RegisterCartApiHandler(service.Server(), &handler.CartApi{CartService: cartService}); err != nil {
		log.Error(err)
	}

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

type clientWrapper struct {
	client.Client
}

func (c *clientWrapper) Call(ctx context.Context, req client.Request, rsp interface{}, opts ...client.CallOption) error {
	return hystrix.Do(req.Service()+"."+req.Endpoint(), func() error {
		//run 正常执行
		fmt.Println(req.Service() + "." + req.Endpoint())
		return c.Client.Call(ctx, req, rsp, opts...)
	}, func(err error) error {
		fmt.Println(err)
		return err
	})
}

// 初始化Wapper
func NewClientHystrixWrapper() client.Wrapper {
	return func(i client.Client) client.Client {
		return &clientWrapper{i}
	}
}

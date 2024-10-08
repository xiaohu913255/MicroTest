package main

import (
	"github.com/go-micro/plugins/v4/registry/consul"
	"github.com/go-micro/plugins/v4/wrapper/monitoring/prometheus"
	ratelimit "github.com/go-micro/plugins/v4/wrapper/ratelimiter/uber"
	opentracingFn "github.com/go-micro/plugins/v4/wrapper/trace/opentracing"
	"github.com/jinzhu/gorm"
	"github.com/opentracing/opentracing-go"
	"github.com/xiaohu913255/MicroTest/payment/common"
	"github.com/xiaohu913255/MicroTest/payment/domain/repository"
	service2 "github.com/xiaohu913255/MicroTest/payment/domain/service"
	"github.com/xiaohu913255/MicroTest/payment/handler"
	"go-micro.dev/v4"
	log "go-micro.dev/v4/logger"
	"go-micro.dev/v4/registry"

	_ "github.com/jinzhu/gorm/dialects/mysql"

	payment "github.com/xiaohu913255/MicroTest/payment/proto/payment"
)

func main() {
	// 配置中心
	consulConfig, err := common.GetConsulConfig("localhost", 8500, "/micro/config")
	if err != nil {
		common.Error(err)
	}
	//注册中心
	consul := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			"localhost:8500",
		}
	})
	//jaeger 链路追踪
	t, io, err := common.NewTracer("go.micro.service.payment", "localhost:6831")
	if err != nil {
		common.Error(err)
	}
	defer io.Close()
	opentracing.SetGlobalTracer(t)

	//mysql 设置
	mysqlInfo := common.GetMysqlFromConsul(consulConfig, "mysql")
	//初始化数据库
	db, err := gorm.Open("mysql", mysqlInfo.User+":"+mysqlInfo.Pwd+"@/"+mysqlInfo.Database+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		common.Error(err)
	}
	defer db.Close()
	//禁止复数表
	db.SingularTable(true)

	//创建表
	tableInit := repository.NewPaymentRepository(db)
	tableInit.InitTable()

	//监控
	common.PrometheusBoot(9089)

	// New Service
	service := micro.NewService(
		micro.Name("go.micro.service.payment"),
		micro.Version("latest"),
		micro.Address("0.0.0.0:8089"),
		//添加注册中心
		micro.Registry(consul),
		//添加链路追踪
		micro.WrapHandler(opentracingFn.NewHandlerWrapper(opentracing.GlobalTracer())),
		//加载限流
		micro.WrapHandler(ratelimit.NewHandlerWrapper(1000)),
		//加载监控
		micro.WrapHandler(prometheus.NewHandlerWrapper()),
	)

	// Initialise service
	service.Init()

	paymentDataService := service2.NewPaymentDataService(repository.NewPaymentRepository(db))

	// Register Handler
	payment.RegisterPaymentHandler(service.Server(), &handler.Payment{PaymentDataService: paymentDataService})

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"github.com/go-micro/plugins/v4/registry/consul"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/xiaohu913255/MicroTest/category/domain/repository"
	service2 "github.com/xiaohu913255/MicroTest/category/domain/service"
	"github.com/xiaohu913255/MicroTest/category/handler"
	"github.com/xiaohu913255/MicroTest/common"
	"go-micro.dev/v4"
	log "go-micro.dev/v4/logger"
	"go-micro.dev/v4/registry"

	category "github.com/xiaohu913255/MicroTest/category/proto/category"
)

func main() {
	//配置中心
	consulConfig, err := common.GetConsulConfig("127.0.0.1", 8500, "/micro/config")
	if err != nil {
		log.Error(err)
	}
	//注册中心
	consulRegistry := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			"127.0.0.1:8500",
		}
	})

	// New Service
	service := micro.NewService(
		micro.Name("go.micro.service.category"),
		micro.Version("latest"),
		//这里设置地址和需要暴露的端口
		micro.Address("127.0.0.1:8082"),
		//添加consul 作为注册中心
		micro.Registry(consulRegistry),
	)

	//获取mysql配置,路径中不带前缀
	mysqlInfo := common.GetMysqlFromConsul(consulConfig, "mysql")

	//连接数据库
	db, err := gorm.Open("mysql", mysqlInfo.User+":"+mysqlInfo.Pwd+"@/"+mysqlInfo.Database+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Error(err)
	}
	defer db.Close()
	//禁止复表
	db.SingularTable(true)

	//rp := repository.NewCategoryRepository(db)
	//rp.InitTable()
	// Initialise service
	service.Init()

	categoryDataService := service2.NewCategoryDataService(repository.NewCategoryRepository(db))

	err = category.RegisterCategoryHandler(service.Server(), &handler.Category{CategoryDataService: categoryDataService})
	if err != nil {
		log.Error(err)
	}

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

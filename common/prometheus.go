package common

import (
	"github.com/prometheus/client_golang/prometheus/promhttp"
	//"github.com/prometheus/common/log"
	_ "github.com/prometheus/common/promlog"
	"net/http"
	"strconv"
	"testing"
)

var t *testing.T

func PrometheusBoot(port int) {
	http.Handle("/metrics", promhttp.Handler())
	//启动web 服务
	go func() {
		err := http.ListenAndServe("0.0.0.0:"+strconv.Itoa(port), nil)
		if err != nil {
			t.Fatal("启动失败")
		}
		//t.Info("监控启动,端口为：" + strconv.Itoa(port))
	}()
}

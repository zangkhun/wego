package main

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	. "wego/app/conf"
	"wego/library"
)

// ServerConfs是全局配置数据
var ServConfs = func() ServerConf {
	var s = new(ServerConf)
	var err = library.ReadYamlConfig(s, "./app/conf/settings.yaml")
	if err != nil {
		panic(errors.New("loading yaml file err, " + err.Error()))
	}
	return *s
}()

func main() {
	//gin.ForceConsoleColor() // web服务控制台日志颜色
	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		log.Printf("API %-6s %-25s --> %s (%d handlers)\n", httpMethod, absolutePath, handlerName, nuHandlers)
	}

	fmt.Println(ServConfs)
	var conf = library.ServiceConfig{
		IPAddr: ServConfs.SysConf.IP + ":" + ServConfs.SysConf.Port,
		Debug:  ServConfs.SysConf.Debug,
	}

	server := library.NewService(conf)
	server.UseMidwares(library.CORSMiddleware())
	server.RegisterRouters(Routers)

	server.Run()
}

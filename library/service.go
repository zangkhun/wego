package library

import "github.com/gin-gonic/gin"

// ServiceConfig 是web服务的配置信息
type ServiceConfig struct {
	IPAddr string // ip:port
	Debug  bool   // 是否开启Debug
}

// Service 是web服务引擎
type Service struct {
	conf   ServiceConfig //服务配置
	engine *gin.Engine   // 服务引擎
}

// NewService 创建一个服务
func NewService(conf ServiceConfig) *Service {
	return &Service{
		conf:   ServiceConfig{},
		engine: gin.Default(),
	}
}

// 启动服务
func (s *Service) Run() error {
	return s.engine.Run(s.conf.IPAddr)
}

// RegisterMidwares 注册 Routers
func (s *Service) UseMidwares(midwares ...gin.HandlerFunc) {
	s.engine.Use(midwares...)
}

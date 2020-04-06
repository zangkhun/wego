package library

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"strings"
)

// Router 是一个服务的路由属性
type Router struct {
	Method 		string
	URI 		string
	Handler 	gin.HandlerFunc
}

// RouterGroup 是一组路由
type RouterGroup struct {
	Group 		map[string][]Router
}

// NewRouter 创建一个 Router
func NewRouter(m string, u string, h gin.HandlerFunc) Router {
	return Router{
		Method:  m,
		URI:     u,
		Handler: h,
	}
}

// registerOne 向服务注册路由(uri 和处理函数)
func (s *Service) registerOne(method string, uri string, handler gin.HandlerFunc) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Fprintf(os.Stderr, "register error, %s, %s \n", method, uri)
		}
	}()
	s.engine.Handle(strings.ToUpper(method), uri, handler)
}

// registerOneToGroup 向服务注册分组路由(uri 和处理函数)
func (s *Service) registerOneToGroup(group *gin.RouterGroup, method string, uri string, handler gin.HandlerFunc) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Fprintf(os.Stderr, "register error, %s, %s \n", method, uri)
		}
	}()
	group.Handle(strings.ToUpper(method), uri, handler)
}

// RegisterRouters 注册路由列表
func (s *Service) RegisterRouters(routers []Router) {
	for _, router := range routers {
		s.registerOne(router.Method, router.URI, router.Handler)
	}
}

// RegisterRouterGroups 注册一组路由，这里一组路由是有相同中间件或URI组
func (s *Service) RegisterRouterGroups(groups map[string][]Router) {
	for k, v := range groups {
		g := s.engine.Group(k)
		for _, router := range v {
			s.registerOneToGroup(g, router.Method, router.URI, router.Handler)
		}
	}
}

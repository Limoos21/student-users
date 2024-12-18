package router

import (
	"github.com/gin-gonic/gin"
)

type RouterGroup struct {
	Prefix      string
	Middlewares []gin.HandlerFunc
	Engine      *gin.Engine
}

func NewRouterGroup(prefix string, Middlewares []gin.HandlerFunc, Engine *gin.Engine) *gin.RouterGroup {
	router := RouterGroup{prefix, Middlewares, Engine}
	return Engine.Group(router.Prefix, router.Middlewares...)
}

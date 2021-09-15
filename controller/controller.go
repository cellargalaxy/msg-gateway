package controller

import (
	"fmt"
	"github.com/cellargalaxy/go_common/util"
	"github.com/cellargalaxy/msg-gateway/model"
	"github.com/cellargalaxy/msg-gateway/static"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func Controller() error {
	engine := gin.Default()
	engine.Use(util.GinLogId)
	engine.Use(util.GinLog)
	engine.GET("/ping", util.Ping)
	engine.POST("/ping", validate, util.Ping)

	engine.Use(staticCache)
	engine.StaticFS("/static", http.FS(static.StaticFile))

	engine.GET("/api/listAllTemplate", validate, listAllTemplate)
	engine.POST("/api/sendTemplateToTag", validate, sendTemplateToTag)
	engine.POST("/api/sendTemplateToCommonTag", validate, sendTemplateToCommonTag)

	engine.POST("/api/createTag", validate, createTag)
	engine.POST("/api/deleteTag", validate, deleteTag)
	engine.GET("/api/listAllTag", validate, listAllTag)
	engine.POST("/api/addTagToUser", validate, addTagToUser)
	engine.POST("/api/deleteTagFromUser", validate, deleteTagFromUser)

	engine.GET("/api/listAllUserInfo", validate, listAllUserInfo)

	engine.POST("/api/sendTgMsg2ConfigChatId", validate, sendTgMsg2ConfigChatId)

	err := engine.Run(model.ListenAddress)
	if err != nil {
		panic(fmt.Errorf("web服务启动，异常: %+v", err))
	}
	return nil
}

func staticCache(c *gin.Context) {
	if strings.HasPrefix(c.Request.RequestURI, "/static") {
		c.Header("Cache-Control", "max-age=86400")
	}
}

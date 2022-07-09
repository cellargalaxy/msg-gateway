package controller

import (
	"fmt"
	"github.com/cellargalaxy/go_common/util"
	"github.com/cellargalaxy/msg_gateway/model"
	"github.com/cellargalaxy/msg_gateway/static"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func Controller() error {
	engine := gin.Default()
	engine.Use(claims)
	engine.Use(util.GinLog)
	engine.GET(util.PingPath, util.Ping)
	engine.POST(util.PingPath, validate, util.Ping)

	engine.Use(staticCache)
	engine.StaticFS("/static", http.FS(static.StaticFile))

	engine.GET(model.ListAllTemplatePath, validate, listAllTemplate)
	engine.POST(model.SendTemplateToTagPath, validate, sendTemplateToTag)
	engine.POST(model.SendTemplateToCommonTagPath, validate, sendTemplateToCommonTag)

	engine.POST(model.CreateTagPath, validate, createTag)
	engine.POST(model.DeleteTagPath, validate, deleteTag)
	engine.GET(model.ListAllTagPath, validate, listAllTag)
	engine.POST(model.AddTagToUserPath, validate, addTagToUser)
	engine.POST(model.DeleteTagFromUserPath, validate, deleteTagFromUser)

	engine.GET(model.ListAllUserInfoPath, validate, listAllUserInfo)

	engine.POST(model.SendTgMsg2ConfigChatIdPath, validate, sendTgMsg2ConfigChatId)

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

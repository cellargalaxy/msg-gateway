package controller

import (
	"fmt"
	common_model "github.com/cellargalaxy/go_common/model"
	"github.com/cellargalaxy/go_common/util"
	"github.com/cellargalaxy/msg_gateway/model"
	"github.com/cellargalaxy/msg_gateway/static"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func Controller() error {
	engine := gin.Default()
	engine.Use(claims)
	engine.Use(util.GinLog)

	debug := engine.Group(common_model.DebugPath, validate)
	pprof.RouteRegister(debug, common_model.PprofPath)

	engine.GET(common_model.PingPath, util.Ping)
	engine.POST(common_model.PingPath, validate, util.Ping)

	engine.Use(staticCache)
	engine.StaticFS(common_model.StaticPath, http.FS(static.StaticFile))

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
	if strings.HasPrefix(c.Request.RequestURI, common_model.StaticPath) {
		c.Header("Cache-Control", "max-age=86400")
	}
}

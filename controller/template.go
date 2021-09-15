package controller

import (
	"github.com/cellargalaxy/go_common/util"
	"github.com/cellargalaxy/msg-gateway/model"
	"github.com/cellargalaxy/msg-gateway/service/controller"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

//获取全部模板
func listAllTemplate(ctx *gin.Context) {
	var request model.ListAllTemplateRequest
	err := ctx.BindQuery(&request)
	if err != nil {
		logrus.WithContext(ctx).WithFields(logrus.Fields{"request": request, "err": err}).Error("获取全部模板，请求参数解析异常")
		ctx.JSON(http.StatusOK, util.CreateErrResponse(err.Error()))
		return
	}
	logrus.WithContext(ctx).WithFields(logrus.Fields{"request": request}).Info("获取全部模板")
	ctx.JSON(http.StatusOK, util.CreateResponse(controller.ListAllTemplate(ctx, request)))
}

//给标签用户发送模板消息
func sendTemplateToTag(ctx *gin.Context) {
	var request model.SendTemplateToTagRequest
	err := ctx.BindJSON(&request)
	if err != nil {
		logrus.WithContext(ctx).WithFields(logrus.Fields{"request": request, "err": err}).Error("给标签用户发送模板消息，请求参数解析异常")
		ctx.JSON(http.StatusOK, util.CreateErrResponse(err.Error()))
		return
	}
	logrus.WithContext(ctx).WithFields(logrus.Fields{"request": request}).Info("给标签用户发送模板消息")
	ctx.JSON(http.StatusOK, util.CreateResponse(controller.SendTemplateToTag(ctx, request)))
}

//给通用标签用户发送模板消息
func sendTemplateToCommonTag(ctx *gin.Context) {
	var request model.SendTemplateToCommonTagRequest
	err := ctx.BindJSON(&request)
	if err != nil {
		logrus.WithContext(ctx).WithFields(logrus.Fields{"request": request, "err": err}).Error("给通用标签用户发送模板消息，请求参数解析异常")
		ctx.JSON(http.StatusOK, util.CreateErrResponse(err.Error()))
		return
	}
	logrus.WithContext(ctx).WithFields(logrus.Fields{"request": request}).Info("给通用标签用户发送模板消息")
	ctx.JSON(http.StatusOK, util.CreateResponse(controller.SendTemplateToCommonTag(ctx, request)))
}

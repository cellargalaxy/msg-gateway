package controller

import (
	"github.com/cellargalaxy/go_common/util"
	"github.com/cellargalaxy/msg-gateway/model"
	"github.com/cellargalaxy/msg-gateway/service/controller"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

//创建标签
func createTag(ctx *gin.Context) {
	var request model.CreateTagRequest
	err := ctx.BindJSON(&request)
	if err != nil {
		logrus.WithContext(ctx).WithFields(logrus.Fields{"request": request, "err": err}).Error("创建标签，请求参数解析异常")
		ctx.JSON(http.StatusOK, util.CreateErrResponse(err.Error()))
		return
	}
	logrus.WithContext(ctx).WithFields(logrus.Fields{"request": request}).Info("创建标签")
	ctx.JSON(http.StatusOK, util.CreateResponse(controller.CreateTag(ctx, request)))
}

//删除标签
func deleteTag(ctx *gin.Context) {
	var request model.DeleteTagRequest
	err := ctx.BindJSON(&request)
	if err != nil {
		logrus.WithContext(ctx).WithFields(logrus.Fields{"request": request, "err": err}).Error("删除标签，请求参数解析异常")
		ctx.JSON(http.StatusOK, util.CreateErrResponse(err.Error()))
		return
	}
	logrus.WithContext(ctx).WithFields(logrus.Fields{"request": request}).Info("删除标签")
	ctx.JSON(http.StatusOK, util.CreateResponse(controller.DeleteTag(ctx, request)))
}

//获取所有标签
func listAllTag(ctx *gin.Context) {
	var request model.ListAllTagRequest
	err := ctx.BindQuery(&request)
	if err != nil {
		logrus.WithContext(ctx).WithFields(logrus.Fields{"request": request, "err": err}).Error("获取所有标签，请求参数解析异常")
		ctx.JSON(http.StatusOK, util.CreateErrResponse(err.Error()))
		return
	}
	logrus.WithContext(ctx).WithFields(logrus.Fields{"request": request}).Info("获取所有标签")
	ctx.JSON(http.StatusOK, util.CreateResponse(controller.ListAllTag(ctx, request)))
}

//为用户加标签
func addTagToUser(ctx *gin.Context) {
	var request model.AddTagToUserRequest
	err := ctx.BindJSON(&request)
	if err != nil {
		logrus.WithContext(ctx).WithFields(logrus.Fields{"request": request, "err": err}).Error("为用户加标签，请求参数解析异常")
		ctx.JSON(http.StatusOK, util.CreateErrResponse(err.Error()))
		return
	}
	logrus.WithContext(ctx).WithFields(logrus.Fields{"request": request}).Info("为用户加标签")
	ctx.JSON(http.StatusOK, util.CreateResponse(controller.AddTagToUser(ctx, request)))
}

//为用户删标签
func deleteTagFromUser(ctx *gin.Context) {
	var request model.DeleteTagFromUserRequest
	err := ctx.BindJSON(&request)
	if err != nil {
		logrus.WithContext(ctx).WithFields(logrus.Fields{"request": request, "err": err}).Error("为用户删标签，请求参数解析异常")
		ctx.JSON(http.StatusOK, util.CreateErrResponse(err.Error()))
		return
	}
	logrus.WithContext(ctx).WithFields(logrus.Fields{"request": request}).Info("为用户删标签")
	ctx.JSON(http.StatusOK, util.CreateResponse(controller.DeleteTagFromUser(ctx, request)))
}

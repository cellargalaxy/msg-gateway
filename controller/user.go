package controller

import (
	"github.com/cellargalaxy/go_common/util"
	"github.com/cellargalaxy/msg-gateway/model"
	"github.com/cellargalaxy/msg-gateway/service/controller"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

//获取全部用户信息
func listAllUserInfo(ctx *gin.Context) {
	var request model.ListAllUserInfoRequest
	err := ctx.BindQuery(&request)
	if err != nil {
		logrus.WithContext(ctx).WithFields(logrus.Fields{"request": request, "err": err}).Error("获取全部用户信息，请求参数解析异常")
		ctx.JSON(http.StatusOK, util.CreateErrResponse(err.Error()))
		return
	}
	logrus.WithContext(ctx).WithFields(logrus.Fields{"request": request}).Info("获取全部用户信息")
	ctx.JSON(http.StatusOK, util.CreateResponse(controller.ListAllUserInfo(ctx, request)))
}

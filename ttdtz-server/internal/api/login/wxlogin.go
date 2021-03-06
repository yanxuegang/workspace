package login

import (
	"log"
	"net/http"
	"ttdtz-server/global"
	"ttdtz-server/internal/managers"
	"ttdtz-server/pkg/app"
	"ttdtz-server/pkg/errcode"

	"github.com/gin-gonic/gin"
)

// type wxloginRequestInfo struct {
// 	Cmd    int    `json:"cmd"`
// 	Params params `json:"params"`
// }

// type wxloginResponseInfo struct {
// 	Code   int    `json:"code"`
// 	OpenId string `json:"open_id"`
// 	Sign   int    `json:"sign"`
// 	Money  int    `json:"money"`
// }

// @Summary 微信登录
// @Produce json
// @Param wxloginRequestInfo query wxloginRequestInfo true "wxloginRequestInfo"
// @Param Params query params true "params"
// @Success 200 {object} wxloginResponseInfo "成功"
// @Failure 400 {object} err "请求错误"
// @failure 500 {object} err "内部错误"
// @Router /api/v1/login [get]
func WxLogin(c *gin.Context) {
	log.Printf("WxLogin req 1 = %+v", c)
	req := managers.LoginRequest{}
	response := app.NewResponse(c)
	err := c.ShouldBind(&req)
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{"error": err.Error()})
		return
	}
	log.Printf("WxLogin req = %+v", req)

	c.String(http.StatusOK, "")
	//todo log check params
	//svc := managers.New(c.Request.Context())
	respdata, err := managers.WxLogin(c.Request.Context(), &req)
	if err != nil {
		global.Logger.Error(err)
		response.ToErrorResponse(errcode.InvalidParams)
		return
	}
	response.ToResponse(respdata)
	return
}

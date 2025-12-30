package util

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Response 统一响应结构
type Response struct {
	Code    int    `json:"code"`           // 业务码，0 表示成功，非 0 表示失败
	Message string `json:"message"`        // 提示信息
	Data    any    `json:"data,omitempty"` // 返回数据
}

// Success 构建成功响应
func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code:    0,
		Message: "success",
		Data:    data,
	})
}

// Error 构建错误响应
func Error(c *gin.Context, httpCode int, errCode int, msg string) {
	c.JSON(httpCode, Response{
		Code:    errCode,
		Message: msg,
		Data:    nil,
	})
}

// HandleError 统一错误处理，自动识别 BizError 类型
func HandleError(c *gin.Context, err error) {
	if bizErr, ok := err.(*BizError); ok {
		c.JSON(bizErr.HttpCode, Response{
			Code:    bizErr.Code,
			Message: bizErr.Msg,
			Data:    nil,
		})
		return
	}
	// 未知错误统一返回 500
	c.JSON(http.StatusInternalServerError, Response{
		Code:    50000,
		Message: "服务器内部错误",
		Data:    nil,
	})
}

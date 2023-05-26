package core

import (
	"github.com/gin-gonic/gin"
	"gulimall-product/internal/pkg/errno"
	"net/http"
)

// Response 定义了返回消息.包括成功和失败的
type Response struct {
	// Code 指定了业务错误码.
	Code int `json:"code"`

	// Message 包含了可以直接对外展示的错误信息.
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// WriteResponse 将错误或响应数据写入 HTTP 响应主体。
// WriteResponse 使用 errno.Decode 方法，根据错误类型，尝试从 errno 中提取业务错误码和错误信息.
func WriteResponse(c *gin.Context, err error, data interface{}) {
	if err != nil {
		hcode, code, message := errno.Decode(err)
		c.JSON(hcode, Response{
			Code:    code,
			Message: message,
			Data:    data,
		})

		return
	}

	c.JSON(http.StatusOK, Response{
		Code:    http.StatusOK,
		Message: "success",
		Data:    data,
	})
}

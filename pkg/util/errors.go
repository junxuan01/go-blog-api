package util

import "net/http"

// BizError 业务错误类型
type BizError struct {
	HttpCode int    // HTTP 状态码
	Code     int    // 业务错误码
	Msg      string // 错误信息
}

// Error 实现 error 接口
func (e *BizError) Error() string {
	return e.Msg
}

// NewBizError 创建业务错误
func NewBizError(httpCode, code int, msg string) *BizError {
	return &BizError{
		HttpCode: httpCode,
		Code:     code,
		Msg:      msg,
	}
}

// 预定义常用业务错误（可根据业务扩展）
var (
	// 客户端错误 4xx
	ErrBadRequest         = NewBizError(http.StatusBadRequest, 40000, "请求参数错误")
	ErrInvalidParam       = NewBizError(http.StatusBadRequest, 40001, "参数校验失败")
	ErrInvalidCredentials = NewBizError(http.StatusBadRequest, 40002, "用户名或密码错误")
	ErrUnauthorized       = NewBizError(http.StatusUnauthorized, 40100, "未授权，请先登录")
	ErrTokenExpired       = NewBizError(http.StatusUnauthorized, 40101, "登录已过期")
	ErrForbidden          = NewBizError(http.StatusForbidden, 40300, "无权限访问")
	ErrNotFound           = NewBizError(http.StatusNotFound, 40400, "资源不存在")
	ErrUserNotFound       = NewBizError(http.StatusNotFound, 40401, "用户不存在")
	ErrArticleNotFound    = NewBizError(http.StatusNotFound, 40402, "文章不存在")
	ErrConflict           = NewBizError(http.StatusConflict, 40900, "资源冲突")
	ErrUsernameExists     = NewBizError(http.StatusConflict, 40901, "用户名已存在")
	ErrEmailExists        = NewBizError(http.StatusConflict, 40902, "邮箱已被注册")

	// 服务端错误 5xx
	ErrInternal = NewBizError(http.StatusInternalServerError, 50000, "服务器内部错误")
	ErrDatabase = NewBizError(http.StatusInternalServerError, 50001, "数据库错误")
)

// WithMsg 复制错误并替换消息（用于动态消息场景）
func (e *BizError) WithMsg(msg string) *BizError {
	return &BizError{
		HttpCode: e.HttpCode,
		Code:     e.Code,
		Msg:      msg,
	}
}

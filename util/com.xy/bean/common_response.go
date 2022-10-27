package bean

import (
	"base-framework/util/com.xy/exception"
)

/**
 * @description: 统一返回对象包装信息
 * @author:xy
 * @date:2022/8/1 17:05
 * @Version: 1.0
 */

type CommonResponse struct {
	Code int `json:"code"`

	Msg string `json:"msg"`

	Data any `json:"data"`
}

func SuccessResponse(data any) *CommonResponse {
	return &CommonResponse{
		Code: exception.OK.Code(),
		Msg:  exception.OK.String(),
		Data: data,
	}
}

func FailResponse(code int, msg string) *CommonResponse {
	return &CommonResponse{
		Code: code,
		Msg:  msg,
	}
}

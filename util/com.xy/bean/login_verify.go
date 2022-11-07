package bean

import (
	"base-framework/util/com.xy/exception"
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/metadata"
	"github.com/go-kratos/kratos/v2/transport"
)

/**
 * @description:
 * @author:xy
 * @date:2022/11/1 14:13
 * @Version: 1.0
 */

var AddressKey = "x-md-global-address"

//VerifyLoginAddress .验证http请求头全局地址
func VerifyLoginAddress(ctx context.Context) (context.Context, error) {
	if serverContext, ok := transport.FromServerContext(ctx); ok {
		get := serverContext.RequestHeader().Get("loginAddress")
		if CheckNonZero(get) {
			background := context.Background()
			background = metadata.AppendToClientContext(ctx, AddressKey, get)
			return background, nil
		}
	}
	return ctx, errors.New(exception.LoginAddress.Code(), exception.LoginAddress.String(), exception.LoginAddress.String())
}

//GetLoginAddress .元数据传输
func GetLoginAddress(ctx context.Context) string {
	md, _ := metadata.FromServerContext(ctx)
	return md.Get(AddressKey)
}

package exception

/**
 * @description: 公共枚举
 * @author:xy
 * @date:2022/8/1 17:15
 * @Version: 1.0
 */

type ErrCode int

const (
	OK           ErrCode = 200  //请求成功
	ServerError  ErrCode = 500  //请求异常
	InvalidParam ErrCode = 4006 //参数错误
	LoginAddress ErrCode = 401  //登录用户获取失败
)

var errMsg = map[ErrCode]string{
	OK:           "OK",
	ServerError:  "SERVER_ERROR",
	InvalidParam: "INVALID_PARAM",
	LoginAddress: "LOGIN_USER_ERROR",
}

func (e ErrCode) String() string {
	return errMsg[e]
}

func (e ErrCode) Code() int {
	return int(e)
}

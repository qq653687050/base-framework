package exception

import "github.com/go-kratos/kratos/v2/errors"

/**
 * @description: 中台异常枚举  范围20000~29999
 * @author:xy
 * @date:2022/10/25 17:22
 * @Version: 1.0
 */

type FacadeErrCode int

const (
	TimeRangeError FacadeErrCode = 20001 //时间范围错误
)

var facadeErrReason = map[FacadeErrCode]string{
	TimeRangeError: "TIME_RANGE_ERROR",
}

var facadeErrMsg = map[FacadeErrCode]string{
	TimeRangeError: "时间范围错误",
}

func (e FacadeErrCode) ReasonString() string {
	return facadeErrReason[e]
}

func (e FacadeErrCode) Code() int {
	return int(e)
}

func (e FacadeErrCode) MsgString() string {
	return facadeErrMsg[e]
}

func (e FacadeErrCode) Error() *errors.Error {
	return errors.New(e.Code(), e.ReasonString(), e.MsgString())
}

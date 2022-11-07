package exception

import "github.com/go-kratos/kratos/v2/errors"

/**
 * @description: 中心异常枚举  范围10000~19999
 * @author:xy
 * @date:2022/10/25 16:17
 * @Version: 1.0
 */

type CenterErrCode int

const (
	ActivityIsNot          CenterErrCode = 10001 //活动不存在
	ActivityRankRepetition CenterErrCode = 10002 //活动排序重复
	AddressIsNot           CenterErrCode = 10003 //地址不存在
)

var centerErrReason = map[CenterErrCode]string{
	ActivityIsNot:          "ACTIVITY_IS_NOT",
	ActivityRankRepetition: "ACTIVITY_RANK_REPETITION",
	AddressIsNot:           "ADDRESS_IS_NOT",
}

var centerErrMsg = map[CenterErrCode]string{
	ActivityIsNot:          "活动不存在",
	ActivityRankRepetition: "活动排序重复",
	AddressIsNot:           "地址不存在",
}

func (e CenterErrCode) ReasonString() string {
	return centerErrMsg[e]
}

func (e CenterErrCode) Code() int {
	return int(e)
}

func (e CenterErrCode) MsgString() string {
	return centerErrMsg[e]
}

func (e CenterErrCode) Error() *errors.Error {
	return errors.New(e.Code(), e.ReasonString(), e.MsgString())
}

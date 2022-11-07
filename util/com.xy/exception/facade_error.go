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
	TimeRangeError     FacadeErrCode = 20001 //时间范围错误
	CreationTimeError  FacadeErrCode = 20002 //不在活动时间范围内
	CreationAddedError FacadeErrCode = 20003 //已加入活动
	PollWordsError     FacadeErrCode = 20004 //重复点赞或取消
	WordsError         FacadeErrCode = 20005 //作品错误
)

var facadeErrReason = map[FacadeErrCode]string{
	TimeRangeError:     "TIME_RANGE_ERROR",
	CreationTimeError:  "CREATION_TIME_ERROR",
	CreationAddedError: "CREATION_ADDED_ERROR",
	PollWordsError:     "POLL_WORDS_ERROR",
	WordsError:         "WORDS_ERROR",
}

var facadeErrMsg = map[FacadeErrCode]string{
	TimeRangeError:     "时间范围错误",
	CreationTimeError:  "不在活动时间范围内",
	CreationAddedError: "已加入活动",
	PollWordsError:     "请勿重复操作",
	WordsError:         "作品错误",
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

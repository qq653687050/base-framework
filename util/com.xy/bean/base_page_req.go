package bean

/**
 * @description: 分页请求基础结构
 * @author:xy
 * @date:2022/8/1 16:51
 * @Version: 1.0
 */

type BasePageReq struct {
	/**
	 * 当前页
	 */
	PageNum int32

	/**
	 * 每页条数
	 */
	PageSize int32
}

func (b *BasePageReq) GetLimit() int32 {
	return b.PageSize
}

func (b *BasePageReq) GetOffset() int32 {
	return (b.PageNum - 1) * b.PageSize
}

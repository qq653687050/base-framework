package bean

/**
 * @description: 分页请求返回基础结构
 * @author:xy
 * @date:2022/8/1 16:51
 * @Version: 1.0
 */

type BasePageResp[T any] struct {
	/**
	 * 数据总条数
	 */
	Total int64 `json:"total"`

	/**
	 * 当前页
	 */
	PageNum int32 `json:"pageNum"`

	/**
	 * 每页大小
	 */
	PageSize int32 `json:"pageSize"`

	/**
	 * 返回结果
	 */
	ResultList []T `json:"resultList"`
}

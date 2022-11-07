package bean

/**
 * @description:
 * @author:xy
 * @date:2022/8/11 10:42
 * @Version: 1.0
 */

type whereParamPair struct {
	Query string // 条件
	Val   []any  // 参数
}

type orderByCol struct {
	Column string // 排序字段
	Asc    bool   // 是否正序
}

type Wrapper struct {
	whereParams []whereParamPair // where条件
	setParams   map[string]any   // set字段
	orders      []orderByCol     // 排序
	paging      BasePageReq      //分页
	group       []string
}

func NewWrapper() *Wrapper {
	w := &Wrapper{}
	return w
}

func (w *Wrapper) where(query string, val ...any) *Wrapper {
	w.whereParams = append(w.whereParams, whereParamPair{Query: query, Val: val})
	return w
}

func (w *Wrapper) Eq(column string, val any) *Wrapper {

	return w.ConEq(true, column, val)
}

func (w *Wrapper) ConEq(condition bool, column string, val any) *Wrapper {
	if condition {
		w.where("`"+column+"` = ?", val)
	}
	return w
}

func (w *Wrapper) NotEq(column string, val any) *Wrapper {

	return w.ConNotEq(true, column, val)
}

func (w *Wrapper) ConNotEq(condition bool, column string, val any) *Wrapper {
	if condition {
		w.where("`"+column+"` <> ?", val)
	}
	return w
}

func (w *Wrapper) Gt(column string, val any) *Wrapper {

	return w.ConGt(true, column, val)
}

func (w *Wrapper) ConGt(condition bool, column string, val any) *Wrapper {
	if condition {
		w.where("`"+column+"` > ?", val)
	}
	return w
}

func (w *Wrapper) Gte(column string, val any) *Wrapper {

	return w.ConGte(true, column, val)
}

func (w *Wrapper) ConGte(condition bool, column string, val any) *Wrapper {
	if condition {
		w.where("`"+column+"` >= ?", val)
	}
	return w
}

func (w *Wrapper) Lt(column string, val any) *Wrapper {

	return w.ConLt(true, column, val)
}

func (w *Wrapper) ConLt(condition bool, column string, val any) *Wrapper {
	if condition {
		w.where("`"+column+"` < ?", val)
	}
	return w
}

func (w *Wrapper) Lte(column string, val any) *Wrapper {

	return w.ConLte(true, column, val)
}

func (w *Wrapper) ConLte(condition bool, column string, val any) *Wrapper {
	if condition {
		w.where("`"+column+"` <= ?", val)
	}
	return w
}

func (w *Wrapper) Like(column string, val string) *Wrapper {
	return w.ConLike(true, column, val)
}

func (w *Wrapper) ConLike(condition bool, column string, val string) *Wrapper {
	if condition {
		w.where("`"+column+"` LIKE ?", "%"+val+"%")
	}
	return w
}

func (w *Wrapper) Between(column string, val1 any, val2 any) *Wrapper {

	return w.ConBetween(true, column, val1, val2)
}

func (w *Wrapper) ConBetween(condition bool, column string, val1 any, val2 any) *Wrapper {
	if condition {
		w.where("`"+column+"`between ? and ?", val1, val2)
	}
	return w
}

func (w *Wrapper) Set(column string, val any) *Wrapper {
	return w.ConSet(true, column, val)
}

func (w *Wrapper) ConSet(condition bool, column string, val any) *Wrapper {
	if condition {
		if w.setParams == nil {
			w.setParams = make(map[string]any)
		}
		w.setParams[column] = val
	}
	return w
}

func (w *Wrapper) In(column string, val any) *Wrapper {
	w.where("`"+column+"` in (?) ", val)
	return w
}

func (w *Wrapper) Order(asc bool, column string) *Wrapper {
	if column == "" {
		column = "id"
	}
	w.orders = append(w.orders, orderByCol{
		Column: "`" + column + "`",
		Asc:    asc,
	})
	return w
}

func (w *Wrapper) Asc(column string) *Wrapper {
	w.orders = append(w.orders, orderByCol{
		Column: "`" + column + "`",
		Asc:    true,
	})
	return w
}

func (w *Wrapper) Desc(column string) *Wrapper {
	w.orders = append(w.orders, orderByCol{
		Column: "`" + column + "`",
		Asc:    false,
	})
	return w
}

func (w *Wrapper) Group(column string) *Wrapper {
	w.group = append(w.group, "`"+column+"`")
	return w
}

func (w *Wrapper) Limit(limit int32) *Wrapper {
	w.paging.PageSize = limit
	return w
}

func (w *Wrapper) Page(pageNum, pageSize int32) *Wrapper {
	if pageNum == 0 {
		w.paging.PageNum = 1
	} else {
		w.paging.PageNum = pageNum
	}
	if pageSize == 0 {
		w.paging.PageSize = 10
	} else {
		w.paging.PageSize = pageSize
	}
	return w

}

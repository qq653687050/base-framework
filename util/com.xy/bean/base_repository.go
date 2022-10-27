package bean

import (
	"base-framework/inf/base_data"
	"context"
	"gorm.io/gorm"
	"reflect"
)

/**
 * @description:
 * @author:xy
 * @date:2022/8/10 17:05
 * @Version: 1.0
 */

type BaseRepository[T any] struct {
	Data *base_data.Data
}

func (b *BaseRepository[T]) Insert(ctx context.Context, entity any) int64 {
	if create := b.Data.DB(ctx).Create(entity); create.Error != nil {
		panic(create.Error.Error())
	} else {
		return create.RowsAffected
	}
}

func (b *BaseRepository[T]) InsertBatch(ctx context.Context, entity []any) int64 {
	if create := b.Data.DB(ctx).Create(entity); create.Error != nil {
		panic(create.Error.Error())
	} else {
		return create.RowsAffected
	}
}

func (b *BaseRepository[T]) InsertBatches(ctx context.Context, entity []any, num int) int64 {
	if batches := b.Data.DB(ctx).CreateInBatches(entity, num); batches.Error != nil {
		panic(batches.Error.Error())
	} else {
		return batches.RowsAffected
	}
}

func (b *BaseRepository[T]) DeleteById(ctx context.Context, id any) int64 {
	if del := b.Data.DB(ctx).Delete(new(T), "id = ?", id); del.Error != nil {
		panic(del.Error.Error())
	} else {
		return del.RowsAffected
	}
}

func (b *BaseRepository[T]) DeleteWrapper(ctx context.Context, w *Wrapper) int64 {
	db := b.Data.DB(ctx)
	db = where(db, w)
	if del := db.Delete(new(T)); del.Error != nil {
		panic(del.Error.Error())
	} else {
		return del.RowsAffected
	}
}

func (b *BaseRepository[T]) UpdateById(ctx context.Context, entity any) int64 {
	of := reflect.TypeOf(entity)
	if of.Kind() != reflect.Struct {
		panic("entity is not struct")
	}
	name, b2 := of.FieldByName("ID")
	byName, b3 := of.FieldByName("Id")
	value := reflect.ValueOf(entity)
	var id any
	if b2 {
		id = value.FieldByName(name.Name).Interface()
	}
	if b3 {
		id = value.FieldByName(byName.Name).Interface()
	}
	if id == "" || id == nil {
		panic("primary key is empty")
	}
	if updates := b.Data.DB(ctx).Model(new(T)).Updates(entity); updates.Error != nil {
		return updates.RowsAffected
	} else {
		panic(updates.Error.Error())
	}

}

func (b *BaseRepository[T]) Update(ctx context.Context, w *Wrapper) int64 {
	db := b.Data.DB(ctx).Model(new(T))
	db = where(db, w)
	if updates := db.Updates(w.setParams); updates.Error != nil {
		panic(updates.Error.Error())
	} else {
		return updates.RowsAffected
	}
}

func (b *BaseRepository[T]) SelectById(id any) T {
	var t T
	if find := b.Data.GDB().Find(&t, id); find.Error != nil {
		panic(find.Error.Error())
	}
	return t
}

func (b *BaseRepository[T]) SelectBatchIds(id []any) []T {
	var ts []T
	if find := b.Data.GDB().Find(&ts, id); find.Error != nil {
		panic(find.Error.Error())
	}
	return ts
}

func (b *BaseRepository[T]) SelectOne(w *Wrapper) T {
	var t T
	db := b.Data.GDB()
	db = where(db, w)
	db = order(db, w)
	if take := db.Find(&t).Limit(1); take.Error != nil {
		panic(take.Error.Error())
	}
	return t
}

func (b *BaseRepository[T]) Count(w *Wrapper) (count int64) {
	db := b.Data.GDB()
	db = where(db, w)
	if d := db.Model(new(T)).Count(&count); d.Error != nil {
		panic(d.Error.Error())
	}
	return
}

func (b *BaseRepository[T]) List(w *Wrapper) []T {
	db := b.Data.GDB()
	db = where(db, w)
	db = order(db, w)
	db = group(db, w)
	var ts []T
	if find := db.Find(&ts); find.Error != nil {
		panic(find.Error.Error())
	}
	return ts
}

func (b *BaseRepository[T]) SelectPage(w *Wrapper) BasePageResp[T] {
	db := b.Data.GDB()
	db = where(db, w)
	db = order(db, w)
	db = group(db, w)
	db = setPage(db, w)
	var t []T

	if find := db.Find(&t); find.Error != nil {
		panic(find.Error.Error())
	}

	resp := BasePageResp[T]{
		Total:      0,
		PageNum:    w.paging.PageNum,
		PageSize:   w.paging.PageSize,
		ResultList: t,
	}
	if count := db.Count(&resp.Total); count.Error != nil {
		panic(count.Error.Error())
	}

	return resp
}

func where(db *gorm.DB, w *Wrapper) *gorm.DB {
	if w != nil {
		if len(w.whereParams) > 0 {
			for _, param := range w.whereParams {
				db = db.Where(param.Query, param.Val...)
			}
		}
	}
	return db
}

func order(db *gorm.DB, w *Wrapper) *gorm.DB {
	if w != nil {
		if len(w.orders) > 0 {
			for _, param := range w.orders {
				if param.Asc {
					db = db.Order(param.Column + " asc")
				} else {
					db = db.Order(param.Column + " desc")
				}

			}
		}
	}
	return db
}

func group(db *gorm.DB, w *Wrapper) *gorm.DB {
	if w != nil {
		if len(w.group) > 0 {
			for _, param := range w.group {
				db = db.Group(param)
			}
		}
	}
	return db
}

func setPage(db *gorm.DB, w *Wrapper) *gorm.DB {
	if w == nil {
		w = &Wrapper{paging: BasePageReq{1, 10}}
	}
	db = db.Limit(int(w.paging.GetLimit())).Offset(int(w.paging.GetOffset()))
	return db
}

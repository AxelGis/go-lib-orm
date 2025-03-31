package orm

import (
	"github.com/AxelGis/go-lib-orm/util"
	"gorm.io/gorm"
)

type BaseORM[T any] struct {
	db *gorm.DB
}

func NewBaseORM[T any](db *gorm.DB) *BaseORM[T] {
	return &BaseORM[T]{db: db}
}

func (r *BaseORM[T]) Create(entity *T) error {
	return r.db.Create(entity).Error
}

func (r *BaseORM[T]) Update(entity *T) error {
	return r.db.Save(entity).Error
}

func (r *BaseORM[T]) UpdateWhere(conditions map[string]interface{}, updates map[string]interface{}) error {
	return r.db.Model(new(T)).Where(conditions).Updates(updates).Error
}

func (r *BaseORM[T]) FindById(id uint) (*T, error) {
	var entity T
	err := r.db.First(&entity, id).Error
	if err != nil {
		return nil, err
	}
	return &entity, nil
}

func (r *BaseORM[T]) FindOne(filters []util.Filter, sorts []util.Sort, relations ...string) (*T, error) {
	var entity T
	query := r.applyFiltersAndSorts(r.db, filters, sorts)
	for _, relation := range relations {
		query = query.Preload(relation)
	}
	err := query.First(&entity).Error
	if err != nil {
		return nil, err
	}
	return &entity, nil
}

func (r *BaseORM[T]) FindAll(pagination *PageInput, filters []util.Filter, sorts []util.Sort, relations ...string) ([]T, error) {
	var entities []T
	query := r.applyFiltersAndSorts(r.db, filters, sorts)
	for _, relation := range relations {
		query = query.Preload(relation)
	}

	// Apply pagination
	if pagination != nil {
		offset := (pagination.Num - 1) * pagination.Size
		query = query.Offset(offset).Limit(pagination.Size)
	}

	err := query.Find(&entities).Error
	if err != nil {
		return nil, err
	}
	return entities, nil
}

func (r *BaseORM[T]) UpdateBy(conditions map[string]interface{}, updates map[string]interface{}) (*T, error) {
	err := r.db.Model(new(T)).Where(conditions).Updates(updates).Error
	if err != nil {
		return nil, err
	}

	var entity T
	err = r.db.Where(conditions).First(&entity).Error
	if err != nil {
		return nil, err
	}
	return &entity, nil
}

func (r *BaseORM[T]) Delete(id uint) error {
	return r.db.Delete(new(T), id).Error
}

func (r *BaseORM[T]) applyFiltersAndSorts(query *gorm.DB, filters []util.Filter, sorts []util.Sort) *gorm.DB {
	// Apply filters
	for _, filter := range filters {
		if filter.CustomQuery != "" {
			query = query.Where(filter.CustomQuery)
			continue
		}

		query = query.Where(util.Conditions[filter.Op].String(), filter.Field, filter.Value)
	}

	// Apply sorting
	for _, sort := range sorts {
		query = query.Order(sort.Field + util.Orders[sort.Direction].String())
	}

	return query
}

type PageInput struct {
	Num  int
	Size int
}

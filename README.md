# go-lib-orm

Golang ORM library for managing database operations with GORM.

## Usage

```go
package main

import (
    "context"
    "gorm.io/gorm"
    "your_project/orm"
    "your_project/entity"
)

type repo struct {
    db        *gorm.DB
    bussiness *orm.BaseORM[entity.Business]
}

func NewRepo(_ context.Context, db *gorm.DB) Repo {
    return &repo{
        db:        db,
        bussiness: orm.NewBaseORMentity.Business,
    }
}

// Entity work
func (r *repo) CreateBusiness(ctx context.Context, item entity.Business) (*entity.Business, error) {
    err := r.bussiness.Create(&item)
    if err != nil {
        return nil, err
    }

    return r.bussiness.FindById(item.ID)
}

func (r *repo) GetBusiness(ctx context.Context, id uint) (*entity.Business, error) {
    return r.bussiness.FindById(id)
}

func (r *repo) UpdateBusiness(ctx context.Context, id uint, item entity.Business) (*entity.Business, error) {
    err := r.bussiness.Update(&item)
    if err != nil {
        return nil, err
    }

    return r.bussiness.FindById(item.ID)
}

func (r *repo) DeleteBusiness(ctx context.Context, id uint) error {
    return r.bussiness.Delete(id)
}
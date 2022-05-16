package service

import (
	"context"

	"github.com/tris-tux/go-lis/backend/db"
	"github.com/tris-tux/go-lis/backend/schema"
)

func GetAll(ctx context.Context) ([]schema.Task, error) {
	return db.GetAll(ctx)
}

func Insert(ctx context.Context, task *schema.Task) (int, error) {
	return db.Insert(ctx, task)
}

func Update(ctx context.Context, task *schema.Task) error {
	return db.Update(ctx, task)
}

func Delete(ctx context.Context, id int) error {
	return db.Delete(ctx, id)
}

func Close(ctx context.Context) {
	db.Close(ctx)
}

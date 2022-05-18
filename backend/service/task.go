package service

import (
	"context"

	"github.com/tris-tux/go-lis/backend/db"
	"github.com/tris-tux/go-lis/backend/schema"
)

func GetAll(ctx context.Context) ([]schema.Result, error) {
	return db.GetAll(ctx)
}

func Insert(ctx context.Context, task *schema.Task) (string, error) {
	return db.Insert(ctx, task)
}

func Update(ctx context.Context, result *schema.Result) (string, error) {
	return db.Update(ctx, result)
}

func Delete(ctx context.Context, task_id int) (string, error) {
	return db.Delete(ctx, task_id)
}

func Close(ctx context.Context) {
	db.Close(ctx)
}

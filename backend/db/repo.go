package db

import (
	"context"

	"github.com/tris-tux/go-lis/backend/schema"
)

const repoKey = "repoKey"

type Repo interface {
	GetAll() ([]schema.Result, error)
	Insert(task *schema.Task) (string, error)
	Update(result *schema.Result) (string, error)
	Delete(task_id int) (string, error)
	Close()
}

func SetRepo(ctx context.Context, repo Repo) context.Context {
	return context.WithValue(ctx, repoKey, repo)
}

func getRepo(ctx context.Context) Repo {
	return ctx.Value(repoKey).(Repo)
}

func GetAll(ctx context.Context) ([]schema.Result, error) {
	return getRepo(ctx).GetAll()
}

func Insert(ctx context.Context, task *schema.Task) (string, error) {
	return getRepo(ctx).Insert(task)
}

func Update(ctx context.Context, result *schema.Result) (string, error) {
	return getRepo(ctx).Update(result)
}

func Delete(ctx context.Context, task_id int) (string, error) {
	return getRepo(ctx).Delete(task_id)
}

func Close(ctx context.Context) {
	getRepo(ctx).Close()
}

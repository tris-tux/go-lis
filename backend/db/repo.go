package db

import (
	"context"

	"github.com/tris-tux/go-lis/backend/schema"
)

const repoKey = "repoKey"

type Repo interface {
	GetAll() ([]schema.Task, error)
	Insert(task *schema.Task) (int, error)
	Update(task *schema.Task) error
	Delete(id int) error
	Close()
}

func SetRepo(ctx context.Context, repo Repo) context.Context {
	return context.WithValue(ctx, repoKey, repo)
}

func getRepo(ctx context.Context) Repo {
	return ctx.Value(repoKey).(Repo)
}

func GetAll(ctx context.Context) ([]schema.Task, error) {
	return getRepo(ctx).GetAll()
}

func Insert(ctx context.Context, task *schema.Task) (int, error) {
	return getRepo(ctx).Insert(task)
}

func Update(ctx context.Context, task *schema.Task) error {
	return getRepo(ctx).Update(task)
}

func Delete(ctx context.Context, id int) error {
	return getRepo(ctx).Delete(id)
}

func Close(ctx context.Context) {
	getRepo(ctx).Close()
}

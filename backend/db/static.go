package db

import (
	"github.com/tris-tux/go-lis/backend/schema"
)

type Static struct{}

func (s *Static) GetAll() ([]schema.Task, error) {
	taskList := []schema.Task{
		{
			TaskId:   1,
			Title: []string{"AAA", "AA"},
			AcctionTime: 1652461200,
			// CreateTime: 1652461200,
			// UpdateTime: 1652461200,
			// IsFinished: false,
		},
		{
			TaskId:   2,
			Title: []string{"BBB", "BB"},
			AcctionTime: 1652462100,
			// CreateTime: 1652462100,
			// UpdateTime: 1652462100,
			// IsFinished: false,
		},
		{
			TaskId:   3,
			Title: []string{"CCC", "CC"},
			AcctionTime: 1652661200,
			// CreateTime: 1652661200,
			// UpdateTime: 1652661200,
			// IsFinished: false,
		},
	}
	return taskList, nil
}

func (s *Static) Insert(task *schema.Task) ([]string, error) {
	return []string{"0"}, nil
}

func (s *Static) Update(task *schema.Task) error {
	return nil
}

func (s *Static) Delete(id int) error {
	return nil
}

func (s *Static) Close() {}

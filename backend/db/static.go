package db

import (
	"github.com/tris-tux/go-lis/backend/schema"
)

type Static struct{}

func (s *Static) GetAll() ([]schema.Result, error) {
	List := []schema.List{
		{
			DetailId: 1,
			ObjectName: "SAD",
			IsFinished: false,
		},
	}
	taskList := []schema.Result{
		{
			TaskId:   1,
			Title: "SATU",
			AcctionTime: 1652461200,
			// ObjectiveList: []string{"AAA", "AA"},
			CreateTime: 1652461200,
			UpdateTime: 1652461200,
			IsFinished: false,
			ObjectiveList: List,
		},
		// {
		// 	DetailId: 1,
		// 	ObjectName: "SAD",
		// 	IsFinished: false,
		// },
		// {
		// 	TaskId:   1,
		// 	Title: "SATU",
		// 	AcctionTime: 1652461200,
		// 	// ObjectiveList: []string{"AAA", "AA"},
		// 	CreateTime: 1652461200,
		// 	UpdateTime: 1652461200,
		// 	IsFinished: false,
		// },
		// {
		// 	TaskId:   2,
		// 	Title: "DUA",
		// 	AcctionTime: 1652462100,
		// 	// ObjectiveList: []string{"BBB", "BB"},
		// 	CreateTime: 1652462100,
		// 	UpdateTime: 1652462100,
		// 	IsFinished: false,
		// },
		// {
		// 	TaskId:   3,
		// 	Title: "TIGA",
		// 	AcctionTime: 1652661200,
		// 	// ObjectiveList: []string{"CCC", "CC"},
		// 	CreateTime: 1652661200,
		// 	UpdateTime: 1652661200,
		// 	IsFinished: false,
		// },
	}
	return taskList, nil
}

func (s *Static) Insert(task *schema.Task) (string, error) {
	return "success", nil
}

func (s *Static) Update(result *schema.Result) (string, error) {
	return "success", nil
}

func (s *Static) Delete(task_id int) (string, error) {
	return "success", nil
}

func (s *Static) Close() {}

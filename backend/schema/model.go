package schema

type Task struct {
	TaskId int `json:"task_id"`
	Title string `json:"title"`
	AcctionTime int `json:"acction_time"`
	ObjectiveList []string `json:"objective_list"`
	// CreateTime int `json:"create_time"`
	// UpdateTime int `json:"update_time"`
	// IsFinished bool `json:"is_finished"`
}

type Result struct{
	TaskId int `json:"task_id"`
	Title string `json:"title"`
	AcctionTime int `json:"acction_time"`
	CreateTime int `json:"create_time"`
	UpdateTime int `json:"update_time"`
	IsFinished bool `json:"is_finished"`
	ObjectiveList []List `json:"obejct_list"`
}

type Tasks struct {
	TaskId int `json:"task_id"`
	Title string `json:"title"`
	AcctionTime int `json:"acction_time"`
	CreateTime int `json:"create_time"`
	UpdateTime int `json:"update_time"`
	IsFinished bool `json:"is_finished"`
	// ObjectiveList List `json:"obejct_list"`
}

type List struct {
	DetailId int `json:"detail_id`
	ObjectName string `json:"object_name"`
	IsFinished bool `json:"is_finished"`
}


// type Detail struct {
// 	DetailId int `json:"detail_id"`
// 	ObjectTaskFk int `json:"object_task_fk"`
// 	ObjectName string `json:"object_name"`
// 	IsFinished bool `json:"is_finished"`
// }
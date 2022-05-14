package schema

type Task struct {
	TaskId int `json:"taskId"`
	Title string `json:"title"`
	AcctionTime int `json:"acctionTime"`
	CreateTime int `json:"createTime"`
	UpdateTime int `json:"updateTime"`
	IdFinished bool `json:"idFinished"`
}

type Detail struct {
	DetailId int `json:"detailId"`
	ObjectTaskFk int `json:"objectTaskFk"`
	ObjectName string `json:"objectName"`
	IdFinished int `json:"idFinished"`
}
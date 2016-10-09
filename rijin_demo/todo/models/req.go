package models

type AddTask struct {
	Id          string `json:"id"`
	OwnerId     string `json:"owner_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type GetTaskReq struct {
	Id string `json:"id"`
}

type GetTaskResp struct {
	Id          string `json:"id"`
	OwnerId     string `json:"owner_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

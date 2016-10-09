package handlers

import (
	"github.com/gogap/spirit"

	"git.rd.rijin.com/rijin_demo/todo/infrastructure"
	"git.rd.rijin.com/rijin_demo/todo/models"
)

func GetTask(msg *spirit.Payload) (result interface{}, err error) {
	req := &models.GetTaskReq{}
	if err = msg.FillContentToObject(req); err != nil {
		return
	}

	todo := &models.Task{
		Id: req.Id,
	}

	if _, err = infrastructure.SQLEngine.Get(todo); err != nil {
		return
	}

	resp := &models.GetTaskResp{
		Id:          req.Id,
		OwnerId:     todo.OwnerId,
		Title:       todo.Title,
		Description: todo.Description,
	}

	result = resp
	return
}

package handlers

import (
	"fmt"
	"github.com/gogap/spirit"
	"github.com/pborman/uuid"

	"git.rd.rijin.com/rijin_demo/todo/infrastructure"
	"git.rd.rijin.com/rijin_demo/todo/models"
)

func AddTask(msg *spirit.Payload) (result interface{}, err error) {
	fmt.Println(">>>>>>>>")
	req := &models.AddTask{}
	if err = msg.FillContentToObject(req); err != nil {
		return
	}

	todo := &models.Task{
		Id:          uuid.NewUUID().String(),
		OwnerId:     req.OwnerId,
		Title:       req.Title,
		Description: req.Description,
	}

	if _, err = infrastructure.SQLEngine.Insert(todo); err != nil {
		return
	}

	return
}

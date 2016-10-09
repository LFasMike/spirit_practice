package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	"github.com/gogap/spirit"

	"git.rd.rijin.com/components/utils"
	"git.rd.rijin.com/rijin_demo/todo/handlers"
	"git.rd.rijin.com/rijin_demo/todo/infrastructure"
)

var (
	sqlEngine *xorm.Engine
)

func initComponent() (err error) {
	engines, err := utils.InitalXORMEngineFromConfig("conf/databases.conf")
	if err != nil {
		utils.Log().Error("init xorm failed:", err)
		return
	}

	for _, v := range engines {
		v.Logger().SetLevel(core.LOG_INFO)
	}
	sqlEngine, _ = engines["rijin"]
	infrastructure.SetXormEngines(engines)
	return
}

func main() {
	todoSpirit := spirit.NewClassicSpirit(
		"todo",
		"this is a todo list project",
		"1.0.0",
		[]spirit.Author{
			{Name: "wangxigang", Email: "wangxigang@rijin.com"},
		},
	)

	todoComponent := spirit.NewBaseComponent("todo")

	todoComponent.RegisterHandler("todo_add", handlers.AddTask)
	todoComponent.RegisterHandler("todo_get", handlers.GetTask)

	todoSpirit.Hosting(todoComponent, initComponent).Build().Run()
}

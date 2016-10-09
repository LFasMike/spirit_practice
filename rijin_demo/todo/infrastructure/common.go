package infrastructure

import (
	"github.com/go-xorm/xorm"
)

var (
	xormEngines map[string]*xorm.Engine
	SQLEngine   *xorm.Engine
)

func SetXormEngines(engines map[string]*xorm.Engine) {
	xormEngines = engines
	SQLEngine = xormEngines["todo"]
}

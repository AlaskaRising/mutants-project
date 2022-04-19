package routers

import (
	"mutants-project/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Router("/mutant", &controllers.MutantController{})
	beego.Router("/stats", &controllers.StatsController{})
}

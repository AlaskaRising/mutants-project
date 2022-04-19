package controllers

import (
	"mutants-project/models"

	beego "github.com/beego/beego/v2/server/web"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var statsValue, statsColl = "statsCollection", "stats"
var statsCollId = 10

//  StatsController operations for Stats
type StatsController struct {
	beego.Controller
	sequenceManager models.SequenceManager
}

// URLMapping ...
func (c *StatsController) URLMapping() {
	c.Mapping("Get", c.Get)
}

// Get ...
// @Title Get
// @Description get Stats
// @Success 200 {object} models.Stats
// @Failure 403 :id is empty
// @router [get]
func (c *StatsController) Get() {
	var tempResult interface{}
	if _ = c.sequenceManager.GetSequenceByItem(statsCollId, statsValue, statsColl, &tempResult); tempResult != nil {
		temp := tempResult.(primitive.D)
		metadata := temp.Map()
		var statsResp models.StatsResponse
		statsResp.Count_human_dna = metadata["count_human_dna"].(float64)
		statsResp.Count_mutant_dna = metadata["count_mutant_dna"].(float64)
		statsResp.Ratio = metadata["ratio"].(float64)
		c.Data["json"] = statsResp
	}
	c.ServeJSON()
}

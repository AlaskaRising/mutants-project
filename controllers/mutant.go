package controllers

import (
	"encoding/json"
	"log"
	"mutants-project/helpers"
	"mutants-project/models"

	beego "github.com/beego/beego/v2/server/web"
)

// MutantController operations for Mutant
type MutantController struct {
	beego.Controller
	sequenceManager models.SequenceManager
	statsManager    models.StatsManager
}

// URLMapping ...
func (c *MutantController) URLMapping() {
	c.Mapping("Post", c.Post)
}

// Post ...
// @Title Create
// @Description create Mutant
// @Param	dna		body 	[]string	true		"body for Mutant content"
// @Success 200
// @Failure 403 body is empty
// @router / [post]
func (c *MutantController) Post() {
	var dnaSequence models.Sequence
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &dnaSequence); err != nil {
		log.Panicln(err.Error())
	} else if helpers.IsMutant(dnaSequence.Dna) {
		dnaSequence.Mutant = true
		c.Ctx.ResponseWriter.WriteHeader(200)
	} else {
		dnaSequence.Mutant = false
		c.Ctx.ResponseWriter.WriteHeader(403)
	}
	c.sequenceManager.AddDna(&dnaSequence, "mutant-stats")
	c.statsManager.UpdateOneStats(helpers.UpdateStats(dnaSequence.Mutant))
}

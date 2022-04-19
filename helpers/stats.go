package helpers

import (
	"mutants-project/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var statsCollectionBson = "statsCollection"
var collectionMongo = "stats"

func UpdateStats(mutant bool) (statsUpdated models.Stats, statsCollection int32) {
	var stats, statsToUpdate models.Stats
	stats.StatsCollection = 10.0
	var sequenceManager models.SequenceManager
	var tempResult interface{}

	if _ = sequenceManager.GetSequenceByItem(stats.StatsCollection, statsCollectionBson, collectionMongo, &tempResult); tempResult != nil {
		temp := tempResult.(primitive.D)
		metadata := temp.Map() // map to map[string]interface{}
		statsToUpdate.ID = metadata["_id"].(primitive.ObjectID)
		statsToUpdate.Count_human_dna = metadata["count_human_dna"].(float64)
		statsToUpdate.Count_mutant_dna = metadata["count_mutant_dna"].(float64)
		statsToUpdate.StatsCollection = metadata["statsCollection"].(int32)
		if mutant {
			statsToUpdate.Count_mutant_dna++
		} else {
			statsToUpdate.Count_human_dna++
		}
		if statsToUpdate.Count_human_dna > 0 {
			statsToUpdate.Ratio = statsToUpdate.Count_mutant_dna / statsToUpdate.Count_human_dna
		}
	} else {
		statsToUpdate.ID, statsToUpdate.Count_human_dna = primitive.NewObjectID(), 0.0
		statsToUpdate.Count_mutant_dna, statsToUpdate.StatsCollection = 0.0, 10
		statsToUpdate.Ratio = 0.0
		if mutant {
			statsToUpdate.Count_mutant_dna++
		} else {
			statsToUpdate.Count_human_dna++
		}
		if statsToUpdate.Count_human_dna > 0 {
			statsToUpdate.Ratio = statsToUpdate.Count_mutant_dna / statsToUpdate.Count_human_dna
		}
		sequenceManager.AddStats(&statsToUpdate)
	}
	return statsToUpdate, statsToUpdate.StatsCollection
}

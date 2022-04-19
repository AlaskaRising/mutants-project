package models

import (
	"context"
	"errors"
	"log"
	"mutants-project/db"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var collectionStats = "stats"

type Stats struct {
	ID               primitive.ObjectID `json:"ID" bson:"_id"`
	StatsCollection  int32              `json:"StatsCollection" bson:"statsCollection"`
	Count_mutant_dna float64            `json:"Count_mutant_dna" bson:"count_mutant_dna"`
	Count_human_dna  float64            `json:"Count_human_dna" bson:"count_human_dna"`
	Ratio            float64            `json:"Ratio" bson:"ratio"`
}

type StatsManager struct {
	Ctx context.Context
}

// UpdateOneStats Update Stats
func (m *StatsManager) UpdateOneStats(data interface{}, statsCollectionId int32) (err error) {
	coll, err := db.GetCollection(collectionStats)
	if err != nil {
		return err
	}
	filter := make(map[string]interface{})
	filter["statsCollection"] = statsCollectionId

	update := map[string]interface{}{
		"$set": data,
	}

	res := coll.FindOneAndUpdate(m.Ctx, filter, update)
	if res.Err() != nil {
		log.Println("error:", res.Err().Error())
		return errors.New("cannot-update-stats")
	}
	return
}

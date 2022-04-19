package models

import (
	"context"
	"errors"
	"mutants-project/db"
	"strings"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var collectionMongo = "mutant-stats"
var dnaBson = "dna"

type Sequence struct {
	Dna    []string `json:"Dna" bson:"dna"`
	Mutant bool     `json:"Mutant" bson:"mutant"`
}

type SequenceManager struct {
	Ctx context.Context
}

// AddDna Add Dna sequence and Mutant validation if not exist
func (m *SequenceManager) AddDna(data *Sequence, collName string) (generatedID string, err error) {
	//Validate Collention mongoDb
	coll, err := db.GetCollection(collName)
	if err != nil {
		return "", err
	}
	//Validate if DNA exist in DB
	var tempResult interface{}
	if _ = m.GetSequenceByItem(data.Dna, dnaBson, collectionMongo, &tempResult); tempResult != nil {
		return
	} else {
		resul, err := coll.InsertOne(m.Ctx, data)
		if err != nil {
			if strings.Contains(err.Error(), "dup key") {
				return "", errors.New("duplicated-Dna")
			}
			return "", err
		}
		generatedID, ok := resul.InsertedID.(string)
		if !ok {
			generatedID = resul.InsertedID.(primitive.ObjectID).Hex()
			if generatedID == "" {
				return "", errors.New("cannot-get-coll-id")
			}
		}
	}
	return
}

// GetSequenceByItem get one Sequence by it's item by nameItem.
func (m *SequenceManager) GetSequenceByItem(item interface{}, nameBson, collName string, resul interface{}) (err error) {
	coll, err := db.GetCollection(collName)
	if err != nil {
		return err
	}
	filter := make(map[string]interface{})
	filter[nameBson] = item
	err = coll.FindOne(context.TODO(), filter).Decode(resul)
	if err != mongo.ErrNoDocuments {
		return errors.New("Sequence-not-found-by-Dna")
	}
	return
}

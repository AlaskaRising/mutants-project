package models

type StatsResponse struct {
	Count_mutant_dna float64 `json:"count_mutant_dna" bson:"count_mutant_dna"`
	Count_human_dna  float64 `json:"count_human_dna" bson:"count_human_dna"`
	Ratio            float64 `json:"ratio" bson:"ratio"`
}

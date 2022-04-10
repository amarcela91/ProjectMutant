package models

type Mutant struct {
	Dna      []string `json:"dna" bson:"dna"`
	IsMutant bool     `json:"is_mutant" bson:"ismutant"`
}

type Stats struct {
	CountMutant int     `json:"count_mutant_dna"`
	CountHuman  int     `json:"count_human_dna"`
	Ratio       float64 `json:"ratio"`
}

package services

import (
	"ProjectMutant/internal/connection"
	"ProjectMutant/internal/models"
	"ProjectMutant/internal/utils"
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

func CreateMutant(newMutant models.Mutant) (models.Mutant, error) {
	newMutant.IsMutant = utils.IsMutant(newMutant.Dna)
	_, errorMongo := connection.GetCollection().InsertOne(context.TODO(), newMutant)
	if errorMongo != nil {
		return models.Mutant{}, errorMongo
	}
	return newMutant, nil
}

func CalculateStats() (models.Stats, error) {
	var results []*models.Mutant
	var count int
	stats := new(models.Stats)
	cur, errorMongo := connection.GetCollection().Find(context.TODO(), bson.D{{}})
	if errorMongo != nil {
		return models.Stats{}, errorMongo
	}

	for cur.Next(context.TODO()) {
		var elem models.Mutant
		if err := cur.Decode(&elem); err != nil {
			return models.Stats{}, err
		}
		if elem.IsMutant {
			count++
		}
		results = append(results, &elem)
	}

	stats.CountMutant = count
	stats.CountHuman = len(results)
	stats.Ratio = utils.ToFixed(float64(stats.CountMutant)/float64(stats.CountHuman), 2)

	return *stats, nil

}

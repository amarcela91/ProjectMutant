package services

import (
	"ProjectMutant/internal/connection"
	"ProjectMutant/internal/models"
	"ProjectMutant/internal/utils"
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
)

func CreateMutant(newMutant models.Mutant) (models.Mutant, error) {
	newMutant.IsMutant = utils.IsMutant(newMutant.Dna)
	findMutan := new(models.Mutant)
	err := connection.GetCollection().FindOne(context.TODO(), bson.D{{"dna", newMutant.Dna}}).Decode(&findMutan)

	//Si encuentra el mutante el err es nulo
	if err == nil {
		return models.Mutant{}, errors.New("El ADN que quiere analizar ya existe")
	}
	_, errorMongo := connection.GetCollection().InsertOne(context.TODO(), newMutant)
	if errorMongo != nil {
		//return models.Mutant{}, errorMongo
		return models.Mutant{}, errors.New("No se logró crear el mutante")
	}

	return newMutant, nil
}

func CalculateStats() (models.Stats, error) {
	var results []*models.Mutant
	var count int
	stats := new(models.Stats)
	cur, errorMongo := connection.GetCollection().Find(context.TODO(), bson.D{{}})
	if errorMongo != nil {
		return models.Stats{}, errors.New("No se logró calcular las estadísticas de verificación de ADN")
	}

	for cur.Next(context.TODO()) {
		var elem models.Mutant
		if err := cur.Decode(&elem); err != nil {
			return models.Stats{}, errors.New("No se logró calcular las estadísticas de verificación de ADN")
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

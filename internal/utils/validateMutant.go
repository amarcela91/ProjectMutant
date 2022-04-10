package utils

import (
	"strings"
)

func transformMatrix(data []string) [][]string {
	var result [][]string
	for i := 0; i < len(data); i++ {
		char := strings.Split(data[i], "")
		result = append(result, char)
	}
	return result
}

func countSequenceDNA(data [][]string) int {
	count := 0
	for i := 0; i < len(data[0]); i++ {
		for j := 0; j < len(data[0]); j++ {
			if j < 3 {
				//contador horizontal
				if data[i][j] == data[i][j+1] && data[i][j+1] == data[i][j+2] && data[i][j+2] == data[i][j+3] {
					count++
					if count >= 2 {
						return count
					}
				}
				//contador vertical
				if data[j][i] == data[j+1][i] && data[j+1][i] == data[j+2][i] && data[j+2][i] == data[j+3][i] {
					count++
					if count >= 2 {
						return count
					}
				}
			}
			if i < 3 && j < 3 {
				if data[j][i] == data[j+1][i+1] && data[j+1][i+1] == data[j+2][i+2] && data[j+2][i+2] == data[j+3][i+3] {
					count++
					if count >= 2 {
						return count
					}
				}
			}
		}
	}
	return count
}

func IsMutant(data []string) bool {
	count := countSequenceDNA(transformMatrix(data))
	return count >= 2
}

/*func main() {
	array := []string{
	"ATGCGA",
	"CAGGGC",
	"TCATGT",
	"TGCAGG",
	"TCCCCA",
	"TCACTG"}hay 6 coincidencias
	array := []string{"ATGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG"}
	array := []string{
		"ATGCGA",
		"CAGTGC",
		"TTATGT",
		"AGAAAG",
		"CCCTTA",
		"TCACTG"}
	fmt.Println(isMutant(array))
}*/

package knn

import (
	"math"
	"sort"
)

type Data struct {
	Age         float64 `json:"age,omitempty"`
	Cases       float64 `json:"cases,omitempty"`
	GenderClass string  `json:"gender_class,omitempty"`
	K           int     `json:"k,omitempty"`
}

func Predict(train_data [][]float64, testData []float64, k int) float64 {

	//in order to save the nearest neighbors, we calculate each distance
	//and make a map where key = distance, value = class it belongs to

	m := make(map[float64]float64)

	//looping over the dataset, we calculate the distance from our test point to every other point
	// row[2] is the class

	for _, row := range train_data {
		distance := euclideanDistance(row, testData)
		m[distance] = row[2]
	}

	//we make a new slice to save distances in order to sort them from lowest
	distances := make([]float64, 0, len(m))
	for dist := range m {
		distances = append(distances, dist)
	}
	//sort the distances
	//reorganize the distances starting from the smallest
	//so u can grap the closest neighboring
	//lower to highest

	sort.Float64Slice.Sort(distances)
	//slice is list in golang
	//initialize counters for neighbouring classes
	countF := 0
	countM := 0

	for i, d := range distances {
		//now that the distances are sorted from low to high
		//we loop over them for k iterations, finding the k closest neighbours
		//extract the classes belonging to those neighbors
		// 0 for male, 1 for female
		class := m[d]

		if class == 0 {
			countM += 1
		} else {
			countF += 1
		}
		//only for k times - because you want k closest neighbours
		if i == (k - 1) {
			break
		}
	}

	//if there were more female neighbours than male, return class for female
	//or the contrary
	if countF > countM {
		return 1
	} else {
		return 0
	}
}

func euclideanDistance(row1, row2 []float64) (distance float64) {
	//loop over all points of the vectors for each axis
	//in this case points are age_group and cases
	// (x1-x2)^2 + (y1-y2)^2 + ... and so on
	// square root on the total
	//towardsdatascience.com/knn-k-nearest-neighbors-1-a4707b24bd1d
	for i, _ := range row2 {
		distance += math.Pow((row1[i] - row2[i]), 2)
	}
	return math.Sqrt(distance)
}

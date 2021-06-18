package api

import (
	"encoding/json"
	"knn-golang/data"
	"knn-golang/knn"
	"log"
	"strconv"

	"net/http"
)

//incoming data from the frontend
type DataReq struct {
	Age         string `json:"age,omitempty"`
	Cases       string `json:"cases,omitempty"`
	GenderClass string `json:"gender_class,omitempty"`
	K           string `json:"k,omitempty"`
}

func Predict(w http.ResponseWriter, r *http.Request) {
	//set http headers
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	//default value of k
	k := 3

	//data struct
	testData := &knn.Data{}

	//parse JSON request
	dataReq := &DataReq{}
	_ = json.NewDecoder(r.Body).Decode(&dataReq)

	//incoming data will have to parsed from string
	age, err := strconv.ParseFloat(dataReq.Age, 64)
	if err != nil {
		log.Println(err)
	}

	//18 points will represent ages of people from 0 to 90
	// to scale incoming age multiply with 18/90
	testData.Age = age * 18 / 90

	//parse to float
	testData.Cases, err = strconv.ParseFloat(dataReq.Cases, 64)
	if err != nil {
		log.Println(err)
	}
	//parse to int
	testData.K, err = strconv.Atoi(dataReq.K)
	if err != nil {
		log.Println(err)
	}

	//add test data to a slice
	testRow := []float64{}
	testRow = append(testRow, testData.Age)
	testRow = append(testRow, testData.Cases)
	if testData.K != 0 {
		k = testData.K
	}

	//get the useful info from the dataset
	err, dataset := data.PreprocessData()
	if err != nil {
		log.Println(err)
	}
	//call predict to predict which gender the group belongs to
	class := knn.Predict(dataset, testRow, k)
	if class == 0 {
		testData.GenderClass = "Hombre"
	} else {
		testData.GenderClass = "Cocinera"
	}
	//return JSON encoded data
	json.NewEncoder(w).Encode(testData)
}

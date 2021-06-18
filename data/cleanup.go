package data

import (
	"encoding/csv"
	"errors"
	"io"
	"net/http"
	"os"
)

func PreprocessData() (err error, data [][]float64) {
	err = downloadDataset()
	if err != nil {
		return err, nil
	}
	err, data = cleanupDataset()
	if err != nil {
		return err, nil
	}
	return nil, data
}

func downloadDataset() error {
	//url to download from
	url := "https://raw.githubusercontent.com/Lutimi/dataset/master/pm_ener_2021.csv"
	//filename to save as
	filename := "pm_ener_2021.csv"
	//get from url
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	//check if downloaded ok
	if resp.StatusCode != 200 {
		return errors.New("received non 200 response code fetching dataset")
	}

	//create new file
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	//paste into it
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return err
	}
	defer file.Close()
	return nil
}

func cleanupDataset() (error, [][]float64) {
	//open the downloaded file
	csvFile, err := os.Open("pm_ener_2021.csv")
	if err != nil {
		return err, nil
	}
	defer csvFile.Close()

	//read from it
	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		return err, nil
	}

	//a map to count cases per each group
	// where key is numerical representation of age group, value is number of cases
	//
	//to turn age groups to numerical form:
	// 0-4 years --> 0
	// 5 - 9 --> 1
	// 10 - 13 --> 2
	//...
	// and so on
	//so that they keep the same incremental distance

	casesPerAgeGroupMale := make(map[float64]float64)
	casesPerAgeGroupFemale := make(map[float64]float64)

	for _, line := range csvLines {
		//if covid+
		if line[2] == "POSITIVO" {
			if line[4] == "Masculino" {
				switch line[3] {
				case "0-4años":
					casesPerAgeGroupMale[0] += 1
				case "5-9años":
					casesPerAgeGroupMale[1] += 1
				case "10-14años":
					casesPerAgeGroupMale[2] += 1
				case "15-19años":
					casesPerAgeGroupMale[3] += 1
				case "20-24años":
					casesPerAgeGroupMale[4] += 1
				case "25-29años":
					casesPerAgeGroupMale[5] += 1
				case "30-34años":
					casesPerAgeGroupMale[6] += 1
				case "35-39años":
					casesPerAgeGroupMale[7] += 1
				case "40-44años":
					casesPerAgeGroupMale[8] += 1
				case "45-49años":
					casesPerAgeGroupMale[9] += 1
				case "50-54años":
					casesPerAgeGroupMale[10] += 1
				case "55-59años":
					casesPerAgeGroupMale[11] += 1
				case "60-64años":
					casesPerAgeGroupMale[12] += 1
				case "65-69años":
					casesPerAgeGroupMale[13] += 1
				case "70-74años":
					casesPerAgeGroupMale[14] += 1
				case "75-79años":
					casesPerAgeGroupMale[15] += 1
				case "80-84años":
					casesPerAgeGroupMale[16] += 1
				case "85-89años":
					casesPerAgeGroupMale[17] += 1
				case "mas_de_90años":
					casesPerAgeGroupMale[18] += 1
				}
			} else {
				//if female
				switch line[3] {
				case "0-4años":
					casesPerAgeGroupFemale[0] += 1
				case "5-9años":
					casesPerAgeGroupFemale[1] += 1
				case "10-14años":
					casesPerAgeGroupFemale[2] += 1
				case "15-19años":
					casesPerAgeGroupFemale[3] += 1
				case "20-24años":
					casesPerAgeGroupFemale[4] += 1
				case "25-29años":
					casesPerAgeGroupFemale[5] += 1
				case "30-34años":
					casesPerAgeGroupFemale[6] += 1
				case "35-39años":
					casesPerAgeGroupFemale[7] += 1
				case "40-44años":
					casesPerAgeGroupFemale[8] += 1
				case "45-49años":
					casesPerAgeGroupFemale[9] += 1
				case "50-54años":
					casesPerAgeGroupFemale[10] += 1
				case "55-59años":
					casesPerAgeGroupFemale[11] += 1
				case "60-64años":
					casesPerAgeGroupFemale[12] += 1
				case "65-69años":
					casesPerAgeGroupFemale[13] += 1
				case "70-74años":
					casesPerAgeGroupFemale[14] += 1
				case "75-79años":
					casesPerAgeGroupFemale[15] += 1
				case "80-84años":
					casesPerAgeGroupFemale[16] += 1
				case "85-89años":
					casesPerAgeGroupFemale[17] += 1
				case "mas_de_90años":
					casesPerAgeGroupFemale[18] += 1
				}
			}
		}

	}

	totalData := [][]float64{}
	//create a 2D matrix
	//make a slice of [[age-group, cases, gender]...]

	for key, value := range casesPerAgeGroupMale {
		temp := []float64{}
		temp = append(temp, key)
		temp = append(temp, value)
		// class 0 for male
		temp = append(temp, 0)
		//create the slice [age-group, cases, male]
		//and add it to the slice of slices
		totalData = append(totalData, temp)
	}

	for key, value := range casesPerAgeGroupFemale {
		temp := []float64{}
		temp = append(temp, key)
		temp = append(temp, value)
		// class 1 for female
		temp = append(temp, 1)
		//create the slice [age-group, cases, female]
		//and add it to the slice of slices
		totalData = append(totalData, temp)
	}

	// now we have a 2D matrix
	// [[age, nr of cases, gender]
	// [age, nr of cases, gender]]
	return nil, totalData
}

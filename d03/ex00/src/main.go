package main

import (
	"encoding/csv"
	"encoding/json"
	// "fmt"
	"github.com/elastic/go-elasticsearch/v7"
	// "io/ioutil"
	"log"
	"os"
	"strings"
)

type RestorantData struct {
	Name     string    `json:"name"`
	Address  string    `json:"address"`
	Phone    string    `json:"phone"`
	Location GLocation `json:"location"`
}
type GLocation struct {
	Longitude string `json:"lon"`
	Latitude  string `json:"lat"`
}

func createRestorantData(data [][]string) []RestorantData {
	var restorantData []RestorantData

	for i, line := range data {
		if i > 0 {
			var rec RestorantData
			for j, field := range line {
				if j == 1 {
					rec.Name = field
				} else if j == 2 {
					rec.Address = field
				} else if j == 3 {
					rec.Phone = field
				} else if j == 4 {
					rec.Location.Longitude = field
				} else if j == 5 {
					rec.Location.Latitude = field
				}
			}
			restorantData = append(restorantData, rec)
		}
	}
	return restorantData
}

func makeJson() []byte {
	f, err := os.Open("../materials/dataMini.csv")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	// csvReader.FieldsPerRecord = -1
	csvReader.Comma = '\t'
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(data[15])

	restorantData := createRestorantData(data)

	jsonData, err := json.MarshalIndent(restorantData, "", "    ")
	if err != nil {
		log.Fatal(err)
	}

	// err = ioutil.WriteFile("test.json", jsonData, 0644)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	return jsonData

}

func main() {

	client, err := elasticsearch.NewDefaultClient()
	if err != nil {
		log.Fatal(err)
	}
	index := "placesss"
	// schema, err := os.ReadFile("schema.json")

	mapping := `
    {
      "settings": {
        "number_of_shards": 1
      },
      "mappings": {
		"properties": {
		  "name": {
			  "type":  "text"
		  },
		  "address": {
			  "type":  "text"
		  },
		  "phone": {
			  "type":  "text"
		  },
		  "location": {
			"type": "geo_point"
		  }
		}
	  }
    }`

	res, err := client.Indices.Create(
		index,
		client.Indices.Create.WithBody(strings.NewReader(mapping)),
	)
	if err != nil {
		log.Fatal(err)
	}
	// res = nil
	log.Println(res)
	// jsonData := makeJson()

	// byteSlice, err := ioutil.ReadAll(jsonData)

}

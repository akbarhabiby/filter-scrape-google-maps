package helpers

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/akbarhabiby/filter-scrape-google-maps/models"
)

func ExportJSONtoCSV(fileName string, jsonData interface{}) {
	fmt.Printf("Exporting %s ...\n", fileName)
	timeLog := Timelog(fmt.Sprintf("%s %s", "Export", fileName))

	rows, ok := jsonData.([]models.Scrapes)
	if !ok {
		panic("Failed to parse jsonData to rows")
	}

	file, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	if err := writer.Write([]string{"full_address", "number", "district", "city", "province", "postal_code", "country", "latitude", "longitude", "plus_code"}); err != nil {
		panic(err)
	}

	for _, row := range rows {
		if err := writer.Write([]string{row.FullAddress, row.Number, row.District, row.City, row.Province, row.PostalCode, row.Country, row.Latitude, row.Longitude, row.PlusCode}); err != nil {
			panic(err)
		}
	}

	timeLog()
}

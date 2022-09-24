package helpers

import (
	"encoding/csv"
	"fmt"
	"math"
	"os"
	"strings"

	"github.com/akbarhabiby/filter-scrape-google-maps/models"
)

func ExportJSONtoCSV(fileName string, jsonData interface{}) {
	rows, ok := jsonData.([]models.Scrapes)
	if !ok {
		panic("Failed to parse jsonData to rows")
	}

	maxLoop := int(1e+6) // * For handling MS Excel max rows
	totalLoop := 0
	rowsLength := len(rows)

	for i := 0; i < int(math.Ceil(float64(rowsLength)/float64(maxLoop))); i++ {
		fName := fmt.Sprintf("%s-%v.csv", strings.ReplaceAll(fileName, ".csv", ""), i)
		fmt.Printf("Exporting %s ...\n", fName)
		timeLog := Timelog(fmt.Sprintf("%s %s", "Export", fName))
		tempLoop := 0
		file, err := os.Create(fName)
		if err != nil {
			panic(err)
		}
		defer file.Close()

		writer := csv.NewWriter(file)
		defer writer.Flush()

		if err := writer.Write([]string{"full_address", "number", "district", "city", "province", "postal_code", "country", "latitude", "longitude", "plus_code"}); err != nil {
			panic(err)
		}

		for {
			if err := writer.Write([]string{rows[totalLoop].FullAddress, rows[totalLoop].Number, rows[totalLoop].District, rows[totalLoop].City, rows[totalLoop].Province, rows[totalLoop].PostalCode, rows[totalLoop].Country, rows[totalLoop].Latitude, rows[totalLoop].Longitude, rows[totalLoop].PlusCode}); err != nil {
				panic(err)
			}
			tempLoop++
			totalLoop++
			if tempLoop == maxLoop || totalLoop == rowsLength {
				timeLog()
				break
			}
		}
	}
}

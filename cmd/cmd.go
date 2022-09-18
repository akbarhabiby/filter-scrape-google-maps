package cmd

import (
	"fmt"
	"log"
	"path"
	"strconv"
	"strings"

	"github.com/akbarhabiby/filter-scrape-google-maps/constants"
	"github.com/akbarhabiby/filter-scrape-google-maps/helpers"
	"github.com/akbarhabiby/filter-scrape-google-maps/models"
)

func Run(fileName string) (total int, success int, failed int) {
	var result models.Result
	helpers.ImportJSON(fileName, &result)

	listSuccess := make([]models.Scrapes, 0)
	listFail := make([]models.Scrapes, 0)

	fmt.Printf("Working on %s ...\n", fileName)
	timeLog := helpers.Timelog("Work")
	for _, address := range result.Scrapes {
		number := strings.TrimSpace(address.Number)
		if strings.Contains(number, "-") {
			numbersArr := strings.Split(number, "-")

			x, err := helpers.GetInt(numbersArr[0])
			if err != nil {
				newNumber := helpers.RemoveUnusedStringAndToUpper(address.Number)
				listFail = append(listFail, helpers.GenerateNewAddressModel(address, newNumber))
				log.Println(fmt.Errorf("failed to parse [%s] => [%s]", number, newNumber))
				continue
			}

			y, err := helpers.GetInt(numbersArr[1])
			if err != nil {
				newNumber := helpers.RemoveUnusedStringAndToUpper(address.Number)
				listFail = append(listFail, helpers.GenerateNewAddressModel(address, newNumber))
				log.Println(fmt.Errorf("failed to parse [%s] => [%s]", number, newNumber))
				continue
			}

			if y < x {
				z := y
				y = x
				x = z
			}

			for i := x; i <= y; i++ {
				listSuccess = append(listSuccess, helpers.GenerateNewAddressModel(address, fmt.Sprint(i)))
			}
		} else {
			newNumber := helpers.RemoveUnusedStringAndToUpper(address.Number)
			i, err := strconv.Atoi(newNumber)
			if err == nil {
				newNumber = fmt.Sprint(i)
			}
			listSuccess = append(listSuccess, helpers.GenerateNewAddressModel(address, newNumber))
		}
	}

	timeLog()

	jsonCSVData := make([]models.Scrapes, 0)
	jsonCSVData = append(jsonCSVData, listSuccess...)
	jsonCSVData = append(jsonCSVData, listFail...)
	helpers.ExportJSONtoCSV(path.Join(constants.OUTPUT_DIR, strings.ReplaceAll(strings.ReplaceAll(fileName, "input/", ""), ".json", ".csv")), jsonCSVData)

	total = len(result.Scrapes)
	success = len(listSuccess)
	failed = len(listFail)
	return
}

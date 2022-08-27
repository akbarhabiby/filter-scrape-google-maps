package cmd

import (
	"fmt"
	"log"
	"os"
	"path"
	"strings"

	"github.com/akbarhabiby/filter-scrape-google-maps/constants"
	"github.com/akbarhabiby/filter-scrape-google-maps/helpers"
	"github.com/akbarhabiby/filter-scrape-google-maps/models"
)

func init() {
	if err := os.Mkdir(constants.LOG_DIR, os.ModePerm); err == nil {
		fmt.Printf("[INIT] %s folder not found, created.\n", constants.LOG_DIR)
	}
	file, err := os.OpenFile(path.Join(constants.LOG_DIR, constants.LOG_FILE), os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}
	log.SetOutput(file)
}

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
				listFail = append(listFail, address)
				log.Println(fmt.Errorf("failed to parse [%s]", number))
				continue
			}

			y, err := helpers.GetInt(numbersArr[1])
			if err != nil {
				listFail = append(listFail, address)
				log.Println(fmt.Errorf("failed to parse [%s]", number))
				continue
			}

			if y < x {
				z := y
				y = x
				x = z
			}

			numReplacer := address.Number

			for i := x; i <= y; i++ {
				newNumber := fmt.Sprint(i)
				newFullAddress := strings.ReplaceAll(address.FullAddress, numReplacer, newNumber)
				newAddress := models.Scrapes{
					FullAddress: newFullAddress,
					Number:      newNumber,
					District:    address.District,
					City:        address.City,
					Province:    address.Province,
					PostalCode:  address.PostalCode,
					Country:     address.Country,
					Latitude:    address.Latitude,
					Longitude:   address.Longitude,
					PlusCode:    address.PlusCode,
					CreatedAt:   address.CreatedAt,
				}
				listSuccess = append(listSuccess, newAddress)
			}
		} else {
			listSuccess = append(listSuccess, address)
		}
	}

	timeLog()

	if err := os.Mkdir(constants.OUTPUT_DIR, os.ModePerm); err == nil {
		fmt.Printf("[INIT] %s folder not found, created.\n", constants.OUTPUT_DIR)
	}
	helpers.ExportJSON(path.Join(constants.OUTPUT_DIR, fmt.Sprintf("%s-%s", constants.SUCCESS, strings.ReplaceAll(fileName, "input/", ""))), listSuccess)
	helpers.ExportJSON(path.Join(constants.OUTPUT_DIR, fmt.Sprintf("%s-%s", constants.ERROR, strings.ReplaceAll(fileName, "input/", ""))), listFail)

	total = len(result.Scrapes)
	success = len(listSuccess)
	failed = len(listFail)
	return
}

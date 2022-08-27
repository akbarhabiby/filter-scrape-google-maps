package helpers

import (
	"strings"

	"github.com/akbarhabiby/filter-scrape-google-maps/models"
)

func GenerateNewAddressModel(address models.Scrapes, newNumber string) (newAddress models.Scrapes) {
	newFullAddress := strings.ReplaceAll(address.FullAddress, address.Number, newNumber)
	newAddress = models.Scrapes{
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
	return
}

package main

import (
	"fmt"
	"os"
	"path"
	"runtime"
	"strings"
	"sync"

	"github.com/akbarhabiby/filter-scrape-google-maps/cmd"
	"github.com/akbarhabiby/filter-scrape-google-maps/constants"
	"github.com/akbarhabiby/filter-scrape-google-maps/helpers"
)

func main() {
	defer helpers.Timelog("Export")()
	runtime.GOMAXPROCS(runtime.NumCPU())

	if err := os.Mkdir(constants.INPUT_DIR, os.ModePerm); err == nil {
		fmt.Printf("[INIT] %s folder not found, created.\n", constants.INPUT_DIR)
	}
	dir, err := os.ReadDir(constants.INPUT_DIR)
	if err != nil {
		panic(err)
	}

	files := make([]string, 0)

	for _, file := range dir {
		if !file.IsDir() {
			item := strings.Split(file.Name(), ".")
			if strings.ToLower(item[len(item)-1]) == "json" {
				files = append(files, path.Join(constants.INPUT_DIR, file.Name()))
			}
		}
	}

	var wg sync.WaitGroup

	wg.Add(len(files))

	for _, file := range files {
		go func(fileName string) {
			defer wg.Done()
			total, success, failed := cmd.Run(fileName)
			fmt.Printf("Total [%v] | Success [%v] | Failed [%v]\n", total, success, failed)
		}(file)
	}

	wg.Wait()
}

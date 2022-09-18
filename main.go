package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"runtime"
	"strings"
	"sync"

	"github.com/akbarhabiby/filter-scrape-google-maps/cmd"
	"github.com/akbarhabiby/filter-scrape-google-maps/constants"
	"github.com/akbarhabiby/filter-scrape-google-maps/helpers"
)

const banner = `
  _____.__.__   __                                                           
_/ ____|__|  |_/  |_  ___________    ______ ________________  ______   ____  
\   __\|  |  |\   ___/ __ \_  __ \  /  ____/ ___\_  __ \__  \ \____ \_/ __ \ 
 |  |  |  |  |_|  | \  ___/|  | \/  \___ \\  \___|  | \// __ \|  |_> \  ___/ 
 |__|  |__|____|__|  \___  |__|    /____  >\___  |__|  (____  |   __/ \___  >
                         \/             \/     \/           \/|__|        \/ 
                                                                             
`

func init() {
	if err := os.Mkdir(constants.LOG_DIR, os.ModePerm); err == nil {
		fmt.Printf("[INIT] %s folder not found, created.\n", constants.LOG_DIR)
	}
	if err := os.Mkdir(constants.OUTPUT_DIR, os.ModePerm); err == nil {
		fmt.Printf("[INIT] %s folder not found, created.\n", constants.OUTPUT_DIR)
	}
	if err := os.Mkdir(constants.INPUT_DIR, os.ModePerm); err == nil {
		fmt.Printf("[INIT] %s folder not found, created.\n", constants.INPUT_DIR)
	}
	file, err := os.OpenFile(path.Join(constants.LOG_DIR, constants.LOG_FILE), os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}
	log.SetOutput(file)

	file, err = os.Open(constants.STR_REPLACE_FILE)
	if err != nil {
		file, err = os.Create(constants.STR_REPLACE_FILE)
		if err != nil {
			panic(err)
		}
		fmt.Printf("[INIT] %s file not found, created.\n", constants.STR_REPLACE_FILE)
		defer file.Close()
		err = json.Unmarshal([]byte(constants.DEFAULT_REPLACER), &helpers.ReplacerStrings)
		if err != nil {
			panic(err)
		}
		bt, err := json.MarshalIndent(helpers.ReplacerStrings, "", "  ")
		if err != nil {
			panic(err)
		}
		file.Write(bt)
		return
	}

	defer file.Close()
	bt, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(bt, &helpers.ReplacerStrings)
	if err != nil {
		panic(err)
	}
}

func main() {
	fmt.Print(banner)
	timeLog := helpers.Timelog("Export")
	runtime.GOMAXPROCS(runtime.NumCPU())

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
			fmt.Printf("\n\n>> Input   : %s\n>> Total   : %v\n>> Success : %v\n>> Failed  : %v\n\n", fileName, total, success, failed)
		}(file)
	}

	wg.Wait()
	timeLog()

	fmt.Println("Press 'Enter' to exit")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}

package config

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// our loaded configurations
var configs = make(map[string]string)

// load configuration for one or more files,
// providing file names is optional, if not provided it will
// look for at file called [executable_name].ini
func Load(requiredCondfigs []string, fileNames ...string) {

	//fmt.Println(executable)

	if len(fileNames) == 0 {
		executable := filepath.Base(os.Args[0])
		fileNames = []string{executable + ".ini"}
	}

	for _, fileName := range fileNames {
		parse(fileName, requiredCondfigs)
	}
}

func parse(fileName string, requiredCondfigs []string) {
	file, err := os.Open(fileName) // from "current" dir?
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		text := strings.Trim(scanner.Text(), " ")
		if strings.HasPrefix(text, ";") || strings.HasPrefix(text, "[") {
			continue
		}
		qIndex := strings.Index(text, "=")
		if qIndex > -1 {
			key := strings.Trim(string(text[:qIndex]), " ")
			value := strings.Trim(string(text[qIndex:]), "= ")
			if key != "" && value != "" {
				configs[key] = value
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Print(err)
	}

	// check for *required* configs!
	for _, conf := range requiredCondfigs {
		_, exists := configs[conf]
		if !exists {
			log.Fatal("Missing **required** configuration: " + conf)
		}

	}
}

func Get(key string) string {
	return configs[key]
}

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/ez-connect/go-rest/core"
	"github.com/ez-connect/go-rest/generator"
)

func main() {
	workingDir := "example"
	fmt.Println("Working dir:", workingDir)

	dirs, err := ioutil.ReadDir(fmt.Sprintf("%s/routes", workingDir))
	if err != nil {
		log.Fatal(err)
	}

	for _, dir := range dirs {
		if dir.IsDir() {
			fmt.Println(workingDir, "/routes/", dir.Name())
			config := generator.Config{}
			err := core.LoadConfig(fmt.Sprintf("%s/routes/%s/settings.yml", workingDir, dir.Name()), &config)
			if err != nil {
				log.Fatal(err)
			}

			err = os.MkdirAll(fmt.Sprintf("%s/generated/%s", workingDir, dir.Name()), os.ModeDir)
			if err != nil {
				log.Fatal(err)
			}

			// Generate
			generator.GenerateFile(workingDir, dir.Name(), "model", config)
			generator.GenerateFile(workingDir, dir.Name(), "repository", config)
			generator.GenerateFile(workingDir, dir.Name(), "handler", config)
			generator.GenerateFile(workingDir, dir.Name(), "router", config)

		}
	}
}

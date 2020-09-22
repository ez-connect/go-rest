package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/ez-connect/go-rest/core"
	"github.com/ez-connect/go-rest/generator"
)

func main() {
	workingDir := "example/routes"
	fmt.Println("Working dir:", workingDir)

	files, err := ioutil.ReadDir(workingDir)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		if f.IsDir() {
			fmt.Println(f.Name())
			c := generator.ModelConfig{}
			err := core.LoadConfig(fmt.Sprintf("%s/%s/settings.yml", workingDir, f.Name()), &c)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(c)
		}
	}
}

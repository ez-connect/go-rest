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

	files, err := ioutil.ReadDir(fmt.Sprintf("%s/routes", workingDir))
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		if f.IsDir() {
			fmt.Println(f.Name())
			c := generator.ModelConfig{}
			err := core.LoadConfig(fmt.Sprintf("%s/routes/%s/settings.yml", workingDir, f.Name()), &c)
			if err != nil {
				log.Fatal(err)
			}

			os.MkdirAll(fmt.Sprintf("%s/generated/%s", workingDir, f.Name()), os.ModeDir)

			// Model
			f, err := os.Create(fmt.Sprintf("%s/generated/%s/model.go", workingDir, f.Name()))
			if err != nil {
				log.Fatal(err)
			}

			defer f.Close()
			v := generator.GenerateModel(f.Name(), c)
			_, err = f.WriteString(v)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(c)
		}
	}
}

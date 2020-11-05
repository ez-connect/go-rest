package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/ez-connect/go-rest/cli/generator"
	"github.com/ez-connect/go-rest/core"
)

func main() {
	dir := flag.String("dir", ".", "Working dir")
	new := flag.String("new", "", "Create new service")
	flag.Parse()

	workingDir := *dir
	fmt.Println("Working dir:", workingDir)

	// New service generator
	service := *new
	fmt.Println(service)
	if service != "" {
		err := os.MkdirAll(fmt.Sprintf("%s/services/%s", workingDir, service), os.ModeDir)
		if err != nil {
			log.Fatal(err)
		}

		// Generate
		generator.WriteService(workingDir, service, generator.Settings)
		generator.WriteService(workingDir, service, generator.Model)
		generator.WriteService(workingDir, service, generator.Repository)
		generator.WriteService(workingDir, service, generator.Handler)
		generator.WriteService(workingDir, service, generator.Router)

		return
	}

	// Generate source from services
	dirs, err := ioutil.ReadDir(fmt.Sprintf("%s/services", workingDir))
	if err != nil {
		log.Fatal(err)
	}

	for _, dir := range dirs {
		if dir.IsDir() && strings.Index(dir.Name(), "_") != 0 {
			fmt.Println(workingDir, "/services/", dir.Name())
			config := generator.Config{}
			err := core.LoadConfig(fmt.Sprintf("%s/services/%s/%v", workingDir, dir.Name(), generator.Settings), &config)
			if err != nil {
				log.Fatal(err)
			}

			err = os.MkdirAll(fmt.Sprintf("%s/generated/%s", workingDir, dir.Name()), os.ModeDir)
			if err != nil {
				log.Fatal(err)
			}

			// Generate
			generator.WriteSource(workingDir, dir.Name(), generator.Model, config)
			generator.WriteSource(workingDir, dir.Name(), generator.Repository, config)
			generator.WriteSource(workingDir, dir.Name(), generator.Handler, config)
			generator.WriteSource(workingDir, dir.Name(), generator.Router, config)
		}
	}

	// // Format code
	// cmd := exec.Command("golangci-lint", "run", "--fix", "--disable-all", "--enable", "goimports", workingDir+"/...")
	// var er bytes.Buffer
	// cmd.Stderr = &er
	// if err = cmd.Run(); err != nil {
	// 	fmt.Println("lint error: ", er.String())
	// }

	fmt.Println("All services are generated")
}

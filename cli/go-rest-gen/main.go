package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/ez-connect/go-rest/cli/go-rest-gen/gen"
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
		gen.WriteService(workingDir, service, gen.Settings)
		gen.WriteService(workingDir, service, gen.Model)
		gen.WriteService(workingDir, service, gen.Repository)
		gen.WriteService(workingDir, service, gen.Handler)
		gen.WriteService(workingDir, service, gen.Router)

		return
	}

	// Generate source from services
	dirs, err := ioutil.ReadDir(fmt.Sprintf("%s/services", workingDir))
	if err != nil {
		log.Fatal(err)
	}

	for _, dir := range dirs {
		if !dir.IsDir() {
			continue
		}

		fmt.Println(workingDir, "/services/", dir.Name())
		config := gen.Config{}
		filename := fmt.Sprintf("%s/services/%s/%v", workingDir, dir.Name(), gen.Settings)

		// Check for settings exists
		if _, err := os.Stat(filename); os.IsNotExist(err) {
			fmt.Println(fmt.Sprintf("No %s file found", gen.Settings))
			continue
		}

		// Generate files
		fmt.Println("Load", filename)
		err := core.LoadConfig(filename, &config)
		if err != nil {
			log.Fatal(err)
		}

		err = os.MkdirAll(fmt.Sprintf("%s/generated/%s", workingDir, dir.Name()), os.ModeDir)
		if err != nil {
			log.Fatal(err)
		}

		// Generate
		gen.WriteSource(workingDir, dir.Name(), gen.Model, config)
		gen.WriteSource(workingDir, dir.Name(), gen.Repository, config)
		gen.WriteSource(workingDir, dir.Name(), gen.Handler, config)
		gen.WriteSource(workingDir, dir.Name(), gen.Router, config)
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

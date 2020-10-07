package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/ez-connect/go-rest/cli/generator"
	"github.com/ez-connect/go-rest/core"
)

func main() {
	dir := flag.String("dir", ".", "Working dir")
	init := flag.String("init", "n", "Init generator")
	new := flag.String("new", "", "Create new service")
	flag.Parse()

	workingDir := *dir
	fmt.Println("Working dir:", workingDir)

	// Init generator
	if *init == "y" {
		err := os.MkdirAll(fmt.Sprintf("%s/services/_base", workingDir), os.ModeDir)
		if err != nil {
			log.Fatal(err)
		}

		generator.GenerateBase(workingDir, "handler")
		generator.GenerateBase(workingDir, "repository")

		return
	}

	// New service generator
	service := *new
	fmt.Println(service)
	if service != "" {
		err := os.MkdirAll(fmt.Sprintf("%s/services/%s", workingDir, service), os.ModeDir)
		if err != nil {
			log.Fatal(err)
		}

		// Generate
		generator.GenerateFileExt(workingDir, service, "settings.yml")
		generator.GenerateFileExt(workingDir, service, "model.go")
		generator.GenerateFileExt(workingDir, service, "repository.go")
		generator.GenerateFileExt(workingDir, service, "handler.go")
		generator.GenerateFileExt(workingDir, service, "router.go")

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
			err := core.LoadConfig(fmt.Sprintf("%s/services/%s/settings.yml", workingDir, dir.Name()), &config)
			if err != nil {
				log.Fatal(err)
			}

			err = os.MkdirAll(fmt.Sprintf("%s/generated/%s", workingDir, dir.Name()), os.ModeDir)
			if err != nil {
				log.Fatal(err)
			}

			// Generate
			generator.GenerateFile(workingDir, dir.Name(), "model.go", config)
			generator.GenerateFile(workingDir, dir.Name(), "repository.go", config)
			generator.GenerateFile(workingDir, dir.Name(), "handler.go", config)
			generator.GenerateFile(workingDir, dir.Name(), "router.go", config)
		}
	}

	// Format code
	cmd := exec.Command("golangci-lint", "run", "--fix", "--disable-all", "--enable", "goimports", workingDir+"/...")
	var er bytes.Buffer
	cmd.Stderr = &er
	if err = cmd.Run(); err != nil {
		fmt.Println("lint error: ", er.String())
	}
}

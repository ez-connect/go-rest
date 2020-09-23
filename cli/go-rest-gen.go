package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"

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
	if service != "" {
		err := os.MkdirAll(fmt.Sprintf("%s/services/%s", workingDir, service), os.ModeDir)
		if err != nil {
			log.Fatal(err)
		}

	}

	dirs, err := ioutil.ReadDir(fmt.Sprintf("%s/services", workingDir))
	if err != nil {
		log.Fatal(err)
	}

	for _, dir := range dirs {
		if dir.IsDir() {
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
			generator.GenerateFile(workingDir, dir.Name(), "model", config)
			generator.GenerateFile(workingDir, dir.Name(), "repository", config)
			generator.GenerateFile(workingDir, dir.Name(), "handler", config)
			generator.GenerateFile(workingDir, dir.Name(), "router", config)
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

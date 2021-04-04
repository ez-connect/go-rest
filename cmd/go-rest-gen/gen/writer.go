package gen

import (
	"fmt"
	"log"
	"os"
)

type FileType string

const (
	Settings FileType = "settings.yml"

	Model      FileType = "model.go"
	Repository FileType = "repository.go"
	Handler    FileType = "handler.go"
	Router     FileType = "router.go"
)

func WriteSource(workingDir, packageName string, fileType FileType, config Config) {
	var v string
	switch fileType {
	case Model:
		v = GenerateModel(packageName, config)
	case Repository:
		v = GenerateRepository(packageName, config)
	case Handler:
		v = GenerateHandler(packageName, config)
	case Router:
		v = GenerateRoutes(packageName, config)
	default:
		log.Fatal("Not support type:", fileType)
	}

	filename := fmt.Sprintf("%s/generated/%s/%s", workingDir, packageName, fileType)
	fmt.Println(filename)

	f, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	_, err = f.WriteString(v)
	if err != nil {
		log.Fatal(err)
	}
}

func WriteService(workingDir, packageName string, fileType FileType) {
	var v string
	switch fileType {
	case Settings:
		v = GenerateSettings(packageName)
	case Model:
		v = GenerateModelService(packageName)
	case Repository:
		v = GenerateRepositoryService(packageName)
	case Handler:
		v = GenerateHandlerService(packageName)
	case Router:
		v = GenerateRoutesService(packageName)
	default:
		log.Fatal("Not support type:", fileType)
	}

	filename := fmt.Sprintf("%s/services/%s/%s", workingDir, packageName, fileType)
	fmt.Println(filename)

	f, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	_, err = f.WriteString(v)
	if err != nil {
		log.Fatal(err)
	}
}

func WriteConstants(workingDir string, configs []Config) {
	filename := fmt.Sprintf("%s/generated/constants.go", workingDir)
	fmt.Println(filename)

	f, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	_, err = f.WriteString(GenerateConstants(configs))
	if err != nil {
		log.Fatal(err)
	}
}

func WritePolicy(workingDir string, configs []Config) {
	filename := fmt.Sprintf("%s/generated/policy.csv", workingDir)
	fmt.Println(filename)

	f, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	_, err = f.WriteString(GeneratePolicy(configs))
	if err != nil {
		log.Fatal(err)
	}
}

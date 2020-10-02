package generator

import (
	"fmt"
	"log"
	"os"
)

func GenerateBase(workingDir, fileType string) {
	var v string
	switch fileType {
	case "handler":
		v = GenerateBaseHandler()
	case "repository":
		v = GenerateBaseRepository()
	default:
		log.Fatal("Not support type:", fileType)
	}

	filename := fmt.Sprintf("%s/services/_base/%s.go", workingDir, fileType)
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

func GenerateFile(workingDir, packageName, fileType string, config Config) {
	var v string
	switch fileType {
	case "model.go":
		v = GenerateModel(packageName, config)
	case "repository.go":
		v = GenerateRepository(packageName, config)
	case "handler.go":
		v = GenerateHandler(packageName)
	case "router.go":
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

func GenerateFileExt(workingDir, packageName, fileType string) {
	var v string
	switch fileType {
	case "settings.yml":
		v = GenerateSettings(packageName)
	case "model.go":
		v = GenerateModelExt(packageName)
	case "repository.go":
		v = GenerateRepositoryExt(packageName)
	case "handler.go":
		v = GenerateHandlerExt(packageName)
	case "router.go":
		v = GenerateRoutesExt(packageName)
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

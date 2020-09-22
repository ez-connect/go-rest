package generator

import (
	"fmt"
	"log"
	"os"
)

func GenerateFile(workingDir, packageName, fileType string, config Config) {
	var v string
	switch fileType {
	case "model":
		v = GenerateModel(packageName, config.Model)
	case "repository":
		v = GenerateRepository(packageName, config.Model.Name)
	case "handler":
		v = GenerateHandler(packageName)
	case "router":
		v = GenerateRoutes(packageName, config.Model.Name, config.Routes)
	default:
		log.Fatal("Not support type:", fileType)
	}

	filename := fmt.Sprintf("%s/generated/%s/%s.go", workingDir, packageName, fileType)
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

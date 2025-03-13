package generator

import (
	"bytes"
	"fmt"
	"os"
	"text/template"
)

func GenerateConfig() {
	tmpl, err := template.ParseFiles("./internal/template/goarch.yaml.example.tmpl")
	if err != nil {
		panic(err)
	}

	var buf bytes.Buffer
	err = tmpl.Execute(&buf, nil)
	if err != nil {
		panic(err)
	}

	filename := "./goarch.yaml"

	if _, err := os.Stat(filename); err == nil {
		fmt.Printf("Error: %s already exists. Please delete it before generating a new config.\n", filename)
		return
	} else if !os.IsNotExist(err) {
		panic(err)
	}

	content := buf.String()
	writeFile(filename, content)
}

package generator

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/alramdein/goarch/internal/config"
	"gopkg.in/yaml.v2"
)

type VariableMapping struct {
	EntityLower string
	EntityUpper string
	LayerLower  string
	LayerUpper  string
}

func GenerateProject() {
	config := loadConfig("goarch.yaml")
	generatedFiles := loadGeneratedFiles(config.OutputPath)

	for _, entities := range config.Content.Entities {
		for e, entityCfg := range entities {
			entity := strings.ToLower(e)
			generateModel(config.OutputPath, entity, generatedFiles)

			for _, method := range entityCfg.Methods {
				if method.Name == "default_crud" {
					for _, layer := range config.Content.Layers {
						generateLayer(config.OutputPath, entity, layer, generatedFiles)
					}
				}
			}
		}

	}
	saveGeneratedFiles(config.OutputPath, generatedFiles)
}

func loadConfig(path string) config.Config {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	var config config.Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		panic(err)
	}
	return config
}

func loadGeneratedFiles(outputPath string) map[string]bool {
	files := make(map[string]bool)
	trackFile := filepath.Join(outputPath, ".goarch.lock")
	data, err := os.ReadFile(trackFile)
	if err == nil {
		for _, line := range strings.Split(string(data), "\n") {
			if line != "" {
				files[line] = true
			}
		}
	}
	return files
}

func saveGeneratedFiles(outputPath string, files map[string]bool) {
	trackFile := filepath.Join(outputPath, ".goarch.lock")
	var content []string
	for file := range files {
		content = append(content, file)
	}
	os.WriteFile(trackFile, []byte(strings.Join(content, "\n")), 0644)
}

func generateModel(outputPath, entity string, generatedFiles map[string]bool) {
	filename := filepath.Join(outputPath, "model", fmt.Sprintf("%s.go", entity))
	if !generatedFiles[filename] {
		writeFile(filename, fmt.Sprintf("package model\n\ntype %s struct {\n\tID int\n}\n", strings.Title(entity)))
		generatedFiles[filename] = true
	}
}

func generateLayer(outputPath, entity string, layer string, generatedFiles map[string]bool) {
	filename := filepath.Join(outputPath, layer, fmt.Sprintf("%s_%s.go", entity, layer))
	if generatedFiles[filename] {
		return
	}

	tmpl, err := template.ParseFiles("./internal/template/common.layer.tmpl")
	if err != nil {
		panic(err)
	}

	variableMapping := VariableMapping{
		EntityLower: strings.ToLower(entity),
		EntityUpper: capitalizeFirstLetter(entity),
		LayerLower:  strings.ToLower(layer),
		LayerUpper:  capitalizeFirstLetter(layer),
	}

	var buf bytes.Buffer
	err = tmpl.Execute(&buf, variableMapping)
	if err != nil {
		panic(err)
	}

	content := buf.String()

	writeFile(filename, content)
	generatedFiles[filename] = true
}

func writeFile(filename, content string) {
	os.MkdirAll(filepath.Dir(filename), os.ModePerm)
	os.WriteFile(filename, []byte(content), 0644)
}

func capitalizeFirstLetter(s string) string {
	if len(s) > 0 {
		return strings.ToUpper(string(s[0])) + s[1:]
	}
	return ""
}

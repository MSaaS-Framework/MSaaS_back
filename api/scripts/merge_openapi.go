package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

// OpenAPI 스펙 구조 정의
type OpenAPI struct {
	OpenAPI string                 `yaml:"openapi"`
	Info    map[string]interface{} `yaml:"info"`
	Servers []map[string]string    `yaml:"servers"`
	Paths   map[string]interface{} `yaml:"paths"`
}

func main() {
	// 병합된 OpenAPI 구조 초기화
	mergedSpec := OpenAPI{
		OpenAPI: "3.0.3",
		Info:    map[string]interface{}{"title": "MSaaS API", "version": "1.0.0"},
		Servers: []map[string]string{{"url": "http://localhost:8080"}},
		Paths:   make(map[string]interface{}),
	}

	// `api-docs` 폴더 내 모든 .yaml 파일 병합
	err := filepath.Walk("./api-docs", func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(path) == ".yaml" {
			data, _ := ioutil.ReadFile(path)
			var spec OpenAPI
			_ = yaml.Unmarshal(data, &spec)
			for key, value := range spec.Paths {
				mergedSpec.Paths[key] = value
			}
		}
		return nil
	})
	if err != nil {
		panic(err)
	}

	// 병합된 파일 저장
	mergedData, _ := yaml.Marshal(mergedSpec)
	_ = ioutil.WriteFile("./openapi.yaml", mergedData, 0644)

	fmt.Println("✅ OpenAPI 문서 병합 완료: openapi.yaml")
}

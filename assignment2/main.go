package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type TokenizerConfig struct {
	TokenChars []string `json:"token_chars"`
	MinGram    string   `json:"min_gram"`
	Type       string   `json:"type"`
	MaxGram    string   `json:"max_gram"`
}

type AnalyzerConfig struct {
	Filter    []string        `json:"filter"`
	Tokenizer TokenizerConfig `json:"tokenizer"`
}

type NormalizerConfig struct {
	CaseInsensitive struct {
		Filter     []string      `json:"filter"`
		Type       string        `json:"type"`
		CharFilter []interface{} `json:"char_filter"`
	} `json:"case_insensitive"`
}

type AnalysisConfig struct {
	Normalizer NormalizerConfig `json:"normalizer"`
	Analyzer   AnalyzerConfig   `json:"analyzer"`
	Tokenizer  TokenizerConfig  `json:"tokenizer"`
}

type IndexConfig struct {
	RefreshInterval  string         `json:"refresh_interval"`
	NumberOfShards   string         `json:"number_of_shards"`
	ProvidedName     string         `json:"provided_name"`
	CreationDate     string         `json:"creation_date"`
	Analysis         AnalysisConfig `json:"analysis"`
	NumberOfReplicas string         `json:"number_of_replicas"`
	UUID             string         `json:"uuid"`
	Version          struct {
		Created string `json:"created"`
	} `json:"version"`
}

type SettingsConfig struct {
	Index IndexConfig `json:"index"`
}

type example struct {
	Settings SettingsConfig `json:"settings"`
}
type Assignment1 struct {
	Comp7S20211122 example `json:"comp-7-s-2021.11.22"`
	Comp7S20211123 example `json:"comp-7-s-2021.11.23"`
}

func ReadJSONFile(filename string) (Assignment1, error) {
	content, err := os.Open(filename)
	if err != nil {
		return Assignment1{}, fmt.Errorf("error opening file: %v", err)
	}
	defer content.Close()

	var data Assignment1
	decoder := json.NewDecoder(content)
	err = decoder.Decode(&data)
	if err != nil {
		return Assignment1{}, fmt.Errorf("error decoding JSON: %v", err)
	}

	return data, nil
}

func main() {

	data, err := ReadJSONFile("input.json")
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("Comp7S20211123 Analyzer Tokenizer:", data.Comp7S20211123.Settings.Index.Analysis.Tokenizer)

	fmt.Println("Comp7S20211122 Analyzer auto complete version filter:", data.Comp7S20211122.Settings.Index.Analysis.Analyzer.Filter)

	fmt.Println("Normalizer Type:", data.Comp7S20211122.Settings.Index.Analysis.Normalizer.CaseInsensitive.Type)
	fmt.Println("Analyzer Tokenizer Type:", data.Comp7S20211122.Settings.Index.Analysis.Analyzer.Tokenizer.Type)

}

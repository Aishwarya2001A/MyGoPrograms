package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type AutocompleteTokenizer struct {
	TokenChars []string `json:"token_chars"`
	MinGram    string   `json:"min_gram"`
	Type       string   `json:"type"`
	MaxGram    string   `json:"max_gram"`
}

type Analyzer struct {
	Autocomplete               Autocomplete `json:"autocomplete"`
	AutocompleteVersionNumbers Autocomplete `json:"autocomplete_version_numbers"`
}

type Autocomplete struct {
	Filter    []string `json:"filter"`
	Tokenizer string   `json:"tokenizer"`
}

type Normalizer struct {
	CaseInsensitive CaseInsensitive `json:"case_insensitive"`
}
type CaseInsensitive struct {
	Filter     []string `json:"filter"`
	Type       string   `json:"type"`
	CharFilter []string `json:"char_filter"`
}

type Analysis struct {
	Normalizer Normalizer `json:"normalizer"`
	Analyzer   Analyzer   `json:"analyzer"`
	Tokenizer  Tokenizer  `json:"tokenizer"`
}
type Tokenizer struct {
	AutocompleteVersionNumberTokenizer AutocompleteTokenizer `json:"autocomplete_version_number_tokenizer"`
	AutocompleteTokenizer              AutocompleteTokenizer `json:"autocomplete_tokenizer"`
}

type Version struct {
	Created string `json:"created"`
}

type Index struct {
	RefreshInterval  string   `json:"refresh_interval"`
	NumberOfShards   string   `json:"number_of_shards"`
	ProvidedName     string   `json:"provided_name"`
	CreationDate     string   `json:"creation_date"`
	Analysis         Analysis `json:"analysis"`
	NumberOfReplicas string   `json:"number_of_replicas"`
	UUID             string   `json:"uuid"`
	Version          Version  `json:"version"`
}

type Settings struct {
	Index Index `json:"index"`
}

type example struct {
	Settings Settings `json:"settings"`
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
	fmt.Println("Comp7S20211123 UUID:", data.Comp7S20211123.Settings.Index.UUID)
	fmt.Println("Comp7S20211122 Analyzer auto complete version filter:", data.Comp7S20211122.Settings.Index.Analysis.Analyzer)
	fmt.Println("Normalizer Type:", data.Comp7S20211122.Settings.Index.Analysis.Normalizer.CaseInsensitive.Type)
	fmt.Println("Analyzer Tokenizer Type:", data.Comp7S20211122.Settings.Index.Analysis.Normalizer.CaseInsensitive.Type)

	fmt.Println("------------------------------------------------------------------------")

	fmt.Println("Comp7S20211122 Refresh Interval:", data.Comp7S20211122.Settings.Index.RefreshInterval)
	fmt.Println("Comp7S20211122 Number of Shards:", data.Comp7S20211122.Settings.Index.NumberOfShards)
	fmt.Println("Comp7S20211122 ProvidedName:", data.Comp7S20211122.Settings.Index.ProvidedName)
	fmt.Println("Comp7S20211122 CreationDate:", data.Comp7S20211122.Settings.Index.CreationDate)

	fmt.Println("------------------------------------------------------------------------")

	fmt.Println("Comp7S20211123 Refresh Interval:", data.Comp7S20211123.Settings.Index.RefreshInterval)
	fmt.Println("Comp7S20211123 Number of Shards:", data.Comp7S20211123.Settings.Index.NumberOfShards)
	fmt.Println("Comp7S20211123 ProvidedName:", data.Comp7S20211123.Settings.Index.ProvidedName)
	fmt.Println("Comp7S20211123 CreationDate:", data.Comp7S20211123.Settings.Index.CreationDate)

}

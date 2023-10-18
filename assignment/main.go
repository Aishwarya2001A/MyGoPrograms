package main

import (
	"encoding/json"
	"fmt"
)

type Assignment struct {
	Comp7S20211122 struct {
		Settings struct {
			Index struct {
				RefreshInterval string `json:"refresh_interval"`
				NumberOfShards  string `json:"number_of_shards"`
				ProvidedName    string `json:"provided_name"`
				CreationDate    string `json:"creation_date"`
				Analysis        struct {
					Normalizer struct {
						CaseInsensitive struct {
							Filter     []string      `json:"filter"`
							Type       string        `json:"type"`
							CharFilter []interface{} `json:"char_filter"`
						} `json:"case_insensitive"`
					} `json:"normalizer"`
					Analyzer struct {
						Autocomplete struct {
							Filter    []string `json:"filter"`
							Tokenizer string   `json:"tokenizer"`
						} `json:"autocomplete"`
						AutocompleteVersionNumbers struct {
							Filter    []string `json:"filter"`
							Tokenizer string   `json:"tokenizer"`
						} `json:"autocomplete_version_numbers"`
					} `json:"analyzer"`
					Tokenizer struct {
						AutocompleteVersionNumberTokenizer struct {
							TokenChars []string `json:"token_chars"`
							MinGram    string   `json:"min_gram"`
							Type       string   `json:"type"`
							MaxGram    string   `json:"max_gram"`
						} `json:"autocomplete_version_number_tokenizer"`
						AutocompleteTokenizer struct {
							TokenChars []string `json:"token_chars"`
							MinGram    string   `json:"min_gram"`
							Type       string   `json:"type"`
							MaxGram    string   `json:"max_gram"`
						} `json:"autocomplete_tokenizer"`
					} `json:"tokenizer"`
				} `json:"analysis"`
				NumberOfReplicas string `json:"number_of_replicas"`
				UUID             string `json:"uuid"`
				Version          struct {
					Created string `json:"created"`
				} `json:"version"`
			} `json:"index"`
		} `json:"settings"`
	} `json:"comp-7-s-2021.11.22"`

	Comp7S20211123 struct {
		Settings struct {
			Index struct {
				RefreshInterval string `json:"refresh_interval"`
				NumberOfShards  string `json:"number_of_shards"`
				ProvidedName    string `json:"provided_name"`
				CreationDate    string `json:"creation_date"`
				Analysis        struct {
					Normalizer struct {
						CaseInsensitive struct {
							Filter     []string      `json:"filter"`
							Type       string        `json:"type"`
							CharFilter []interface{} `json:"char_filter"`
						} `json:"case_insensitive"`
					} `json:"normalizer"`
					Analyzer struct {
						Autocomplete struct {
							Filter    []string `json:"filter"`
							Tokenizer string   `json:"tokenizer"`
						} `json:"autocomplete"`
						AutocompleteVersionNumbers struct {
							Filter    []string `json:"filter"`
							Tokenizer string   `json:"tokenizer"`
						} `json:"autocomplete_version_numbers"`
					} `json:"analyzer"`
					Tokenizer struct {
						AutocompleteVersionNumberTokenizer struct {
							TokenChars []string `json:"token_chars"`
							MinGram    string   `json:"min_gram"`
							Type       string   `json:"type"`
							MaxGram    string   `json:"max_gram"`
						} `json:"autocomplete_version_number_tokenizer"`
						AutocompleteTokenizer struct {
							TokenChars []string `json:"token_chars"`
							MinGram    string   `json:"min_gram"`
							Type       string   `json:"type"`
							MaxGram    string   `json:"max_gram"`
						} `json:"autocomplete_tokenizer"`
					} `json:"tokenizer"`
				} `json:"analysis"`
				NumberOfReplicas string `json:"number_of_replicas"`
				UUID             string `json:"uuid"`
				Version          struct {
					Created string `json:"created"`
				} `json:"version"`
			} `json:"index"`
		} `json:"settings"`
	} `json:"comp-7-s-2021.11.23"`
}

func main() {
	jsonData := `{
		"comp-7-s-2021.11.22": {
			"settings": {
				"index": {
					"refresh_interval": "1s",
					"number_of_shards": "5",
					"provided_name": "comp-7-s-2021.11.22",
					"creation_date": "1637661228822",
					"analysis": {
						"normalizer": {
							"case_insensitive": {
								"filter": ["lowercase", "asciifolding"],
								"type": "custom",
								"char_filter": []
							}
						},
						"analyzer": {
							"autocomplete": {
								"filter": ["lowercase"],
								"tokenizer": "autocomplete_tokenizer"
							},
							"autocomplete_version_numbers": {
								"filter": ["lowercase"],
								"tokenizer": "autocomplete_version_number_tokenizer"
							}
						},
						"tokenizer": {
							"autocomplete_version_number_tokenizer": {
								"token_chars": ["letter", "digit", "punctuation"],
								"min_gram": "2",
								"type": "edge_ngram",
								"max_gram": "20"
							},
							"autocomplete_tokenizer": {
								"token_chars": ["letter", "digit"],
								"min_gram": "2",
								"type": "edge_ngram",
								"max_gram": "20"
							}
						}
					},
					"number_of_replicas": "1",
					"uuid": "w72cfW0GR8uCG8snqUI7Dg",
					"version": {
						"created": "6081499"
					}
				}
			}
		},
		"comp-7-s-2021.11.23": {
			"settings": {
				"index": {
					"refresh_interval": "1s",
					"number_of_shards": "5",
					"provided_name": "comp-7-s-2021.11.23",
					"creation_date": "1637661229195",
					"analysis": {
						"normalizer": {
							"case_insensitive": {
								"filter": ["lowercase", "asciifolding"],
								"type": "custom",
								"char_filter": []
							}
						},
						"analyzer": {
							"autocomplete": {
								"filter": ["lowercase"],
								"tokenizer": "autocomplete_tokenizer"
							},
							"autocomplete_version_numbers": {
								"filter": ["lowercase"],
								"tokenizer": "autocomplete_version_number_tokenizer"
							}
						},
						"tokenizer": {
							"autocomplete_version_number_tokenizer": {
								"token_chars": ["letter", "digit", "punctuation"],
								"min_gram": "2",
								"type": "edge_ngram",
								"max_gram": "20"
							},
							"autocomplete_tokenizer": {
								"token_chars": ["letter", "digit"],
								"min_gram": "2",
								"type": "edge_ngram",
								"max_gram": "20"
							}
						}
					},
					"number_of_replicas": "1",
					"uuid": "XjiLB2xYTwqk7F83FGL_hg",
					"version": {
						"created": "6081499"
					}
				}
			}
		}
	}`

	var data Assignment
	if err := json.Unmarshal([]byte(jsonData), &data); err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Comp7S20211122 Refresh Interval:", data.Comp7S20211122.Settings.Index.RefreshInterval)
	fmt.Println("Comp7S20211122 Number of Shards:", data.Comp7S20211122.Settings.Index.NumberOfShards)
	fmt.Println("Comp7S20211122 ProvidedName:", data.Comp7S20211122.Settings.Index.ProvidedName)
	fmt.Println("Comp7S20211122 CreationDate:", data.Comp7S20211122.Settings.Index.CreationDate)

	fmt.Println("------------------------------------------------------------------")

	fmt.Println("Comp7S20211123 Refresh Interval:", data.Comp7S20211123.Settings.Index.RefreshInterval)
	fmt.Println("Comp7S20211123 Number of Shards:", data.Comp7S20211123.Settings.Index.NumberOfShards)
	fmt.Println("Comp7S20211123 ProvidedName:", data.Comp7S20211123.Settings.Index.ProvidedName)
	fmt.Println("Comp7S20211123 CreationDate:", data.Comp7S20211123.Settings.Index.CreationDate)
}

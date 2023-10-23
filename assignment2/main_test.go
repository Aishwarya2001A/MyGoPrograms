package main

import (
	"encoding/json"
	"os"
	"reflect"
	"testing"
)

func TestReadJSONFile(t *testing.T) {

	testFile := "test.json"

	data, err := ReadJSONFile(testFile)
	if err != nil {
		t.Fatalf("Error reading JSON file: %v", err)
	}

	expectedDataFile := "expected.json"
	expectedContent, err := os.Open(expectedDataFile)
	if err != nil {
		t.Fatalf("Error opening expected data file: %v", err)
	}
	defer expectedContent.Close()

	var expectedData Assignment1
	decoder := json.NewDecoder(expectedContent)
	err = decoder.Decode(&expectedData)
	if err != nil {
		t.Fatalf("Error decoding expected data JSON: %v", err)
	}

	if !reflect.DeepEqual(data, expectedData) {
		t.Errorf("Expected %+v, got %+v", expectedData, data)
	}
}

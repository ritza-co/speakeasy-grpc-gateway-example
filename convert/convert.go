package main

import (
	"encoding/json"
	"os"

	"github.com/getkin/kin-openapi/openapi2conv"
	"github.com/getkin/kin-openapi/openapi2"
)

func main() {
	input, err := os.ReadFile("./openapi/speakeasy/v1/speakeasy.swagger.json")
	if err != nil {
		panic(err)
	}
	
	var doc2 openapi2.T

	if err = json.Unmarshal(input, &doc2); err != nil {
		panic(err)
	}

	doc3, err := openapi2conv.ToV3(&doc2)
	if err != nil {
		panic(err)
	}

	output, err := json.MarshalIndent(doc3, "", "  ")
	if err != nil {
		panic(err)
	}
	if err = os.WriteFile("./openapi/speakeasy/v1/speakeasy.openapi.json", output, 0644); err != nil {
		panic(err)
	}
}

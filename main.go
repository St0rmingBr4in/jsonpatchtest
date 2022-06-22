package main

import (
	"fmt"

	jsonpatch "github.com/evanphx/json-patch"
)

func main() {
  original := []byte(`{"object": {"value1": "value1"}}`)
	patchJSON := []byte(`[
        {"op": "add", "path": "/object", "value": {"value2": "value2"}}
	]`)

	patch, err := jsonpatch.DecodePatch(patchJSON)
	if err != nil {
		panic(err)
	}

	modified, err := patch.Apply(original)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Original document: %s\n", original)
	fmt.Printf("Modified document: %s\n", modified)
}

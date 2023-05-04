package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/nsf/jsondiff"

	"github.com/grokify/aitutor/mrranedeer"
)

func main() {
	prompt := mrranedeer.PromptDefault()

	j, err := json.MarshalIndent(prompt, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(j))

	diff, str, err := mrranedeer.PromptDefaultEqual(mrranedeer.VersionLatest, nil)
	if err != nil {
		log.Fatal(err)
	}
	if diff != jsondiff.FullMatch {
		fmt.Println(str)
	}
	fmt.Printf("Match Result (%s)\n", diff.String())

	fmt.Println("DONE")
}

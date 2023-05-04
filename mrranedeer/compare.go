package mrranedeer

import (
	"encoding/json"

	"github.com/nsf/jsondiff"
)

func PromptDefaultEqual(version string, opts *jsondiff.Options) (jsondiff.Difference, string, error) {
	if opts == nil {
		options := jsondiff.DefaultJSONOptions()
		opts = &options
	}

	refBytes, err := VersionBytes(version)
	if err != nil {
		return 2, "", err
	}

	prompt := PromptDefault()
	genBytes, err := json.Marshal(prompt)
	if err != nil {
		return 2, "", err
	}

	d, s := jsondiff.Compare(refBytes, genBytes, opts)
	return d, s, nil
}

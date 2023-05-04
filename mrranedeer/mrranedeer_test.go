package mrranedeer

import (
	"fmt"
	"testing"

	"github.com/nsf/jsondiff"
)

var versionTests = []struct {
	v    string
	comp []byte
}{
	{Version236, Version236Bytes()},
}

func TestVersions(t *testing.T) {
	for _, tt := range versionTests {
		m, _, err := PromptDefaultEqual(tt.v, nil)
		if err != nil {
			t.Errorf("mrranedeer.PromptDefaultEqual(%s, nil) Error: (%s)", tt.v, err.Error())
		}
		fmt.Printf("Match type (%s)\n", m.String())
		if m != jsondiff.FullMatch {
			t.Errorf("mrranedeer.PromptDefaultEqual(%s, nil) Match type is (%s) not (%s)", tt.v, jsondiff.FullMatch.String(), m.String())
		}
	}
}

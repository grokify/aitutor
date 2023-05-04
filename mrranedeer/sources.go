package mrranedeer

import (
	_ "embed"
	"fmt"
	"io"
	"net/http"
)

//go:embed mrranedeer_v2.3.6.json
var mrranedeer236 []byte

//go:embed mrranedeer_v2.0.0.json
var mrranedeer200 []byte

const (
	VersionLatest = "latest"
	Version236    = "2.3.6"
	Version200    = "2.0.0"
)

func Version236Bytes() []byte {
	return mrranedeer236
}

func Version200Bytes() []byte {
	return mrranedeer200
}

const SrcURL = "https://raw.githubusercontent.com/JushBJJ/Mr.-Ranedeer-AI-Tutor/main/Mr_Ranedeer.json"

func versionLatestBytes() ([]byte, error) {
	resp, err := http.Get(SrcURL)
	if err != nil {
		return []byte{}, err
	}
	if resp.StatusCode >= 300 || resp.StatusCode < 200 {
		return []byte{}, fmt.Errorf("error (%d)", resp.StatusCode)
	}
	return io.ReadAll(resp.Body)
}

func VersionBytes(version string) ([]byte, error) {
	switch version {
	case VersionLatest:
		return versionLatestBytes()
	case "": // return latest local version
		return Version236Bytes(), nil
	case Version236:
		return Version236Bytes(), nil
	case Version200:
		return Version200Bytes(), nil
	default:
		return []byte{}, fmt.Errorf("version not supported (%s)", version)
	}
}

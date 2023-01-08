package aminoacids

import (
	_ "embed"
	"encoding/json"
)

//go:embed aminoacids.json
var AminoAcidsData []byte

var AminoAcids map[string]AminoAcid

// Properties of aminoacids
type AminoAcid struct {
	Image string  `json:"image"`
	Mass  float32 `json:"mass"`
}

func LoadData(data *[]byte) map[string]AminoAcid {
	out := make(map[string]AminoAcid)
	err := json.Unmarshal(*data, &out)
	if err != nil {
		panic(err)
	}
	return out
}

func init() {
	AminoAcids = LoadData(&AminoAcidsData)
}

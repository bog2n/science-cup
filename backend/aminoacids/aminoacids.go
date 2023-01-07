package aminoacids

import (
	"encoding/json"
)

// Properties of aminoacids
type AminoAcid struct {
	Image string  `json:"image"`
	Mass  float32 `json:"mass"`
}

var AminoAcids map[string]AminoAcid

func LoadData(data *[]byte) map[string]AminoAcid {
	out := make(map[string]AminoAcid)
	err := json.Unmarshal(*data, &out)
	if err != nil {
		panic(err)
	}
	return out
}

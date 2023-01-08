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
	Mass  float64 `json:"mass"`
}

func LoadData(data *[]byte) map[string]AminoAcid {
	out := make(map[string]AminoAcid)
	err := json.Unmarshal(*data, &out)
	if err != nil {
		panic(err)
	}
	return out
}

func CalculateMass(protein string) float64 {
	var mass float64
	for _, v := range protein {
		mass += AminoAcids[string(v)].Mass
	}
	mass += 18.010564686 // h2o molecule mass
	return mass
}

func init() {
	AminoAcids = LoadData(&AminoAcidsData)
}

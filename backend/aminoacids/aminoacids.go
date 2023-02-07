package aminoacids

import (
	_ "embed"
	"encoding/json"
	"log"
	"math"
)

const Debug = true

func debug(v ...any) {
	if Debug {
		log.Print(v...)
	}
}

//go:embed aminoacids.json
var AminoAcidsData []byte

var AminoAcids map[string]AminoAcid

// Properties of aminoacids
type AminoAcid struct {
	Image     string  `json:"image"`
	Mass      float64 `json:"mass"`
	HydroPhob float64 `json:"hydrophob"`
	Pi        float64 `json:"PI"`
}

// Populates aminoacids map with json data from
func LoadData(data *[]byte) map[string]AminoAcid {
	out := make(map[string]AminoAcid)
	err := json.Unmarshal(*data, &out)
	if err != nil {
		panic(err)
	}
	return out
}

// Calculates mass of protein defined in string
func CalculateMass(protein string) float64 {
	var mass float64
	for _, v := range protein {
		mass += AminoAcids[string(v)].Mass
	}
	mass += 18.010564686 // h2o molecule mass
	return mass
}

// Calculates hydrophobic index of protein defined in string
func CalculateHydroIndex(protein string) float64 {
	var out float64 = 7.9
	for _, v := range protein {
		out += AminoAcids[string(v)].HydroPhob
	}
	return out
}

// Calculates Isoelectric Point for given protein
func CalculatePI(protein string) float64 {
	count := make(map[string]float64)
	// Update counts of proteins
	for _, v := range protein {
		count[string(v)]++
	}
	var ph float64 = 0
	for ph = 0; ph <= 14; ph += 0.05 {
		var q = 0.0
		// Acidic
		q -= 1.0 / (1.0 + math.Pow(10.0, (3.65-ph)))        // COOH
		q -= count["C"] / (1.0 + math.Pow(10.0, (8.18-ph))) // C
		q -= count["D"] / (1.0 + math.Pow(10.0, (3.9-ph)))  // D
		q -= count["E"] / (1.0 + math.Pow(10.0, (4.07-ph))) // E
		q -= count["H"] / (1.0 + math.Pow(10.0, (6.04-ph))) // H
		// Based
		q += 1.0 / (1.0 + math.Pow(10.0, (ph-8.2)))          // NH2
		q += count["K"] / (1.0 + math.Pow(10.0, (ph-10.54))) // K
		q += count["R"] / (1.0 + math.Pow(10.0, (ph-12.48))) // R
		q += count["Y"] / (1.0 + math.Pow(10.0, (ph-10.46))) // Y
		if q <= 0 {
			break
		}
	}
	return ph
}

// Calculates PH Index for given protein
func CalculatePH(protein string) float64 {
	// Isoelectric point
	pi := CalculatePI(protein)
	// Mass
	m := CalculateMass(protein)

	return -math.Log10(pi / m)
}

func init() {
	AminoAcids = LoadData(&AminoAcidsData)
}

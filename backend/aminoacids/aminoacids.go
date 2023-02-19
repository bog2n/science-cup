package aminoacids

import (
	_ "embed"
	"encoding/json"
	"log"
	"math"
)

const Debug = false

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
	Mass      float64 `json:"mass"`
	HydroPhob float64 `json:"hydrophob"`
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
	for s := 1.0; s <= 8; s++ {
		q := calculateCharge(count, ph)
		if q > 0 {
			ph += 14 / math.Pow(2, s)
		} else {
			ph -= 14 / math.Pow(2, s)
		}
		debug("ph ", ph, " step ", s)
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

// Calculates polarity of protein
func CalculatePolarity(protein string) float64 {
	count := make(map[string]float64)
	// Update counts of proteins
	for _, v := range protein {
		count[string(v)]++
	}
	var ph float64 = 7.4 // Neutral pH

	return math.Round(calculateCharge(count, ph))
}

func calculateCharge(counts map[string]float64, ph float64) float64 {
	var q float64 = 0

	// Acidic
	q -= 1 / (1 + math.Pow(10.0, (3.65-ph)))         // COOH
	q -= counts["C"] / (1 + math.Pow(10, (8.18-ph))) // C
	q -= counts["D"] / (1 + math.Pow(10, (3.9-ph)))  // D
	q -= counts["E"] / (1 + math.Pow(10, (4.07-ph))) // E
	q -= counts["H"] / (1 + math.Pow(10, (6.04-ph))) // H
	// Based
	q += 1 / (1 + math.Pow(10, (ph-8.2)))             // NH2
	q += counts["K"] / (1 + math.Pow(10, (ph-10.54))) // K
	q += counts["R"] / (1 + math.Pow(10, (ph-12.48))) // R
	q += counts["Y"] / (1 + math.Pow(10, (ph-10.46))) // Y

	return q
}

func init() {
	AminoAcids = LoadData(&AminoAcidsData)
}

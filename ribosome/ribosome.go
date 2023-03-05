// Ribosome package - implements (nearly) all ribosome functions ;)
package ribosome

import (
	"errors"
	"log"
	"strings"
)

const Debug = false

func debug(v ...any) {
	if Debug {
		log.Print(v...)
	}
}

// Error when unknown nucleotide is found in the sequence
var UnknownCodonError error = errors.New("Unknown codon")

// Error when no proteins where found in the sequence
var NoProteinError error = errors.New("No protein found")

// Returns all amino acid strings in given RNA/DNA sequence
func GetAminoAcids(rna string) ([]string, error) {
	var out []string
	var protein string
	var start bool

	rna = DNA2RNA(rna)
	debug("Starting decoding")
	for i := 0; i < len(rna)-len(rna)%3; i += 3 {
		a, ok := AminoAcids[rna[i:i+3]]
		if !ok {
			debug("Unknown codon sequence")
			return []string{""}, UnknownCodonError
		}
		// Look for start codon
		if a == "M" {
			debug("Got START codon")
			start = true
		}
		// Amino sequence started, add acid
		if start {
			if a == "0" {
				debug("Got STOP codon returning: ", out)
				out = append(out, protein)
				// Reset
				protein = ""
				start = false
				continue
			}
			debug("Adding ", a, " to sequence")
			protein += a
		}
	}
	// No amino acids found
	if len(out) == 0 {
		debug("Start and stop codon not found")
		return []string{""}, NoProteinError
	}
	// Return what we got
	return out, nil
}

// Converts DNA string to RNA string and cleans up the string
func DNA2RNA(dna string) string {
	dna = strings.Replace(dna, " ", "", -1)
	dna = strings.Replace(dna, "\n", "", -1)
	dna = strings.Replace(dna, "\t", "", -1)
	return strings.Replace(strings.ToUpper(dna), "T", "U", -1)
}

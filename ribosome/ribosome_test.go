package ribosome

import (
	"reflect"
	"testing"
)

func TestGetAminoAcids(t *testing.T) {
	tests := map[string][]string{
		"AAAAAUAUGAACGAAAAUCUGUucgcuucauucauugcccccacaauccuaggccuacccUGA":    {"MNENLFASFIAPTILGLP"},
		"AAAAAUaugaAC   GAAAAucugUUCGCuuCAUucauugcCCCCACAAUccUAGGCCUaCCCUGA": {"MNENLFASFIAPTILGLP"},
		"AUGAAAUAACCCCCCAUggGGUAG": {"MK", "MG"},
	}

	// Loop over test cases
	for val, want := range tests {
		got, _ := GetAminoAcids(val)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Got %q want %q", got, want)
		}
	}

	// Test errors
	errortests := map[string]error{
		"AAAAAUAUGAACGAAAAUCUGUUCGCUUCAUUCAUUGCCCCCACAAUCCUAGGCCUACCCUGA": nil,
		"BBB": UnknownCodonError,
		"AAAAAUAUGAACGAAAAUCUGUUCGCUUCAUUCAUUGCCCCCACAAUCCUAGGCCUACCC": NoProteinError,
		"AAAACGAAAAUCUGUUCGCUUCAUUCAUUGCCCCCACAAUCCUAGGCCUACCC":        NoProteinError,
	}
	// Loop over test cases
	for val, want := range errortests {
		_, got := GetAminoAcids(val)
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}
}

func TestDNA2RNA(t *testing.T) {
	tests := map[string]string{
		"ACTG": "ACUG",
		"TTTG": "UUUG",
		"TttG": "UUUG",
		"AAaG": "AAAG",
	}
	for val, want := range tests {
		got := DNA2RNA(val)
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}
}

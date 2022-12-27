package ribosome

import "testing"

func TestGetAminoAcids(t *testing.T) {
	tests := map[string][]string{
		"AAAAAUAUGAACGAAAAUCUGUucgcuucauucauugcccccacaauccuaggccuacccUGA":    {"MNENLFASFIAPTILGLP"},
		"AAAAAUaugaAC   GAAAAucugUUCGCuuCAUucauugcCCCCACAAUccUAGGCCUaCCCUGA": {"MNENLFASFIAPTILGLP"},
		"AUGAAAUAACCCCCCAUggGGUAG": {"MK", "MG"},
	}

	for val, wwant := range tests {
		ggot, _ := GetAminoAcids(val)

		// Loop over want items
		for want := range wwant {
			eq := false
			// Loop over output items
			for got := range ggot {
				// If any value is equal then we got a match
				if want == got {
					eq = true
				}
			}
			if eq == false {
				t.Errorf("Got %q want %q", ggot, wwant)
			}
		}
	}
	// Test errors
	errortests := map[string]error{
		"AAAAAUAUGAACGAAAAUCUGUUCGCUUCAUUCAUUGCCCCCACAAUCCUAGGCCUACCCUGA": nil,
		"BBB": unknownCodonError,
		"AAAAAUAUGAACGAAAAUCUGUUCGCUUCAUUCAUUGCCCCCACAAUCCUAGGCCUACCC": noProteinError,
		"AAAACGAAAAUCUGUUCGCUUCAUUCAUUGCCCCCACAAUCCUAGGCCUACCC":        noProteinError,
	}
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

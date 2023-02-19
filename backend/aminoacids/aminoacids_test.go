package aminoacids

import (
	"math"
	"os"
	"reflect"
	"testing"
)

func TestLoadData(t *testing.T) {
	// Read test JSON file
	data, err := os.ReadFile("test.json")
	if err != nil {
		panic(err)
	}
	// Define wanted data
	want := make(map[string]AminoAcid)
	want["U"] = AminoAcid{
		Mass:      22.0,
		HydroPhob: 2.22,
	}
	want["A"] = AminoAcid{
		Mass:      21.3,
		HydroPhob: 1.11,
	}

	got := LoadData(&data)
	if !reflect.DeepEqual(want, got) {
		t.Error("Want: ", want, "\nGot: ", got)
	}
}

func TestCalculateMass(t *testing.T) {
	tests := map[string]float64{
		"GFPCM":                553.2022,
		"EHDGYIFVS":            1065.4752,
		"HFYNRQEKTFH":          1505.7144,
		"AQLSTKERNGMWYFHDCIPV": 2394.1214,
	}
	for val, want := range tests {
		got := CalculateMass(val)
		diff := math.Abs(got - want)
		if diff > 0.1 {
			t.Errorf("For data: %q, got %f want %f | diff = %f", val, got, want, diff)
		}
	}
}

func TestCalculateHydroIndex(t *testing.T) {
	tests := map[string]float64{
		"EGNDH":                19.50,
		"DHFYSA":               12.41,
		"CIPVYWTSLQARNDHFMKEG": 18.20,
	}
	for val, want := range tests {
		got := CalculateHydroIndex(val)
		diff := math.Abs(got - want)
		if diff > 0.1 {
			t.Errorf("For data: %q, got %f want %f | diff = %f", val, got, want, diff)
		}
	}
}

func TestCalculatePI(t *testing.T) {
	tests := map[string]float64{
		"ASDF":                 3.77,
		"CIPVYWTSLQARNDHFMKEG": 6.96,
	}
	for val, want := range tests {
		got := CalculatePI(val)
		diff := math.Abs(got - want)
		if diff > 0.1 {
			t.Errorf("For data: %q, got %f want %f | diff = %f", val, got, want, diff)
		}
	}
}

func TestCalculatePH(t *testing.T) {
	tests := map[string]float64{
		"ASDF":                 2.06,
		"CIPVYWTSLQARNDHFMKEG": 2.55,
		"AQLSTKERNGMWYFHDCIPV": 2.54,
	}
	for val, want := range tests {
		got := CalculatePH(val)
		diff := math.Abs(got - want)
		if diff > 0.1 {
			t.Errorf("For data: %q, got %f want %f | diff = %f", val, got, want, diff)
		}
	}
}

func TestCalculatePolarity(t *testing.T) {
	tests := map[string]float64{
		"ASDF":                 -1,
		"CIPVYWTSLQARNDHFMKEG": 0,
		"RKA":                  2,
	}
	for val, want := range tests {
		got := CalculatePolarity(val)
		if want != got {
			t.Errorf("Got %f want %f", got, want)
		}
	}
}

func BenchmarkCalculatePI(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculatePI("AQLSTKERNGMWYFHDCIPV")
	}
}

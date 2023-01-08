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
		Image: `<path d="m 100 100 v 500 z">`,
		Mass:  22.0,
	}
	want["A"] = AminoAcid{
		Image: `aaaa`,
		Mass:  21.3,
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
		"AQLSTKERNGMWYFHDCIPV": 2394.1214, // all of them
	}
	for val, want := range tests {
		got := CalculateMass(val)
		diff := math.Abs(got - want)
		if diff > 0.1 {
			t.Errorf("For data: %q, got %f want %f | diff = %f", val, got, want, diff)
		}
	}
}

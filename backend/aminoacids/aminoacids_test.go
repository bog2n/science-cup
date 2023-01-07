package aminoacids

import (
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

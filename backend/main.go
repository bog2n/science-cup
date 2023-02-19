package main

import (
	"encoding/json"
	"log"
	"motorola/aminoacids"
	"motorola/image"
	"motorola/ribosome"
	"net/http"
)

type Response struct {
	Ok       bool   `json:"ok"`
	Proteins []Data `json:"proteins"`
}

type Data struct {
	Protein     string  `json:"protein"`
	Mass        float64 `json:"mass"`
	HydroIndex  float64 `json:"hindex"`
	Isoelectric float64 `json:"isopoint"`
	PH          float64 `json:"ph"`
	Polarity    float64 `json:"polarity"`
}

func handleData(w http.ResponseWriter, r *http.Request) {
	genome := r.FormValue("genome")
	var out Response
	out.Ok = true
	for i := 0; i < 3; i++ {
		prot, err := ribosome.GetAminoAcids(genome[i:])
		if err != nil {
			out.Ok = false
			log.Print(err)
			break
		}
		for i := range prot {
			var d Data
			d.Protein = prot[i]
			d.Mass = aminoacids.CalculateMass(prot[i])
			d.HydroIndex = aminoacids.CalculateHydroIndex(prot[i])
			d.Isoelectric = aminoacids.CalculatePI(prot[i])
			d.PH = aminoacids.CalculatePH(prot[i])
			d.Polarity = aminoacids.CalculatePolarity(prot[i])
			out.Proteins = append(out.Proteins, d)
		}
	}
	b, _ := json.Marshal(out)
	w.Header().Set("Content-type", "application/json")
	w.Write(b)
}

func handleImage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "image/svg+xml")
	image.DrawProtein(r.FormValue("protein"), w)
}

func handleMain(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`
<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<title></title>
</head>
<body>
	<form action="/data" method=post>
		<input type=text name=genome>
		<input type=submit value=go>
	</form>
</body>
</html>
`))
}

func main() {
	http.HandleFunc("/", handleMain)
	http.HandleFunc("/data", handleData)
	http.HandleFunc("/image", handleImage)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

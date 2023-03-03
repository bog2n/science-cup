package main

import (
	"compress/gzip"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"mime"
	"motorola/aminoacids"
	"motorola/frontend"
	"motorola/image"
	"motorola/ribosome"
	"net/http"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
)

var (
	bindAddress string
	port        uint
	noBrowser   bool
	bindString  string
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

func openBrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}
}

// Handles finding all proteins in RNA/DNA sequence
func handleData(w http.ResponseWriter, r *http.Request) {
	// TODO file uploading handler
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
			out.Proteins = append(out.Proteins, Data{
				Protein:     prot[i],
				Mass:        aminoacids.CalculateMass(prot[i]),
				HydroIndex:  aminoacids.CalculateHydroIndex(prot[i]),
				Isoelectric: aminoacids.CalculatePI(prot[i]),
				PH:          aminoacids.CalculatePH(prot[i]),
				Polarity:    aminoacids.CalculatePolarity(prot[i]),
			})
		}
	}
	b, _ := json.Marshal(out)
	w.Header().Set("Content-type", "application/json")
	w.Write(b)
}

// Handles /image endpoint for generating images
func handleImage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "image/svg+xml")
	if strings.Contains(r.Header.Get("Accept-encoding"), "gzip") {
		g := gzip.NewWriter(w)
		defer g.Close()
		w.Header().Set("Content-encoding", "gzip")
		image.DrawProtein(r.URL.Path[len("/image/"):], g)
	} else {
		image.DrawProtein(r.URL.Path[len("/image/"):], w)
	}
}

// TODO replace with actual frontend
func handleMain(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	if path == "/" {
		path = "/index.html"
	}
	path = "out" + path
	b, err := frontend.Content.ReadFile(path)
	if err != nil {
		w.WriteHeader(404)
		w.Write(frontend.NotFound)
		return
	}
	ext := filepath.Ext(path)
	m := mime.TypeByExtension(ext)
	w.Header().Set("Content-type", m)
	w.Write(b)
}

func init() {
	// Parse cli flags
	flag.StringVar(&bindAddress, "b", "127.0.0.1", "Address to listen on")
	flag.UintVar(&port, "p", 8080, "Port to listen on")
	flag.BoolVar(&noBrowser, "nobrowser", false, "Do not open browser on startup")
	flag.Parse()
	bindString = bindAddress + ":" + strconv.Itoa(int(port))
}

func main() {
	if !noBrowser {
		openBrowser("http://" + bindString)
	}
	http.HandleFunc("/", handleMain)
	http.HandleFunc("/api/data", handleData)
	http.HandleFunc("/api/image/", handleImage)
	log.Fatal(http.ListenAndServe(bindString, nil))
}

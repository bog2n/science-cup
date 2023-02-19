# Notes

## Roadmap

- ~~add embedded json file with all aminoacids~~
- ~~create functions for calculating properties of proteins~~
	- ~~calculate mass of protein~~
	- ~~calculate hydrophobic index~~
	- ~~calculate pH index~~
	- ~~calculate polarity~~
- ~~add images in svg format~~
- ~~handle image translating and merging~~
- write http server api documentation
- write endpoint handlers

## HTTP API

### POST `/data` - `genome=<RNA/DNA sequence>`

Sample data:
```json
{
  "ok": true,
  "proteins": [
    {
      "protein": "MHFVRTTGLY",
      "mass": 1223.6121131859998,
      "hindex": 8.89,
      "isopoint": 9.3515625,
      "ph": 2.1167595868870497,
      "polarity": 1
    },
    {
      "protein": "MSKFCISLKFHNQDYSTKG",
      "mass": 2233.065994286,
      "hindex": 19.480000000000004,
      "isopoint": 9.2421875,
      "ph": 2.3831267830805074,
      "polarity": 2
    },
    {
      "protein": "MQPNVPFNSHEV",
      "mass": 1397.639784486,
      "hindex": 13.77,
      "isopoint": 3.8828125,
      "ph": 2.5562488354658397,
      "polarity": -2
    }
  ]
}
```

### GET `/image?protein=<protein>`

Returns svg image for given protein

## [PL] Obliczanie masy białka

Testując funkcję obliczającą masę białka otrzymywałem tą samą różnicę pomiędzy obliczoną masą a tą która wychodziła z kalkulatorów:
```
--- FAIL: TestCalculateMass (0.00s)
    aminoacids_test.go:44: For data: "GFPCM", got 535.192311 want 553.202200 | diff = 18.009889
    aminoacids_test.go:44: For data: "EHDGYIFVS", got 1047.466160 want 1065.475200 | diff = 18.009040
    aminoacids_test.go:44: For data: "HFYNRQEKTFH", got 1487.705831 want 1505.714400 | diff = 18.008569
    aminoacids_test.go:44: For data: "AQLSTKERNGMWYFHDCIPV", got 2376.114342 want 2394.121400 | diff = 18.007058
```
dziwne...

Zagłębiając się w kod źródłowy z [jednego](https://www.rapidnovor.com/mass-calculator/) z kalkulatorów okazało się że jest to masa cząsteczki wody którą dobrze widać na początku i na końcu jakiegokolwiek z białka które narysujemy w <https://pepdraw.io>

## [PL] Optymalizowanie algorytmu liczącego punkt izoelektryczny

Obliczanie punktu izoelektrycznego jest prostą operacją, dla pH od 0 do 14 liczymy ładunek elektryczny białka i jeżeli jest on równy 0 to to pH jest punktem izoelektrycznym tego białka, oczywiście wszystko jest tylko przybliżeniem i dla dokładnych wyników należy zbadać takie białko w rzeczywistości.

Najprościej jest to zaimplementować robiąc pętle od pH = 0 do ph = 14 obliczamy np. co 0.01 pH ładunek elektryczny białka, i jeżeli jest on równy 0 to zwracamy tą wartość.

Metoda jest prosta, lecz można to zrobić szybciej stosując [bisekcję](https://pl.wikipedia.org/wiki/Metoda_r%C3%B3wnego_podzia%C5%82u), przy zastosowaniu tej metody w naszej funkcji wydajność wzrosła 12 krotnie!

### Pętla
```
(924ee2f...) go test -bench="."
goos: linux
goarch: amd64
pkg: motorola/aminoacids
cpu: Intel(R) Core(TM) i5-3320M CPU @ 2.60GHz
BenchmarkCalculatePI-4              7183            168556 ns/op
PASS
ok      motorola/aminoacids     2.225s
```

### Bisekcja
```
(2077835...) go test -bench="."
goos: linux
goarch: amd64
pkg: motorola/aminoacids
cpu: Intel(R) Core(TM) i5-3320M CPU @ 2.60GHz
BenchmarkCalculatePI-4             94546             13154 ns/op
PASS
ok      motorola/aminoacids     1.381s
```

## Links and sources

- <https://pepdraw.io>
- <https://en.wikipedia.org/wiki/Amino_acid#/media/File:ProteinogenicAminoAcids.svg>
- <https://www.rapidnovor.com/mass-calculator/>
- <http://izoelektryczny.netmark.pl/files/teoria.html>

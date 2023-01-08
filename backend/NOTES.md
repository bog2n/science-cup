# Notes

## Roadmap

- ~~add embedded json file with all aminoacids~~
- create functions for calculating properties of proteins
	- ~~calculate mass of protein~~
	- calculate hydrophobic index
	- calculate pH index
	- calculate polarity
- add images in svg format
- handle image translating and merging
- write http server api documentation
- write endpoint handlers

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

## Links and sources

- <https://pepdraw.io>
- <https://en.wikipedia.org/wiki/Amino_acid#/media/File:ProteinogenicAminoAcids.svg>
- <https://www.rapidnovor.com/mass-calculator/>

package image

import (
	"io"

	svg "github.com/ajstarks/svgo"
)

func DrawProtein(protein string, writer io.Writer) {
	var pos []int
	image := svg.New(writer)
	image.Start(len(protein)*90+120, 340)
	pos = []int{60, 170}
	image.Text(35, 167, "H", "stroke:black")
	image.Text(43, 155, "+", "stroke:black")
	image.Text(49, 167, "2", "stroke:black;font-size:8px")
	for index, value := range protein {
		switch string(value) {
		case "A":
			pos = a(pos, image, index, (index == 0 || index == len(protein)-1))
		case "R":
			pos = r(pos, image, index)
		case "N":
			pos = n(pos, image, index)
		case "D":
			pos = d(pos, image, index)
		case "C":
			pos = c(pos, image, index)
		case "Q":
			pos = q(pos, image, index)
		case "E":
			pos = e(pos, image, index)
		case "G":
			pos = g(pos, image, index)
		case "H":
			pos = h(pos, image, index)
		case "I":
			pos = i(pos, image, index)
		case "L":
			pos = l(pos, image, index)
		case "K":
			pos = k(pos, image, index)
		case "M":
			pos = m(pos, image, index)
		case "F":
			pos = f(pos, image, index)
		case "P":
			pos = p(pos, image, index)
		case "S":
			pos = s(pos, image, index)
		case "T":
			pos = t(pos, image, index)
		case "W":
			pos = w(pos, image, index)
		case "Y":
			pos = y(pos, image, index)
		case "V":
			pos = v(pos, image, index)
		}
	}
	DrawText(pos, "O-", image, R)
	image.End()
}

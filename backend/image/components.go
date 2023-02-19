// Used to produce images of chemical bonds and aminoacids
package image

import (
	"fmt"
	"math/rand"

	svg "github.com/ajstarks/svgo"
)

// Directions
// L - Left
// R - Right
// U - Up
// D - Down
const (
	L = 1 << iota
	R
	U
	D
)

// Directions
// UR - Up right
// UL - Up left
// DR - Down right
// DL - Down left
const (
	UR = U | R
	UL = U | L
	DR = D | R
	DL = D | L
)

// Draw text according to direction of previous point,
// and reverses it if needed
func DrawText(xy []int, text string, canvas *svg.SVG, direction int) {
	var x, y = xy[0], xy[1]
	var xmul, ymul, yoffset, xoffset int
	if direction&L != 0 {
		xmul = -10
		yoffset = 7
		xoffset = -4
	} else if direction&R != 0 {
		xmul = 10
		xoffset = -7
		yoffset = 7
	} else if direction == D {
		ymul = 15
		xmul = 0
		xoffset = -6
	} else if direction == U {
		ymul = -15
		yoffset = 12
		xoffset = -6
	}
	for index, v := range text {
		i := index + 1
		canvas.Text(x+xmul*i+xoffset, y+ymul*i+yoffset, string(v), "stroke:black;text-anchor:center")
	}
}

// Draws single line in specified direction
func DrawLine(xy []int, canvas *svg.SVG, direction int) []int {
	var x, y = xy[0], xy[1]
	var x2, y2 int
	switch direction {
	case U:
		y2 = -30
	case D:
		y2 = 30
	case L:
		x2 = -30
	case R:
		x2 = 30
	case UL:
		y2 = -15
		x2 = -30
	case UR:
		y2 = -15
		x2 = 30
	case DL:
		y2 = 15
		x2 = -30
	case DR:
		y2 = 15
		x2 = 30
	}
	canvas.Line(x, y, x+x2, y+y2, "stroke:black;stroke-width:2px")
	return []int{x + x2, y + y2}
}

// Draws double line in specified direction
func DrawDLine(xy []int, canvas *svg.SVG, direction int) []int {
	var x, y = xy[0], xy[1]
	if direction^U == 0 || direction^D == 0 {
		drawLine([]int{x - 2, y}, canvas, direction)
		a := drawLine([]int{x + 2, y}, canvas, direction)
		return []int{a[0] - 2, a[1]}
	} else {
		drawLine([]int{x, y - 2}, canvas, direction)
		a := drawLine([]int{x, y + 2}, canvas, direction)
		return []int{a[0], a[1] - 2}
	}
}

// Draws triangle in specified direction and reverses it if needed
func DrawTriangle(xy []int, canvas *svg.SVG, direction int, reverse bool) []int {
	var x1, y1, x2, y2, x, y int = 0, 0, 0, 0, xy[0], xy[1]
	switch direction {
	case DR, UL:
		x1 = 2
		y1 = -3
		x2 = (direction&L*2 - 1) * -30
		y2 = (direction&U/2 - 1) * -15
	case UR, DL:
		x1 = 2
		y1 = 3
		x2 = (direction&L*2 - 1) * -30
		y2 = (direction&U/2 - 1) * -15
	case R, L:
		y1 = 3
		x2 = (direction&L*2 - 1) * -30
	case U, D:
		x1 = 3
		y2 = (direction&U/2 - 1) * -30
	}
	var cords [][]int
	if reverse {
		cords = [][]int{
			{x + x1 + x2, x, x - x1 + x2},
			{y + y1 + y2, y, y - y1 + y2},
		}
	} else {
		cords = [][]int{
			{x + x1, x + x2, x - x1},
			{y + y1, y + y2, y - y1},
		}
	}
	canvas.Polyline(cords[0], cords[1], "stroke:black;stroke-width:0;fill:black")

	return []int{x + x2, y + y2}
}

// Draws striped triangle in specified direction and reverses it if needed
func DrawDTriangle(xy []int, canvas *svg.SVG, direction int, reverse bool) []int {
	cID := newString()
	var x, y int = xy[0], xy[1]
	fmt.Fprintf(canvas.Writer, `<clipPath id="%s">`, cID)
	cords := drawTriangle(xy, canvas, direction, reverse)
	canvas.ClipEnd()
	switch direction {
	case U, D:
		for i := 0; i < 5; i++ {
			canvas.Line(x+3,
				y+i*-6*(direction&U/2-1),
				x-3,
				y+i*-6*(direction&U/2-1),
				"stroke:black;stroke-width:3px;clip-path:url(#"+cID+")")
		}
	case L, R:
		for i := 0; i < 5; i++ {
			canvas.Line(x+i*-6*(direction&L*2-1),
				y-3,
				x+i*-6*(direction&L*2-1),
				y+3,
				"stroke:black;stroke-width:3px;clip-path:url(#"+cID+")")
		}
	case DL, UR:
		for i := 0; i < 5; i++ {
			canvas.Line(x+i*-6*(direction&L*2-1)-2,
				y+i*-3*(direction&U/2-1)-3,
				x+i*-6*(direction&L*2-1)+2,
				y+i*-3*(direction&U/2-1)+3,
				"stroke: black;stroke-width:3px;clip-path:url(#"+cID+")")
		}
	case UL, DR:
		for i := 0; i < 5; i++ {
			canvas.Line(x+i*-6*(direction&L*2-1)-2,
				y+i*-3*(direction&U/2-1)+3,
				x+i*-6*(direction&L*2-1)+2,
				y+i*-3*(direction&U/2-1)-3,
				"stroke: black;stroke-width:3px;clip-path:url(#"+cID+")")
		}
	}
	return cords
}

func newString() string {
	letters := "abcdefghijklmnoprstuvwxyz"
	var code string
	for i := 0; i < 32; i++ {
		code += string(letters[rand.Intn(len(letters))])
	}
	return code
}

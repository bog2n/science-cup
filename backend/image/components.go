// Used to produce images of chemical bonds and aminoacids
package image

import (
	"fmt"
	"math/rand"
	"strings"

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

const style = "stroke:black;stroke-width:2px"

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
	if text == "NH2" && xmul < 0 {
		text = "N2H"
	}
	for index, v := range text {
		i := index + 1
		if strings.Contains("1234567890", string(v)) {
			canvas.Text(x+xmul*i+xoffset+4, y+ymul*i+yoffset,
				string(v), "font-size:8px;stroke:black")
		} else if string(v) == "-" {
			canvas.Text(x+xmul*i+xoffset+3, y+ymul*i+yoffset-6,
				string(v), "stroke:black")
		} else {
			canvas.Text(x+xmul*i+xoffset, y+ymul*i+yoffset,
				string(v), "stroke:black")
		}
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
	canvas.Line(x, y, x+x2, y+y2, style)
	return []int{x + x2, y + y2}
}

// Draws double line in specified direction
func DrawDLine(xy []int, canvas *svg.SVG, direction int) []int {
	var x, y = xy[0], xy[1]
	if direction^U == 0 || direction^D == 0 {
		DrawLine([]int{x - 2, y}, canvas, direction)
		a := DrawLine([]int{x + 2, y}, canvas, direction)
		return []int{a[0] - 2, a[1]}
	} else {
		DrawLine([]int{x, y - 2}, canvas, direction)
		a := DrawLine([]int{x, y + 2}, canvas, direction)
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
	cords := DrawTriangle(xy, canvas, direction, reverse)
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

// Draws that strange circle with nitrogen atoms found in Arginine
func drawNitricCircle(xy []int, canvas *svg.SVG, direction int) {
	var x, y = xy[0], xy[1]

	x1 := (direction&U/2 - 1)
	y1 := (direction&L*2 - 1)
	coords := DrawLine([]int{x + x1*2, y},
		canvas, direction)
	for i := 0; i < 3; i++ {
		canvas.Line(x-x1*2, y+i*-12*x1,
			x-x1*2, y+i*-12*x1+(direction&U/2-1)*-6, style)
	}
	x, y = coords[0], coords[1]

	var d int
	if direction == U {
		d = UR
	} else {
		d = DL
	}
	p := DrawLine([]int{x, y - x1*2}, canvas, d)
	DrawText([]int{p[0], p[1] + x1*2}, "NH2", canvas, d)
	for i := 0; i < 3; i++ {
		canvas.Line(x+i*12*x1, y+x1*(i*6*y1+2),
			x+((i*12)+6)*x1, y+x1*((i*6+3)*y1+2), style)
	}

	if direction == U {
		d = UL
	} else {
		d = DR
	}
	p = DrawLine([]int{coords[0] - x1*2, coords[1] + x1*2}, canvas, d)
	DrawText([]int{p[0] + x1*2, p[1] - x1*2}, "NH2", canvas, d)
	x1 = -x1
	y1 = -y1
	for i := 0; i < 3; i++ {
		canvas.Line(x+i*12*x1, y+x1*(i*6*y1+2),
			x+((i*12)+6)*x1, y+x1*((i*6+3)*y1+2), style)
	}

	canvas.Circle(x, y, 8, "fill:white;stroke:black;stroke-width:2px")
	canvas.Text(x, y+6, "+", "stroke:black;stroke-width:2px;text-anchor:middle")
}

// Draws this pentagon thing in Histidine
func drawHPentagon(xy []int, canvas *svg.SVG, direction int) {
	y1 := direction&U/2 - 1
	p := DrawLine(xy, canvas, R)
	DrawText(p, "N", canvas, R)

	canvas.Line(p[0], p[1], p[0]+15, p[1]-30*y1, style)
	canvas.Line(p[0]-4, p[1]-4*y1, p[0]+7, p[1]-26*y1, style)
	DrawLine([]int{p[0] + 15, p[1] - 30*y1}, canvas, L|direction)

	canvas.Line(xy[0], xy[1], xy[0]-15, xy[1]-30*y1, style)
	canvas.Line(xy[0]+4, xy[1]-4*y1, xy[0]-7, xy[1]-26*y1, style)
	p = DrawLine([]int{xy[0] - 15, xy[1] - 30*y1}, canvas, R|direction)

	DrawText(p, "NH", canvas, direction)
}

// Strange combination of pentagon and benzen (tryptophan)
func drawWThingy(xy []int, canvas *svg.SVG, direction int) {
	y1 := direction&U/2 - 1
	p := DrawLine(xy, canvas, R)
	canvas.Line(xy[0]+2, xy[1]-y1*4, xy[0]+28, xy[1]-y1*4, style)

	DrawBensen([]int{xy[0] - 15, xy[1] - 30*y1}, canvas, direction)
	canvas.Line(xy[0], xy[1], xy[0]-15, xy[1]-30*y1, style)
	DrawLine([]int{p[0] + 15, p[1] - 30*y1}, canvas, L|direction)

	canvas.Line(p[0], p[1], p[0]+15, p[1]-30*y1, style)
	DrawText([]int{xy[0] + 45, xy[1] - 30*y1}, "NH", canvas, R|direction)

}

func DrawBensen(xy []int, canvas *svg.SVG, direction int) {
	var p []int
	y1 := direction&U/2 - 1
	xy = []int{xy[0] - 30, xy[1] - 15*y1}

	p = DrawLine(xy, canvas, R|(direction^(U|D)))
	canvas.Line(p[0]+2, p[1]-y1*6, p[0]+26, p[1]-y1*18, style)
	p = DrawLine(p, canvas, R|direction)
	p = DrawLine(p, canvas, direction)
	canvas.Line(p[0]-4, p[1]+y1*3, p[0]-28, p[1]-y1*9, style)
	p = DrawLine(p, canvas, L|direction)
	p = DrawLine(p, canvas, L|(direction^(U|D)))
	canvas.Line(p[0]+4, p[1]+y1*3, p[0]+4, p[1]+y1*27, style)
	p = DrawLine(p, canvas, (direction ^ (U | D)))
}

func newString() string {
	letters := "abcdefghijklmnoprstuvwxyz"
	var code string
	for i := 0; i < 32; i++ {
		code += string(letters[rand.Intn(len(letters))])
	}
	return code
}

package image

import (
	svg "github.com/ajstarks/svgo"
)

// Valine
func v(xy []int, canvas *svg.SVG, index int) []int {
	var p, out []int = aminobase(xy, canvas, index)
	if index&1 == 0 {
		p = DrawLine(p, canvas, D)
		DrawLine(p, canvas, DR)
		DrawLine(p, canvas, DL)
	} else {
		p = DrawLine(p, canvas, U)
		DrawLine(p, canvas, UR)
		DrawLine(p, canvas, UL)
	}
	return out
}

// Tyrosine
func y(xy []int, canvas *svg.SVG, index int) []int {
	var p, out []int = aminobase(xy, canvas, index)
	if index&1 == 0 {
		p = DrawLine(p, canvas, D)
		p = DrawLine(p, canvas, DR)
		DrawBensen([]int{p[0] + 30, p[1] - 15}, canvas, D)
		p = DrawLine([]int{p[0] + 60, p[1] + 30}, canvas, DR)
		DrawText(p, "OH", canvas, DR)
	} else {
		p = DrawLine(p, canvas, U)
		p = DrawLine(p, canvas, UR)
		DrawBensen([]int{p[0] + 30, p[1] + 15}, canvas, U)
		p = DrawLine([]int{p[0] + 60, p[1] - 30}, canvas, UR)
		DrawText(p, "OH", canvas, UR)
	}
	return out
}

// Tryptophan
func w(xy []int, canvas *svg.SVG, index int) []int {
	var p, out []int = aminobase(xy, canvas, index)
	if index&1 == 0 {
		p = DrawLine(p, canvas, D)
		p = DrawLine(p, canvas, DR)
		drawWThingy(p, canvas, D)
	} else {
		p = DrawLine(p, canvas, U)
		p = DrawLine(p, canvas, UR)
		drawWThingy(p, canvas, U)
	}
	return out
}

// Threonine
func t(xy []int, canvas *svg.SVG, index int) []int {
	var p, out []int = aminobase(xy, canvas, index)
	if index&1 == 0 {
		p = DrawLine(p, canvas, D)
		DrawLine(p, canvas, DL)
		p = DrawLine(p, canvas, DR)
		DrawText(p, "OH", canvas, DR)
	} else {
		p = DrawLine(p, canvas, U)
		DrawLine(p, canvas, UL)
		p = DrawLine(p, canvas, UR)
		DrawText(p, "OH", canvas, UR)
	}
	return out
}

// Serine
func s(xy []int, canvas *svg.SVG, index int) []int {
	var p, out []int = aminobase(xy, canvas, index)
	if index&1 == 0 {
		p = DrawLine(p, canvas, D)
		p = DrawLine(p, canvas, DR)
		DrawText(p, "OH", canvas, DR)
	} else {
		p = DrawLine(p, canvas, U)
		p = DrawLine(p, canvas, UR)
		DrawText(p, "OH", canvas, UR)
	}
	return out
}

// Proline
func p(xy []int, canvas *svg.SVG, index int) []int {
	var p, out []int = aminobase(xy, canvas, index)
	if index&1 == 0 {
		canvas.Line(p[0], p[1], p[0], p[1]+20, "stroke:black;stroke-width:2px")
		canvas.Line(p[0], p[1]+20, p[0]-20, p[1]+30, "stroke:black;stroke-width:2px")
		canvas.Line(p[0]-20, p[1]+30, p[0]-35, p[1]+10, "stroke:black;stroke-width:2px")
		canvas.Line(p[0]-35, p[1]+10, p[0]-30, p[1]-15, "stroke:black;stroke-width:2px")
	} else {
		canvas.Line(p[0], p[1], p[0], p[1]-20, "stroke:black;stroke-width:2px")
		canvas.Line(p[0], p[1]-20, p[0]-20, p[1]-30, "stroke:black;stroke-width:2px")
		canvas.Line(p[0]-20, p[1]-30, p[0]-35, p[1]-10, "stroke:black;stroke-width:2px")
		canvas.Line(p[0]-35, p[1]-10, p[0]-30, p[1]+15, "stroke:black;stroke-width:2px")
	}
	return out

}

// Phenylalanine
func f(xy []int, canvas *svg.SVG, index int) []int {
	var p, out []int = aminobase(xy, canvas, index)
	if index&1 == 0 {
		p = DrawLine(p, canvas, D)
		p = DrawLine(p, canvas, DR)
		DrawBensen([]int{p[0] + 30, p[1] - 15}, canvas, D)
	} else {
		p = DrawLine(p, canvas, U)
		p = DrawLine(p, canvas, UR)
		DrawBensen([]int{p[0] + 30, p[1] + 15}, canvas, U)
	}
	return out
}

// Methionine
func m(xy []int, canvas *svg.SVG, index int) []int {
	var p, out []int = aminobase(xy, canvas, index)
	if index&1 == 0 {
		p = DrawLine(p, canvas, D)
		p = DrawLine(p, canvas, DR)
		p = DrawLine(p, canvas, D)
		DrawText(p, "S", canvas, D)
		DrawLine(p, canvas, DR)
	} else {
		p = DrawLine(p, canvas, U)
		p = DrawLine(p, canvas, UR)
		p = DrawLine(p, canvas, U)
		DrawText(p, "S", canvas, U)
		DrawLine(p, canvas, UR)
	}
	return out
}

// Lysine
func k(xy []int, canvas *svg.SVG, index int) []int {
	var p, out []int = aminobase(xy, canvas, index)
	if index&1 == 0 {
		p = DrawLine(p, canvas, D)
		p = DrawLine(p, canvas, DR)
		p = DrawLine(p, canvas, D)
		p = DrawLine(p, canvas, DR)
		p = DrawLine(p, canvas, D)
		DrawText(p, "NH3", canvas, R)
		DrawText(p, "+", canvas, D)
	} else {
		p = DrawLine(p, canvas, U)
		p = DrawLine(p, canvas, UR)
		p = DrawLine(p, canvas, U)
		p = DrawLine(p, canvas, UR)
		p = DrawLine(p, canvas, U)
		DrawText(p, "NH3", canvas, R)
		DrawText(p, "+", canvas, U)
	}
	return out
}

// Leucine
func l(xy []int, canvas *svg.SVG, index int) []int {
	var p, out []int = aminobase(xy, canvas, index)
	if index&1 == 0 {
		p = DrawLine(p, canvas, D)
		p = DrawLine(p, canvas, DR)
		DrawLine(p, canvas, D)
		DrawLine(p, canvas, R)
	} else {
		p = DrawLine(p, canvas, U)
		p = DrawLine(p, canvas, UR)
		DrawLine(p, canvas, U)
		DrawLine(p, canvas, R)
	}
	return out
}

// Isoleucine
func i(xy []int, canvas *svg.SVG, index int) []int {
	var p, out []int = aminobase(xy, canvas, index)
	if index&1 == 0 {
		p = DrawLine(p, canvas, D)
		DrawDTriangle(p, canvas, DL, true)
		p = DrawLine(p, canvas, DR)
		DrawLine(p, canvas, D)
	} else {
		p = DrawLine(p, canvas, U)
		DrawTriangle(p, canvas, UL, true)
		p = DrawLine(p, canvas, UR)
		DrawLine(p, canvas, U)
	}
	return out
}

// Histidine
func h(xy []int, canvas *svg.SVG, index int) []int {
	var p, out []int = aminobase(xy, canvas, index)
	if index&1 == 0 {
		p = DrawLine(p, canvas, D)
		p = DrawLine(p, canvas, DR)
		drawHPentagon(p, canvas, D)
	} else {
		p = DrawLine(p, canvas, U)
		p = DrawLine(p, canvas, UR)
		drawHPentagon(p, canvas, U)
	}
	return out
}

// Glycine
func g(xy []int, canvas *svg.SVG, index int) []int {
	var p []int
	if index&1 == 0 {
		DrawText(xy, "NH", canvas, U)
		p = DrawLine(xy, canvas, DR)
		p = DrawLine(p, canvas, UR)
		DrawText(DrawDLine(p, canvas, U), "O", canvas, U)
		return DrawLine(p, canvas, DR)
	} else {
		DrawText(xy, "NH", canvas, D)
		p = DrawLine(xy, canvas, UR)
		p = DrawLine(p, canvas, DR)
		DrawText(DrawDLine(p, canvas, D), "O", canvas, D)
		return DrawLine(p, canvas, UR)
	}
}

// Glutamic Acid
func e(xy []int, canvas *svg.SVG, index int) []int {
	var p, out []int = aminobase(xy, canvas, index)
	if index&1 == 0 {
		p = DrawLine(p, canvas, D)
		p = DrawLine(p, canvas, DR)
		p = DrawLine(p, canvas, D)
		DrawText(DrawLine(p, canvas, DR), "O-", canvas, DR)
		DrawText(DrawDLine(p, canvas, DL), "O", canvas, DL)
	} else {
		p = DrawLine(p, canvas, U)
		p = DrawLine(p, canvas, UR)
		p = DrawLine(p, canvas, U)
		DrawText(DrawLine(p, canvas, UR), "O-", canvas, UR)
		DrawText(DrawDLine(p, canvas, UL), "O", canvas, UL)
	}
	return out
}

// Glutamine
func q(xy []int, canvas *svg.SVG, index int) []int {
	var p, out []int = aminobase(xy, canvas, index)
	if index&1 == 0 {
		p = DrawLine(p, canvas, D)
		p = DrawLine(p, canvas, DR)
		p = DrawLine(p, canvas, D)
		DrawText(DrawLine(p, canvas, DR), "NH2", canvas, DR)
		DrawText(DrawDLine(p, canvas, DL), "O", canvas, DL)
	} else {
		p = DrawLine(p, canvas, U)
		p = DrawLine(p, canvas, UR)
		p = DrawLine(p, canvas, U)
		DrawText(DrawLine(p, canvas, UR), "NH2", canvas, UR)
		DrawText(DrawDLine(p, canvas, UL), "O", canvas, UL)
	}
	return out
}

// Cysteine
func c(xy []int, canvas *svg.SVG, index int) []int {
	var p, out []int = aminobase(xy, canvas, index)
	if index&1 == 0 {
		p = DrawLine(p, canvas, D)
		DrawText(DrawLine(p, canvas, DR), "SH", canvas, DR)
	} else {
		p = DrawLine(p, canvas, U)
		DrawText(DrawLine(p, canvas, UR), "SH", canvas, UR)
	}
	return out
}

// Aspartic Acid
func d(xy []int, canvas *svg.SVG, index int) []int {
	var p, out []int = aminobase(xy, canvas, index)
	if index&1 == 0 {
		p = DrawLine(p, canvas, D)
		p = DrawLine(p, canvas, DR)
		DrawText(DrawLine(p, canvas, R), "O-", canvas, R)
		DrawText(DrawDLine(p, canvas, D), "O", canvas, D)
	} else {
		p = DrawLine(p, canvas, U)
		p = DrawLine(p, canvas, UR)
		DrawText(DrawLine(p, canvas, R), "O-", canvas, R)
		DrawText(DrawDLine(p, canvas, U), "O", canvas, U)
	}
	return out
}

// Aspargine
func n(xy []int, canvas *svg.SVG, index int) []int {
	var p, out []int = aminobase(xy, canvas, index)
	if index&1 == 0 {
		p = DrawLine(p, canvas, D)
		p = DrawLine(p, canvas, DR)
		DrawText(DrawLine(p, canvas, R), "NH2", canvas, R)
		DrawText(DrawDLine(p, canvas, D), "O", canvas, D)
	} else {
		p = DrawLine(p, canvas, U)
		p = DrawLine(p, canvas, UR)
		DrawText(DrawLine(p, canvas, R), "NH2", canvas, R)
		DrawText(DrawDLine(p, canvas, U), "O", canvas, U)
	}
	return out
}

// Arginine
func r(xy []int, canvas *svg.SVG, index int) []int {
	var p, out []int = aminobase(xy, canvas, index)
	if index&1 == 0 {
		p = DrawLine(p, canvas, D)
		p = DrawLine(p, canvas, DR)
		p = DrawLine(p, canvas, D)
		p = DrawLine(p, canvas, DR)
		DrawText(p, "NH", canvas, R)
		drawNitricCircle(p, canvas, D)
	} else {
		p = DrawLine(p, canvas, U)
		p = DrawLine(p, canvas, UR)
		p = DrawLine(p, canvas, U)
		p = DrawLine(p, canvas, UR)
		DrawText(p, "NH", canvas, R)
		drawNitricCircle(p, canvas, U)
	}
	return out
}

// Alanine
func a(xy []int, canvas *svg.SVG, index int, end bool) []int {
	var p []int
	if index&1 == 0 {
		DrawText(xy, "NH", canvas, U)
		p = DrawLine(xy, canvas, DR)
		if end {
			DrawDTriangle(p, canvas, D, true)
		} else {
			DrawTriangle(p, canvas, D, true)
		}
		p = DrawLine(p, canvas, UR)
		DrawText(DrawDLine(p, canvas, U), "O", canvas, U)
		return DrawLine(p, canvas, DR)
	} else {
		DrawText(xy, "NH", canvas, D)
		p = DrawLine(xy, canvas, UR)
		if end {
			DrawDTriangle(p, canvas, U, true)
		} else {
			DrawTriangle(p, canvas, U, true)
		}
		p = DrawLine(p, canvas, DR)
		DrawText(DrawDLine(p, canvas, D), "O", canvas, D)
		return DrawLine(p, canvas, UR)
	}
}

// Draws basic aminoacid with triangle between NH and rest of acid
// returns coordinates for rest of acid and end of acid
func aminobase(xy []int, canvas *svg.SVG, index int) ([]int, []int) {
	var p, pp []int
	if index&1 == 0 {
		DrawText(xy, "NH", canvas, U)
		p = DrawDTriangle(xy, canvas, DR, false)
		pp = p
		p = DrawLine(p, canvas, UR)
		DrawText(DrawDLine(p, canvas, U), "O", canvas, U)
		return pp, DrawLine(p, canvas, DR)
	} else {
		DrawText(xy, "NH", canvas, D)
		p = DrawTriangle(xy, canvas, UR, false)
		pp = p
		p = DrawLine(p, canvas, DR)
		DrawText(DrawDLine(p, canvas, D), "O", canvas, D)
		return pp, DrawLine(p, canvas, UR)
	}
}

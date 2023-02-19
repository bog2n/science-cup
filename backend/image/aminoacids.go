package image

import (
	svg "github.com/ajstarks/svgo"
)

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

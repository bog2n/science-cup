package image

import (
	svg "github.com/ajstarks/svgo"
)

func c(xy []int, canvas *svg.SVG, index int) []int {
	var p, pp []int
	if index&1 == 0 {
		DrawText(xy, "NH", canvas, U)
		p = DrawDTriangle(xy, canvas, DR, false)
		pp = p
		// Acidic rest
		pp = DrawLine(pp, canvas, D)
		DrawText(DrawLine(pp, canvas, DR), "SH", canvas, DR)

		p = DrawLine(p, canvas, UR)
		DrawText(DrawDLine(p, canvas, U), "O", canvas, U)
		return DrawLine(p, canvas, DR)
	} else {
		DrawText(xy, "NH", canvas, D)
		p = DrawTriangle(xy, canvas, UR, false)
		pp = p
		// Acidic rest
		pp = DrawLine(pp, canvas, U)
		DrawText(DrawLine(pp, canvas, UR), "SH", canvas, UR)

		p = DrawLine(p, canvas, DR)
		DrawText(DrawDLine(p, canvas, D), "O", canvas, D)
		return DrawLine(p, canvas, UR)
	}
}

func d(xy []int, canvas *svg.SVG, index int) []int {
	var p, pp []int
	if index&1 == 0 {
		DrawText(xy, "NH", canvas, U)
		p = DrawDTriangle(xy, canvas, DR, false)
		pp = p
		// Acidic rest
		pp = DrawLine(pp, canvas, D)
		pp = DrawLine(pp, canvas, DR)
		DrawText(DrawLine(pp, canvas, R), "O-", canvas, R)
		DrawText(DrawDLine(pp, canvas, D), "O", canvas, D)

		p = DrawLine(p, canvas, UR)
		DrawText(DrawDLine(p, canvas, U), "O", canvas, U)
		return DrawLine(p, canvas, DR)
	} else {
		DrawText(xy, "NH", canvas, D)
		p = DrawTriangle(xy, canvas, UR, false)
		pp = p
		// Acidic rest
		pp = DrawLine(pp, canvas, U)
		pp = DrawLine(pp, canvas, UR)
		DrawText(DrawLine(pp, canvas, R), "O-", canvas, R)
		DrawText(DrawDLine(pp, canvas, U), "O", canvas, U)

		p = DrawLine(p, canvas, DR)
		DrawText(DrawDLine(p, canvas, D), "O", canvas, D)
		return DrawLine(p, canvas, UR)
	}
}

func n(xy []int, canvas *svg.SVG, index int) []int {
	var p, pp []int
	if index&1 == 0 {
		DrawText(xy, "NH", canvas, U)
		p = DrawDTriangle(xy, canvas, DR, false)
		pp = p
		// Acidic rest
		pp = DrawLine(pp, canvas, D)
		pp = DrawLine(pp, canvas, DR)
		DrawText(DrawLine(pp, canvas, R), "NH2", canvas, R)
		DrawText(DrawDLine(pp, canvas, D), "O", canvas, D)

		p = DrawLine(p, canvas, UR)
		DrawText(DrawDLine(p, canvas, U), "O", canvas, U)
		return DrawLine(p, canvas, DR)
	} else {
		DrawText(xy, "NH", canvas, D)
		p = DrawTriangle(xy, canvas, UR, false)
		pp = p
		// Acidic rest
		pp = DrawLine(pp, canvas, U)
		pp = DrawLine(pp, canvas, UR)
		DrawText(DrawLine(pp, canvas, R), "NH2", canvas, R)
		DrawText(DrawDLine(pp, canvas, U), "O", canvas, U)

		p = DrawLine(p, canvas, DR)
		DrawText(DrawDLine(p, canvas, D), "O", canvas, D)
		return DrawLine(p, canvas, UR)
	}
}

func r(xy []int, canvas *svg.SVG, index int) []int {
	var p, pp []int
	if index&1 == 0 {
		DrawText(xy, "NH", canvas, U)
		p = DrawDTriangle(xy, canvas, DR, false)
		pp = p
		// Acidic rest
		pp = DrawLine(pp, canvas, D)
		pp = DrawLine(pp, canvas, DR)
		pp = DrawLine(pp, canvas, D)
		pp = DrawLine(pp, canvas, DR)
		DrawText(pp, "NH", canvas, R)
		drawNitricCircle(pp, canvas, D)

		p = DrawLine(p, canvas, UR)
		DrawText(DrawDLine(p, canvas, U), "O", canvas, U)
		return DrawLine(p, canvas, DR)
	} else {
		DrawText(xy, "NH", canvas, D)
		p = DrawTriangle(xy, canvas, UR, false)
		pp = p
		// Acidic rest
		pp = DrawLine(pp, canvas, U)
		pp = DrawLine(pp, canvas, UR)
		pp = DrawLine(pp, canvas, U)
		pp = DrawLine(pp, canvas, UR)
		DrawText(pp, "NH", canvas, R)
		drawNitricCircle(pp, canvas, U)

		p = DrawLine(p, canvas, DR)
		DrawText(DrawDLine(p, canvas, D), "O", canvas, D)
		return DrawLine(p, canvas, UR)
	}
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

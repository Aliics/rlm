package rlm

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"math"
)

type Vec2 rl.Vector2

var (
	Vec2Zero     = Vec2{}
	Vec2Identity = Vec2{}
	Vec2XY       = Vec2{X: 1, Y: 1}
	Vec2Up       = Vec2{Y: 1}
	Vec2Down     = Vec2{Y: -1}
	Vec2Left     = Vec2{X: 1}
	Vec2Right    = Vec2{X: -1}
)

func NewVec2(x, y float32) Vec2 { return Vec2{X: x, Y: y} }
func NewVec2N(v ...float32) Vec2 {
	var (
		result Vec2
		prev   float32
	)
	if len(v) > 0 {
		result.X = v[0]
		prev = result.X
	}
	if len(v) > 1 {
		result.Y = v[1]
	} else {
		result.Y = prev
	}
	return result
}

func (v Vec2) XY() (x, y float32)     { return v.X, v.Y }
func (v Vec2) YX() (y, x float32)     { return v.Y, v.X }
func (v Vec2) XX() (float32, float32) { return v.X, v.X }
func (v Vec2) YY() (float32, float32) { return v.Y, v.Y }

func (v Vec2) Floats() []float32  { return []float32{v.X, v.Y} }
func (v Vec2) Float2() [2]float32 { return [...]float32{v.X, v.Y} }

func (v Vec2) ToVec3(z float32) Vec3 { return Vec3{X: v.X, Y: v.Y, Z: z} }

// Equals compares this and another Vec2 taking into account floating point precision.
func (v Vec2) Equals(v2 Vec2) bool      { return rl.Vector2Equals(rl.Vector2(v), rl.Vector2(v2)) }
func (v Vec2) LessThan(v2 Vec2) bool    { return v.X < v2.Y && v.Y < v2.Y }
func (v Vec2) GreaterThan(v2 Vec2) bool { return v.X > v2.Y && v.Y > v2.Y }

// Abs will create a new Vec2 where all axes are their absolute values.
//
//	(abs(x), abs(y))
func (v Vec2) Abs() Vec2 {
	return Vec2{
		X: float32(math.Abs(float64(v.X))),
		Y: float32(math.Abs(float64(v.Y))),
	}
}

// Add will add this Vec2's and another one's respective axes to form a third Vec2 where each axis is the sum of the
// two Vec2's respective axes.
//
//	(x1+x2, y1+y2)
func (v Vec2) Add(v2 Vec2) Vec2 { return Vec2(rl.Vector2Add(rl.Vector2(v), rl.Vector2(v2))) }

// Subtract will subtract this Vec2's and another one's respective axes to form a third Vec2 where each axis is the
// difference of the two Vec2's respective axes.
//
//	(x1-x2, y1-y2)
func (v Vec2) Subtract(v2 Vec2) Vec2 { return Vec2(rl.Vector2Subtract(rl.Vector2(v), rl.Vector2(v2))) }

// Multiply will multiply this Vec2's and another one's respective axes to form a third Vec2 where each axis is the
// product of the two Vec2's respective axes.
//
//	(x1*x2, y1*y2)
func (v Vec2) Multiply(v2 Vec2) Vec2 { return Vec2(rl.Vector2Multiply(rl.Vector2(v), rl.Vector2(v2))) }

// Divide will divide this Vec2's and another one's respective axes to form a third Vec2 where each axis is the
// product of the two Vec2's respective axes.
//
//	(x1/x2, y1/y2)
func (v Vec2) Divide(v2 Vec2) Vec2 { return Vec2(rl.Vector2Divide(rl.Vector2(v), rl.Vector2(v2))) }

// Max will determine the max for each of the axes.
func (v Vec2) Max(v2 Vec2) Vec2 { return Vec2{X: max(v.X, v2.X), Y: max(v.Y, v2.Y)} }
func (v Vec2) MaxN(vs ...Vec2) Vec2 {
	result := v
	for _, v2 := range vs {
		result = v.Max(v2)
	}
	return result
}

// Min will determine the min for each of the axes.
func (v Vec2) Min(v2 Vec2) Vec2 { return Vec2{X: min(v.X, v2.X), Y: min(v.Y, v2.Y)} }
func (v Vec2) MinN(vs ...Vec2) Vec2 {
	result := v
	for _, v2 := range vs {
		result = v.Min(v2)
	}
	return result
}

// Clamp will provide a Vec2 whose axes are no larger than the respective axes of max and are no smaller than the
// respective axes of min.
func (v Vec2) Clamp(min, max Vec2) Vec2 {
	return Vec2(rl.Vector2Clamp(rl.Vector2(v), rl.Vector2(min), rl.Vector2(max)))
}

// Dot will calculate the dot product of this Vec32 and another.
//
//	x1*x2 + y1*y2
func (v Vec2) Dot(v2 Vec2) float32 { return rl.Vector2DotProduct(rl.Vector2(v), rl.Vector2(v2)) }

// Distance will calculate the distance between this Vec2 and another.
// This value is effectively just pythagoras.
//
//	sqrt(dX^2, dY^2)
func (v Vec2) Distance(v2 Vec2) float32 { return rl.Vector2Distance(rl.Vector2(v), rl.Vector2(v2)) }

// Scale will scale the all axes by the given float.
func (v Vec2) Scale(s float32) Vec2 { return Vec2(rl.Vector2Scale(rl.Vector2(v), s)) }

// Length is the same as Distance from (0, 0).
// This value is effectively just pythagoras.
//
//	sqrt(x^2, y^2)
func (v Vec2) Length() float32 { return rl.Vector2Length(rl.Vector2(v)) }
func (v Vec2) Normalize() Vec2 { return Vec2(rl.Vector2Normalize(rl.Vector2(v))) }
func (v Vec2) Negate() Vec2    { return Vec2(rl.Vector2Negate(rl.Vector2(v))) }
func (v Vec2) Invert() Vec2    { return Vec2(rl.Vector2Invert(rl.Vector2(v))) }

// Lerp linearly interpolate a Vec2 between this Vec2 and another.
func (v Vec2) Lerp(v2 Vec2, x float32) Vec2 {
	return Vec2(rl.Vector2Lerp(rl.Vector2(v), rl.Vector2(v2), x))
}

func (v Vec2) CrossProduct(v2 Vec2) float32 {
	return rl.Vector2CrossProduct(rl.Vector2(v), rl.Vector2(v2))
}

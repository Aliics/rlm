package rlm

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"math"
)

type Vec3 rl.Vector3

var (
	Vec3Zero     Vec3 = Vec3{}
	Vec3Identity Vec3 = Vec3{}
	Vec3XYZ      Vec3 = Vec3{X: 1, Y: 1, Z: 1}
	Vec3XY       Vec3 = Vec3{X: 1, Y: 1}
	Vec3YZ       Vec3 = Vec3{Y: 1, Z: 1}
	Vec3XZ       Vec3 = Vec3{X: 1, Z: 1}
	Vec3Up       Vec3 = Vec3{Y: 1}
	Vec3Down     Vec3 = Vec3{Y: -1}
	Vec3Left     Vec3 = Vec3{X: 1}
	Vec3Right    Vec3 = Vec3{X: -1}
	Vec3Front    Vec3 = Vec3{Z: 1}
	Vec3Back     Vec3 = Vec3{Z: -1}
)

func NewVec3(x, y, z float32) Vec3 { return Vec3{X: x, Y: y, Z: z} }
func NewVec3N(v ...float32) Vec3 {
	var (
		result Vec3
		prev   float32
	)
	if len(v) > 0 {
		result.X = v[0]
		prev = result.X
	}
	if len(v) > 1 {
		result.Y = v[1]
		prev = result.Y
	} else {
		result.Y = prev
	}
	if len(v) > 2 {
		result.Z = v[2]
	} else {
		result.Z = prev
	}
	return result
}

func (v Vec3) XXX() (float32, float32, float32)         { return v.X, v.X, v.X }
func (v Vec3) XXY() (x1 float32, x2 float32, y float32) { return v.X, v.X, v.Y }
func (v Vec3) XYX() (x1 float32, y float32, x2 float32) { return v.X, v.Y, v.X }
func (v Vec3) YXX() (y float32, x1 float32, x2 float32) { return v.Y, v.X, v.X }
func (v Vec3) XXZ() (x1 float32, x2 float32, z float32) { return v.X, v.X, v.Z }
func (v Vec3) XZX() (x1 float32, z float32, x2 float32) { return v.X, v.Z, v.X }
func (v Vec3) ZXX() (z float32, x1 float32, x2 float32) { return v.Z, v.X, v.X }

func (v Vec3) YYY() (float32, float32, float32)         { return v.Y, v.Y, v.Y }
func (v Vec3) YYX() (y1 float32, y2 float32, x float32) { return v.Y, v.Y, v.X }
func (v Vec3) YXY() (y1 float32, x float32, y2 float32) { return v.Y, v.X, v.Y }
func (v Vec3) XYY() (x float32, y1 float32, y2 float32) { return v.X, v.Y, v.Y }
func (v Vec3) YYZ() (y1 float32, y2 float32, z float32) { return v.Y, v.Y, v.Z }
func (v Vec3) YZY() (y1 float32, z float32, y2 float32) { return v.Y, v.Z, v.Y }
func (v Vec3) ZYY() (z float32, y1 float32, y2 float32) { return v.Z, v.Y, v.Y }

func (v Vec3) ZZZ() (float32, float32, float32)         { return v.Z, v.Z, v.Z }
func (v Vec3) ZZX() (z1 float32, z2 float32, x float32) { return v.Z, v.Z, v.X }
func (v Vec3) ZXZ() (z1 float32, x float32, z2 float32) { return v.Z, v.X, v.Z }
func (v Vec3) XZZ() (x float32, z1 float32, z2 float32) { return v.X, v.Z, v.Z }
func (v Vec3) ZZY() (z1 float32, z2 float32, y float32) { return v.Z, v.Z, v.Y }
func (v Vec3) ZYZ() (z1 float32, y float32, z2 float32) { return v.Z, v.Y, v.Z }
func (v Vec3) YZZ() (y float32, z1 float32, z2 float32) { return v.Y, v.Z, v.Z }

func (v Vec3) XYZ() (x, y, z float32) { return v.X, v.Y, v.Z }
func (v Vec3) XZY() (x, z, y float32) { return v.X, v.Z, v.Y }
func (v Vec3) YZX() (y, z, x float32) { return v.Y, v.Z, v.X }
func (v Vec3) YXZ() (y, x, z float32) { return v.Y, v.X, v.Z }
func (v Vec3) ZYX() (z, y, x float32) { return v.Z, v.Y, v.X }
func (v Vec3) ZXY() (z, x, y float32) { return v.Z, v.X, v.Y }

func (v Vec3) XY() (x, y float32)     { return v.X, v.Y }
func (v Vec3) YX() (y, x float32)     { return v.Y, v.X }
func (v Vec3) YZ() (y, z float32)     { return v.Y, v.Z }
func (v Vec3) ZY() (z, y float32)     { return v.Z, v.Y }
func (v Vec3) XZ() (x, z float32)     { return v.X, v.Z }
func (v Vec3) ZX() (z, x float32)     { return v.Z, v.X }
func (v Vec3) XX() (float32, float32) { return v.X, v.X }
func (v Vec3) YY() (float32, float32) { return v.Y, v.Y }
func (v Vec3) ZZ() (float32, float32) { return v.Z, v.Z }

// Equals compares this and another Vec3 taking into account floating point precision.
func (v Vec3) Equals(v3 Vec3) bool      { return rl.Vector3Equals(rl.Vector3(v), rl.Vector3(v3)) }
func (v Vec3) LessThan(v3 Vec3) bool    { return v.X < v3.Y && v.Y < v3.Y && v.Z < v3.Z }
func (v Vec3) GreaterThan(v3 Vec3) bool { return v.X > v3.Y && v.Y > v3.Y && v.Z > v3.Z }

// Abs will create a new Vec3 where all axes are their absolute values.
//
//	(abs(x), abs(y), abs(z))
func (v Vec3) Abs() Vec3 {
	return Vec3{
		X: float32(math.Abs(float64(v.X))),
		Y: float32(math.Abs(float64(v.Y))),
		Z: float32(math.Abs(float64(v.Z))),
	}
}

// Add will add this Vec3's and another one's respective axes to form a third Vec3 where each axis is the sum of the
// two Vec3's respective axes.
//
//	(x1+x2, y1+y2, z1+z2)
func (v Vec3) Add(v3 Vec3) Vec3 { return Vec3(rl.Vector3Add(rl.Vector3(v), rl.Vector3(v3))) }

// Subtract will subtract this Vec3's and another one's respective axes to form a third Vec3 where each axis is the
// difference of the two Vec3's respective axes.
//
//	(x1-x2, y1-y2, z1-z2)
func (v Vec3) Subtract(v3 Vec3) Vec3 { return Vec3(rl.Vector3Subtract(rl.Vector3(v), rl.Vector3(v3))) }

// Multiply will multiply this Vec3's and another one's respective axes to form a third Vec3 where each axis is the
// product of the two Vec3's respective axes.
//
//	(x1*x2, y1*y2, z1*z2)
func (v Vec3) Multiply(v3 Vec3) Vec3 { return Vec3(rl.Vector3Multiply(rl.Vector3(v), rl.Vector3(v3))) }

// Divide will divide this Vec3's and another one's respective axes to form a third Vec3 where each axis is the
// product of the two Vec3's respective axes.
//
//	(x1/x2, y1/y2, z1/z2)
func (v Vec3) Divide(v3 Vec3) Vec3 { return Vec3(rl.Vector3Divide(rl.Vector3(v), rl.Vector3(v3))) }

// Max will determine the max for each of the axes.
func (v Vec3) Max(v3 Vec3) Vec3 { return Vec3(rl.Vector3Max(rl.Vector3(v), rl.Vector3(v3))) }
func (v Vec3) MaxN(vs ...Vec3) Vec3 {
	result := v
	for _, v3 := range vs {
		result = v.Max(v3)
	}
	return result
}

// Min will determine the min for each of the axes.
func (v Vec3) Min(v3 Vec3) Vec3 { return Vec3(rl.Vector3Min(rl.Vector3(v), rl.Vector3(v3))) }
func (v Vec3) MinN(vs ...Vec3) Vec3 {
	result := v
	for _, v3 := range vs {
		result = v.Min(v3)
	}
	return result
}

// Clamp will provide a Vec3 whose axes are no larger than the respective axes of max and are no smaller than the
// respective axes of min.
func (v Vec3) Clamp(min, max Vec3) Vec3 {
	return Vec3(rl.Vector3Clamp(rl.Vector3(v), rl.Vector3(min), rl.Vector3(max)))
}

// Dot will calculate the dot product of this Vec3 and another.
//
//	x1*x2 + y1*y2 + z1*z2
func (v Vec3) Dot(v3 Vec3) float32 { return rl.Vector3DotProduct(rl.Vector3(v), rl.Vector3(v3)) }

// Distance will calculate the distance between this Vec3 and another.
// This value is effectively just pythagoras in 3 dimensions.
//
//	sqrt(dX^2, dY^2, dZ^2)
func (v Vec3) Distance(v3 Vec3) float32 { return rl.Vector3Distance(rl.Vector3(v), rl.Vector3(v3)) }

// Scale will scale the all axes by the given float.
func (v Vec3) Scale(s float32) Vec3 { return Vec3(rl.Vector3Scale(rl.Vector3(v), s)) }

// Length is the same as Distance from (0, 0, 0).
// This value is effectively just pythagoras in 3 dimensions.
//
//	sqrt(x^2, y^2, z^2)
func (v Vec3) Length() float32 { return rl.Vector3Length(rl.Vector3(v)) }
func (v Vec3) Normalize() Vec3 { return Vec3(rl.Vector3Normalize(rl.Vector3(v))) }
func (v Vec3) Negate() Vec3    { return Vec3(rl.Vector3Negate(rl.Vector3(v))) }
func (v Vec3) Invert() Vec3    { return Vec3(rl.Vector3Invert(rl.Vector3(v))) }

// Lerp linearly interpolate a Vec3 between this Vec3 and another.
func (v Vec3) Lerp(v3 Vec3, x float32) Vec3 {
	return Vec3(rl.Vector3Lerp(rl.Vector3(v), rl.Vector3(v3), x))
}

// CrossProduct will calculate the cross between this Vec3 and another.
func (v Vec3) CrossProduct(v3 Vec3) Vec3 {
	return Vec3(rl.Vector3CrossProduct(rl.Vector3(v), rl.Vector3(v3)))
}

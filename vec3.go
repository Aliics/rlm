package rlm

import rl "github.com/gen2brain/raylib-go/raylib"

type Vec3 rl.Vector3

var (
	Vec3Zero  Vec3 = Vec3{}
	Vec3XYZ   Vec3 = Vec3{X: 1, Y: 1, Z: 1}
	Vec3XY    Vec3 = Vec3{X: 1, Y: 1}
	Vec3YZ    Vec3 = Vec3{Y: 1, Z: 1}
	Vec3XZ    Vec3 = Vec3{X: 1, Z: 1}
	Vec3Up    Vec3 = Vec3{Y: 1}
	Vec3Down  Vec3 = Vec3{Y: -1}
	Vec3Left  Vec3 = Vec3{X: 1}
	Vec3Right Vec3 = Vec3{X: -1}
	Vec3Front Vec3 = Vec3{Z: 1}
	Vec3Back  Vec3 = Vec3{Z: -1}
)

func NewVec3(x, y, z float32) Vec3 { return Vec3{X: x, Y: y, Z: z} }
func NewVec3N(v float32) Vec3      { return Vec3{X: v, Y: v, Z: v} }

func (v Vec3) XXX() (float32, float32, float32)         { return v.X, v.X, v.X }
func (v Vec3) XXY() (x1 float32, x2 float32, y float32) { return v.X, v.X, v.Y }
func (v Vec3) XYX() (x1 float32, y float32, x2 float32) { return v.X, v.Y, v.X }
func (v Vec3) YXX() (y float32, x1 float32, x2 float32) { return v.Y, v.X, v.X }
func (v Vec3) XXZ() (x1 float32, x2 float32, z float32) { return v.X, v.X, v.Z }
func (v Vec3) XZX() (x1 float32, z float32, x2 float32) { return v.X, v.Z, v.X }
func (v Vec3) ZXX() (z float32, x1 float32, x2 float32) { return v.Z, v.X, v.X }

func (v Vec3) YYY() (float32, float32, float32)         { return v.Y, v.Y, v.Y }
func (v Vec3) YYX() (y1 float32, y2 float32, y float32) { return v.Y, v.Y, v.X }
func (v Vec3) YXY() (y1 float32, y float32, y2 float32) { return v.Y, v.X, v.Y }
func (v Vec3) XYY() (y float32, y1 float32, y2 float32) { return v.X, v.Y, v.Y }
func (v Vec3) YYZ() (y1 float32, y2 float32, z float32) { return v.Y, v.Y, v.Z }
func (v Vec3) YZY() (y1 float32, z float32, y2 float32) { return v.Y, v.Z, v.Y }
func (v Vec3) ZYY() (z float32, y1 float32, y2 float32) { return v.Z, v.Y, v.Y }

func (v Vec3) ZZZ() (float32, float32, float32)         { return v.Z, v.Z, v.Z }
func (v Vec3) ZZX() (z1 float32, z2 float32, Z float32) { return v.Z, v.Z, v.X }
func (v Vec3) ZXZ() (z1 float32, Z float32, z2 float32) { return v.Z, v.X, v.Z }
func (v Vec3) XZZ() (Z float32, z1 float32, z2 float32) { return v.X, v.Z, v.Z }
func (v Vec3) ZZY() (z1 float32, z2 float32, Z float32) { return v.Z, v.Z, v.Y }
func (v Vec3) ZYZ() (z1 float32, Z float32, z2 float32) { return v.Z, v.Y, v.Z }
func (v Vec3) YZZ() (z float32, z1 float32, z2 float32) { return v.Y, v.Z, v.Z }

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

func (v Vec3) Add(v3 Vec3) Vec3         { return Vec3(rl.Vector3Add(rl.Vector3(v), rl.Vector3(v3))) }
func (v Vec3) Subtract(v3 Vec3) Vec3    { return Vec3(rl.Vector3Subtract(rl.Vector3(v), rl.Vector3(v3))) }
func (v Vec3) Multiply(v3 Vec3) Vec3    { return Vec3(rl.Vector3Multiply(rl.Vector3(v), rl.Vector3(v3))) }
func (v Vec3) Divide(v3 Vec3) Vec3      { return Vec3(rl.Vector3Divide(rl.Vector3(v), rl.Vector3(v3))) }
func (v Vec3) Max(v3 Vec3) Vec3         { return Vec3(rl.Vector3Max(rl.Vector3(v), rl.Vector3(v3))) }
func (v Vec3) Min(v3 Vec3) Vec3         { return Vec3(rl.Vector3Min(rl.Vector3(v), rl.Vector3(v3))) }
func (v Vec3) Distance(v3 Vec3) float32 { return rl.Vector3Distance(rl.Vector3(v), rl.Vector3(v3)) }
func (v Vec3) Dot(v3 Vec3) float32      { return rl.Vector3DotProduct(rl.Vector3(v), rl.Vector3(v3)) }
func (v Vec3) Scale(s float32) Vec3     { return Vec3(rl.Vector3Scale(rl.Vector3(v), s)) }
func (v Vec3) Normalize() Vec3          { return Vec3(rl.Vector3Normalize(rl.Vector3(v))) }
func (v Vec3) CrossProduct(v3 Vec3) Vec3 {
	return Vec3(rl.Vector3CrossProduct(rl.Vector3(v), rl.Vector3(v3)))
}

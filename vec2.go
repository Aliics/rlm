package rlm

import rl "github.com/gen2brain/raylib-go/raylib"

type Vec2 rl.Vector2

var (
	Vec2Zero  Vec2 = Vec2{}
	Vec2XY    Vec2 = Vec2{X: 1, Y: 1}
	Vec2Up    Vec2 = Vec2{Y: 1}
	Vec2Down  Vec2 = Vec2{Y: -1}
	Vec2Left  Vec2 = Vec2{X: 1}
	Vec2Right Vec2 = Vec2{X: -1}
)

func NewVec2(x, y float32) Vec2 { return Vec2{X: x, Y: y} }
func NewVec2N(v float32) Vec2   { return Vec2{X: v, Y: v} }

func (v Vec2) XY() (x, y float32)     { return v.X, v.Y }
func (v Vec2) YX() (y, x float32)     { return v.Y, v.X }
func (v Vec2) XX() (float32, float32) { return v.X, v.X }
func (v Vec2) YY() (float32, float32) { return v.Y, v.Y }

func (v Vec2) ToVec3(z float32) Vec3 { return Vec3{X: v.X, Y: v.Y, Z: z} }

func (v Vec2) Add(v2 Vec2) Vec2         { return Vec2(rl.Vector2Add(rl.Vector2(v), rl.Vector2(v2))) }
func (v Vec2) Subtract(v2 Vec2) Vec2    { return Vec2(rl.Vector2Subtract(rl.Vector2(v), rl.Vector2(v2))) }
func (v Vec2) Multiply(v2 Vec2) Vec2    { return Vec2(rl.Vector2Multiply(rl.Vector2(v), rl.Vector2(v2))) }
func (v Vec2) Divide(v2 Vec2) Vec2      { return Vec2(rl.Vector2Divide(rl.Vector2(v), rl.Vector2(v2))) }
func (v Vec2) Max(v2 Vec2) Vec2         { return Vec2{X: max(v.X, v2.X), Y: max(v.Y, v2.Y)} }
func (v Vec2) Min(v2 Vec2) Vec2         { return Vec2{X: min(v.X, v2.X), Y: min(v.Y, v2.Y)} }
func (v Vec2) Dot(v2 Vec2) float32      { return rl.Vector2DotProduct(rl.Vector2(v), rl.Vector2(v2)) }
func (v Vec2) Distance(v2 Vec2) float32 { return rl.Vector2Distance(rl.Vector2(v), rl.Vector2(v2)) }
func (v Vec2) Scale(s float32) Vec2     { return Vec2(rl.Vector2Scale(rl.Vector2(v), s)) }
func (v Vec2) Normalize() Vec2          { return Vec2(rl.Vector2Normalize(rl.Vector2(v))) }
func (v Vec2) CrossProduct(v2 Vec2) float32 {
	return rl.Vector2CrossProduct(rl.Vector2(v), rl.Vector2(v2))
}

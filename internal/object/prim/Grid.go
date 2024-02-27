package prim

import (
	"karalis/internal/camera"
	pub_object "karalis/pkg/object"

	raylib "github.com/gen2brain/raylib-go/raylib"
)

var ()

type Grid struct {
	spacing float32
	size    int32
}

func (g *Grid) Init() error {
	g.spacing = 1
	g.size = 10

	return nil
}

func (c *Grid) GetModelMatrix() raylib.Matrix {
	return raylib.MatrixIdentity()
}

func (g *Grid) GetPos() raylib.Vector3 {
	return raylib.NewVector3(0, 0, 0)
}

func (c *Grid) GetPitch() float32 {
	return 0
}

func (c *Grid) SetPitch(p float32) {}

func (c *Grid) GetYaw() float32 {
	return 0
}

func (c *Grid) SetYaw(y float32) {}

func (c *Grid) GetRoll() float32 {
	return 0
}

func (c *Grid) SetRoll(r float32) {}

func (g *Grid) GetVertices() []raylib.Vector3 {
	verts := []raylib.Vector3{}
	return verts
}

func (g *Grid) GetUVs() []raylib.Vector2 {
	uvs := []raylib.Vector2{}
	return uvs
}

func (g *Grid) SetUVs(uvs []raylib.Vector2) {
}

func (c *Grid) GetMaterials() *raylib.Material {
	return &raylib.Material{}
}

func (c *Grid) SetTexture(mat *raylib.Material, tex raylib.Texture2D) {
}

func (c *Grid) GetTexture(mat *raylib.Material) raylib.Texture2D {
	return raylib.Texture2D{}
}

func (g *Grid) Prerender(cam *camera.Cam) []func() {
	return []func(){}
}

func (g *Grid) Render(cam *camera.Cam) []func() {
	raylib.DrawGrid(g.size, g.spacing)
	return []func(){}
}

func (g *Grid) Postrender(cam *camera.Cam) []func() {
	return []func(){}
}

func (g *Grid) Update(dt float32) {
}

func (g *Grid) OnAdd() {
}

func (g *Grid) OnRemove() {
}

func (g *Grid) AddChild(obj pub_object.Object) {
}

func (g *Grid) RemChild(obj pub_object.Object) {
}

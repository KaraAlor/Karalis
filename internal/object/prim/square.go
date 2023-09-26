package prim

import (
	"image/color"

	pub_object "karalis/pkg/object"

	raylib "github.com/gen2brain/raylib-go/raylib"
)

var ()

type Square struct {
	mdl   raylib.Model
	pos   raylib.Vector3
	size  float32
	color color.RGBA
}

func (c *Square) Init() {
	c.pos = raylib.NewVector3(0, 0, 0)
	c.size = 1
	c.color = raylib.White

	c.mdl = raylib.LoadModel("res/prim/square.obj")
}

func (c *Square) GetMaterials() *raylib.Material {
	return c.mdl.Materials
}

func (c *Square) SetTexture(mat *raylib.Material, tex raylib.Texture2D) {
	raylib.SetMaterialTexture(mat, raylib.MapDiffuse, tex)
}

func (c *Square) GetTexture(mat *raylib.Material) raylib.Texture2D {
	return mat.Maps.Texture
}

func (c *Square) Prerender() []func() {
	return []func(){}
}

func (c *Square) Render() []func() {
	raylib.SetTexture(c.mdl.Materials.Maps.Texture.ID)
	raylib.DrawModel(c.mdl, c.pos, c.size, c.color)
	raylib.SetTexture(0)
	return []func(){}
}

func (c *Square) Postrender() []func() {
	return []func(){}
}

func (c *Square) Update(dt float32) {
}

func (c *Square) OnAdd() {
}

func (c *Square) OnRemove() {
}

func (c *Square) AddChild(obj pub_object.Object) {
}

func (c *Square) RemChild(obj pub_object.Object) {
}
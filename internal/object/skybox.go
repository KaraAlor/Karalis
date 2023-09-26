package object

import (
	pub_object "godev/pkg/object"

	raylib "github.com/gen2brain/raylib-go/raylib"
)

var ()

type Skybox struct {
}

func (s *Skybox) Init() {
}

func (s *Skybox) Prerender() []func() {
	return []func(){}
}

func (s *Skybox) Render() []func() {
	return []func(){}
}

func (s *Skybox) Postrender() []func() {
	return []func(){}
}

func (s *Skybox) Update(dt float32) {
}

func (s *Skybox) OnAdd() {
}

func (s *Skybox) OnRemove() {
}

func (s *Skybox) AddChild(obj pub_object.Object) {
}

func (s *Skybox) RemChild(obj pub_object.Object) {
}

func (s *Skybox) GetMaterials() *raylib.Material {
	return &raylib.Material{}
}

func (s *Skybox) SetTexture(mat *raylib.Material, tex raylib.Texture2D) {
}

func (s *Skybox) GetTexture(mat *raylib.Material) raylib.Texture2D {
	return raylib.Texture2D{}
}

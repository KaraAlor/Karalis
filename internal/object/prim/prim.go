package prim

import (
	"image/color"
	"reflect"
	"unsafe"

	"karalis/internal/camera"
	"karalis/pkg/app"
	pub_object "karalis/pkg/object"

	raylib "github.com/gen2brain/raylib-go/raylib"
	lmath "karalis/pkg/lmath"
)

type Prim struct {
	mdl   raylib.Model
	pos   raylib.Vector3
	rot   raylib.Vector3
	scale raylib.Vector3
	color color.RGBA
}

func (p *Prim) init() error {
	p.pos = raylib.NewVector3(0, 0, 0)
	p.rot = raylib.NewVector3(0, 0, 0)
	p.scale = raylib.NewVector3(1, 1, 1)
	p.color = raylib.White
	p.mdl = raylib.Model{}

	return nil
}

func (p *Prim) GetModelMatrix() raylib.Matrix {
	matScale := raylib.MatrixScale(p.scale.X, p.scale.Y, p.scale.Z)
	Quat := lmath.Quat{}
	Quat = *Quat.FromEuler(float64(p.GetPitch()), float64(p.GetYaw()), float64(p.GetRoll()))
	matRotation := raylib.QuaternionToMatrix(raylib.NewQuaternion(float32(Quat.X), float32(Quat.Y), float32(Quat.Z), float32(Quat.W)))
	matTranslation := raylib.MatrixTranslate(p.pos.X, p.pos.Y, p.pos.Z)
	matTransform := raylib.MatrixMultiply(raylib.MatrixMultiply(matScale, matRotation), matTranslation)
	matTransform = raylib.MatrixMultiply(p.mdl.Transform, matTransform)
	return matTransform
}

func (p *Prim) GetColor() color.Color {
	return p.color
}

func (p *Prim) SetColor(col color.Color) {
	switch color := col.(type) {
	case color.RGBA:
		p.color = color
	}
}

func (p *Prim) GetScale() raylib.Vector3 {
	return p.scale
}

func (p *Prim) SetScale(sc raylib.Vector3) {
	p.scale = sc
}

func (p *Prim) GetPos() raylib.Vector3 {
	return p.pos
}

func (p *Prim) SetPos(pos raylib.Vector3) {
	p.pos = pos
}

func (p *Prim) GetPitch() float32 {
	return p.rot.X
}

func (p *Prim) SetPitch(pitch float32) {
	p.rot.X = pitch
}

func (p *Prim) GetYaw() float32 {
	return p.rot.Y
}

func (p *Prim) SetYaw(yaw float32) {
	p.rot.Y = yaw
}

func (p *Prim) GetRoll() float32 {
	return p.rot.Z
}

func (p *Prim) SetRoll(roll float32) {
	p.rot.Z = roll
}

func (p *Prim) GetVertices() []raylib.Vector3 {
	verts := []raylib.Vector3{}
	length := p.mdl.Meshes.VertexCount

	var mdlverts []float32

	header := (*reflect.SliceHeader)(unsafe.Pointer(&mdlverts))
	header.Data = uintptr(unsafe.Pointer(p.mdl.Meshes.Vertices))
	header.Len = int(length)
	header.Cap = int(length)

	for i := 0; i < len(mdlverts); i++ {
		verts = append(verts, raylib.NewVector3(mdlverts[3*i], mdlverts[3*i+1], mdlverts[3*i+2]))
	}
	return verts
}

func (p *Prim) GetUVs() []raylib.Vector2 {
	uvs := []raylib.Vector2{}
	length := p.mdl.Meshes.VertexCount
	var mdluvs []float32

	header := (*reflect.SliceHeader)(unsafe.Pointer(&mdluvs))
	header.Data = uintptr(unsafe.Pointer(p.mdl.Meshes.Texcoords))
	header.Len = int(length)
	header.Cap = int(length)

	for i := 0; i < len(mdluvs); i++ {
		uvs = append(uvs, raylib.NewVector2(mdluvs[2*i], mdluvs[2*i+1]))
	}
	return uvs
}

func (p *Prim) SetUVs(uvs []raylib.Vector2) {
	length := int(p.mdl.Meshes.VertexCount)
	var mdluvs []float32

	header := (*reflect.SliceHeader)(unsafe.Pointer(&mdluvs))
	header.Data = uintptr(unsafe.Pointer(p.mdl.Meshes.Texcoords))
	header.Len = length * 2
	header.Cap = length * 2

	for i := 0; i < len(uvs); i++ {
		mdluvs[i*2] = uvs[i].X
		mdluvs[i*2+1] = uvs[i].Y
	}
	pub_object.UpdateModelUVs(&p.mdl)
}

func (p *Prim) GetMaterials() *raylib.Material {
	return p.mdl.Materials
}

func (p *Prim) SetTexture(mat *raylib.Material, tex raylib.Texture2D) {
	raylib.SetMaterialTexture(mat, raylib.MapDiffuse, tex)
}

func (p *Prim) GetTexture(mat *raylib.Material) raylib.Texture2D {
	return mat.Maps.Texture
}

func (p *Prim) Prerender(cam *camera.Cam) []func() {
	return []func(){}
}

func (p *Prim) Render(cam *camera.Cam) []func() {
	matTransform := p.GetModelMatrix()
	sh := app.CurApp.GetShader()
	p.mdl.Materials.Shader = *sh.GetShader()
	raylib.DrawMesh(*p.mdl.Meshes, *p.mdl.Materials, matTransform)

	return []func(){}
}

func (p *Prim) Postrender(cam *camera.Cam) []func() {
	return []func(){}
}

func (p *Prim) Update(dt float32) {
}

func (p *Prim) OnAdd() {
}

func (p *Prim) OnRemove() {
}

func (p *Prim) AddChild(obj pub_object.Object) {
}

func (p *Prim) RemChild(obj pub_object.Object) {
}

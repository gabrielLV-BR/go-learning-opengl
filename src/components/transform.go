package components

import (
	. "github.com/go-gl/mathgl/mgl64"
)

type Transform struct {
	Position Vec3
	Rotation Quat
	Size     Vec3
}

func (t *Transform) Move(vel Vec3) {
	t.Position.Add(vel)
}

func (t *Transform) Scale(scalar float64) {
	t.Size.Mul(scalar)
}

func (t *Transform) Rotate(rotation Quat) {
	t.Rotation.Mul(rotation)
}

func (t *Transform) GetMatrix() Mat4 {
	var mat Mat4 = Ident4()

	mat.Mul4(Translate3D(t.Position[0], t.Position[1], t.Position[2]))
	mat.Mul4(Scale3D(t.Size[0], t.Size[1], t.Size[2]))
	mat.Mul4(HomogRotate3D(t.Rotation.Len(), t.Rotation.V))

	return mat
}

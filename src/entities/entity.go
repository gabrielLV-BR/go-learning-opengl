package entities

type Entity interface {
	Create()
	Update(float32)
	Render()
	Destroy()
}

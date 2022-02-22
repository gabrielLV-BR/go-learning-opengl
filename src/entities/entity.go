package entities

type Entity interface {
	Setup()
	Update(float32)
	Render()
	Destroy()
}

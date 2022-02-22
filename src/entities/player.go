package entities

import (
	"gl/learning/components"
	"gl/learning/input"

	"github.com/go-gl/glfw/v3.2/glfw"
)

type Player struct {
	components.Transform
	Speed float32
	//
	keys map[uint32]bool
}

func NewPlayer() Player {
	return Player {
		Transform: components.DefaultTransform(),
		Speed: 10.,
		keys: make(map[uint32]bool),
	}
}

func (p *Player) HandleInput(key glfw.Key, action glfw.Action, mods glfw.ModifierKey) {

}

func (p *Player) Setup() {
	input.Subscribe(
		func(k glfw.Key, a glfw.Action, mk glfw.ModifierKey) {
			p.HandleInput(k, a, mk)
		},
	)
}

func (p *Player) Update(float32) {}

func (p *Player) Render() {}

func (p *Player) Destroy() {}

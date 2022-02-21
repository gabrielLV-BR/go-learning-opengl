package input

import (
	. "github.com/go-gl/glfw/v3.2/glfw"
)

type InputListener func(Key, Action, ModifierKey)

var (
	listeners = []InputListener{}
)

func Subscribe(f InputListener) {
	listeners = append(listeners, f)
}

func HandleKeyInput(w *Window, key Key, scancode int, action Action, mods ModifierKey) {
	for _, v := range listeners {
		v(key, action, mods)
	}
}

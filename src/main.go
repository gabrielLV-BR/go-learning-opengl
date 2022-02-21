package main

import (
	"fmt"
	"gl/learning/input"
	"gl/learning/shaders"
	"gl/learning/utils"
	"log"
	"runtime"
	"time"

	"github.com/go-gl/gl/v4.1-core/gl" // OR: github.com/go-gl/gl/v2.1/gl
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/go-gl/mathgl/mgl32"
)

const (
	width  = 500
	height = 500
)

var (
	triangle = []float32{
		0, 0.5, 0,
		-0.5, -0.5, 0,
		0.5, -0.5, 0,
	}
	start_time int64 = 0.0
)

func init() {
	runtime.LockOSThread()
}

func main() {
	// Initialize GLFW and OpenGL
	window := initGlfw()
	program := initOpenGL()

	// Initialize necessary variables
	start_time = time.Now().UnixMilli()

	// Setup callbacks
	window.SetKeyCallback(input.HandleKeyInput)

	//
	vao := makeVao(triangle)

	model := mgl32.Ident4()
	camera := mgl32.LookAt(0., 0., 10., 0., 0., 0., 0., 1., 0.)
	proj := mgl32.Perspective(mgl32.DegToRad(60.), width/height, 0.01, 1000.)

	
	var x float32 = 0.

	input.Subscribe(func(key glfw.Key, action glfw.Action, _ glfw.ModifierKey) {
		if key == glfw.KeyD {
			fmt.Println("Master o f penes")
			model = mgl32.Translate3D(x, 0., 0.)
			x += 0.1
		}
	})

	// Main Loop
	for !window.ShouldClose() {
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

		gl.UseProgram(program.Id)

		program.SetUniform("time", float32(time.Now().UnixMilli()-start_time)/100.)
		program.SetUniform("u_model", model)
		program.SetUniform("u_view", camera)
		program.SetUniform("u_proj", proj)

		gl.BindVertexArray(vao)
		gl.DrawArrays(gl.TRIANGLES, 0, int32(len(triangle)/3))

		glfw.PollEvents()
		window.SwapBuffers()
	}

	// Cleanup
	glfw.Terminate()
}

// initGlfw initializes glfw and returns a Window to use.
func initGlfw() *glfw.Window {
	if err := glfw.Init(); err != nil {
		panic(err)
	}
	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	window, err := glfw.CreateWindow(width, height, "Conway's Game of Life", nil, nil)
	utils.Check(err)

	window.MakeContextCurrent()

	return window
}

// initOpenGL initializes OpenGL and returns an intiialized program.
func initOpenGL() shaders.Program {
	if err := gl.Init(); err != nil {
		panic(err)
	}
	version := gl.GoStr(gl.GetString(gl.VERSION))
	log.Println("OpenGL version", version)

	vertexShader := shaders.NewShader(
		"../assets/shaders/basic.vs",
		gl.VERTEX_SHADER,
	)

	fragmentShader := shaders.NewShader(
		"../assets/shaders/basic.fs",
		gl.FRAGMENT_SHADER,
	)

	return shaders.NewProgram(fragmentShader, vertexShader)
}

// makeVao initializes and returns a vertex array from the points provided.
func makeVao(points []float32) uint32 {
	var vbo uint32
	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, 4*len(points), gl.Ptr(points), gl.STATIC_DRAW)

	var vao uint32
	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)
	gl.EnableVertexAttribArray(0)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 0, nil)

	return vao
}

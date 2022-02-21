package shaders

import (
	"github.com/go-gl/gl/v4.1-core/gl" // OR: github.com/go-gl/gl/v2.1/gl
	"github.com/go-gl/mathgl/mgl32"
)

type Program struct {
	VertexShader   Shader
	FragmentShader Shader
	Id             uint32
	cache          map[string]int32
}

func NewProgram(fs, vs Shader) Program {
	prog := gl.CreateProgram()
	gl.AttachShader(prog, vs.Id)
	gl.AttachShader(prog, fs.Id)
	gl.LinkProgram(prog)

	return Program{
		VertexShader:   vs,
		FragmentShader: fs,
		Id:             prog,
		cache:          make(map[string]int32),
	}
}

func (p *Program) SetUniform(name string, value interface{}) {
	location := p.getCached(name)
	switch val := value.(type) {
	case float32:
		gl.Uniform1f(location, val)
	case float64:
		gl.Uniform1d(location, val)
	case int32:
		gl.Uniform1i(location, val)
	case mgl32.Mat4:
		gl.UniformMatrix4fv(location, 1, false, &val[0])
		// ...
	}
}

func (p *Program) getCached(name string) int32 {
	if val, ok := p.cache[name]; ok {
		return val
	} else {
		loc := gl.GetUniformLocation(p.Id, gl.Str(name+"\x00"))
		p.cache[name] = loc
		return loc
	}
}

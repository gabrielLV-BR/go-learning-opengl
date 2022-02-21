package shaders

import (
	"fmt"
	"gl/learning/utils"
	"os"
	"strings"

	"github.com/go-gl/gl/v4.1-core/gl"
)

type ShaderType uint32

const (
	VERTEX   ShaderType = gl.VERTEX_SHADER
	FRAGMENT ShaderType = gl.FRAGMENT_SHADER
)

type Shader struct {
	Id   uint32
	Kind ShaderType
}

func NewShader(path string, kind uint32) Shader {
	source := readShader(path)
	id, err := compileShader(source, kind)
	utils.Check(err)

	return Shader{Id: id, Kind: ShaderType(kind)}
}

func readShader(path string) string {
	data, err := os.ReadFile(path)
	utils.Check(err)

	return string(data)
}

func compileShader(source string, shaderType uint32) (uint32, error) {
	shader := gl.CreateShader(shaderType)

	csources, free := gl.Strs(source)
	gl.ShaderSource(shader, 1, csources, nil)
	free()
	gl.CompileShader(shader)

	var status int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(shader, logLength, nil, gl.Str(log))

		return 0, fmt.Errorf("failed to compile %v: %v", source, log)
	}

	return shader, nil
}

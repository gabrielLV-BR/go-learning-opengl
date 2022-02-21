#version 410

in vec3 vp;
uniform float time;

void main() {
    gl_Position = vec4(
        vec3(
            vp.x * sin(time),
            vp.yz
        ), 
        1.0
    );
}
#version 410

in vec3 vp;
uniform float time;

uniform mat4 u_model, u_view, u_proj;

void main() {
    // gl_Position = vec4(vp, 1.0);
    gl_Position = u_proj * u_view * u_model * vec4(vp, 1.0);
}
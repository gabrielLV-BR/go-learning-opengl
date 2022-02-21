#version 410

out vec4 frag_colour;

uniform float time;

float n_sin(float x) {
    return (sin(x) + 1.) * 0.5;
}

float n_cos(float x) {
    return (cos(x) + 1.) * 0.5;
}


void main() {
    frag_colour = vec4(n_sin(time), n_cos(time), 1., 1.0);
}
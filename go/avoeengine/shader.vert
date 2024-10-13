#version 330 core

in vec3 position;
in vec3 normal;

uniform mat4 modelViewMatrix;
uniform mat4 projectionMatrix;

void main() {
	gl_Position = projectionMatrix * modelViewMatrix * vec4(position, 1.0);
}

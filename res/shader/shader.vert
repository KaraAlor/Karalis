#version 330

// Input vertex attributes
in vec3 vertexPosition;
in vec2 vertexTexCoord;
in vec3 vertexNormal;
in vec4 vertexColor;

// Input uniform values
uniform mat4 mvp;
uniform mat4 matModel;
uniform mat4 matNormal;

// Output vertex attributes (to fragment shader)
out vec4 fragPosition;
out vec2 fragTexCoord;
out vec4 fragColor;
out vec3 fragNormal;
out vec4 fragScrPos;
out float fragClip;

// Custom Uniforms
uniform vec3 portalPos = vec3(0,0,0);
uniform vec3 portalNorm = vec3(0,0,0);

void main() {
    // Calculate final vertex position
    gl_Position = mvp*vec4(vertexPosition, 1.0);

    // Send vertex attributes to fragment shader
    fragPosition = matModel*vec4(vertexPosition, 1.0);
    fragTexCoord = vertexTexCoord;
    fragColor = vertexColor;
    fragNormal = normalize(vec3(matNormal*vec4(vertexNormal, 1.0)));
    fragScrPos = mvp*vec4(vertexPosition, 1.0);
    fragClip = dot((vec3(fragPosition) - portalPos), portalNorm);
}

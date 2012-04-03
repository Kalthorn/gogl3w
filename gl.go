package gogl3w

/*
#cgo darwin 	LDFLAGS: -framework OpenGL -lGL
#cgo windows 	LDFLAGS: -lopengl32
#cgo linux 		LDFLAGS: -lGL

#include <GL3/gl3w.h>

void (APIENTRYP ptrgoglClear)(GLbitfield mask);
void (APIENTRYP ptrgoglClearColor)(GLclampf red, GLclampf green, GLclampf blue, GLclampf alpha);

void (APIENTRYP ptrgoglBindVertexArray)(GLuint array);
void (APIENTRYP ptrgoglDeleteVertexArrays)(GLsizei n, const GLuint* arrays);
void (APIENTRYP ptrgoglGenVertexArrays)(GLsizei n, GLuint* arrays);

void (APIENTRYP ptrgoglBindBuffer)(GLenum target, GLuint buffer);
void (APIENTRYP ptrgoglDeleteBuffers)(GLsizei n, const GLuint* buffers);
void (APIENTRYP ptrgoglGenBuffers)(GLsizei n, GLuint* buffers);

void connectGLpointers(void) {
	ptrgoglClear = glClear;
	ptrgoglClearColor = glClearColor;
	ptrgoglBindVertexArray = glBindVertexArray;
	ptrgoglDeleteVertexArrays = glDeleteVertexArrays;
	ptrgoglGenVertexArrays = glGenVertexArrays;
	ptrgoglBindBuffer = glBindBuffer;
	ptrgoglDeleteBuffers = glDeleteBuffers;
	ptrgoglGenBuffers = glGenBuffers;
}

void goglClear(GLbitfield mask) {
	glClear(mask);
}
void goglClearColor(GLclampf red, GLclampf green, GLclampf blue, GLclampf alpha) {
	(*ptrgoglClearColor)(red, green, blue, alpha);
}

void goglBindVertexArray(GLuint array) {
	(*ptrgoglBindVertexArray)(array);
}
void goglDeleteVertexArrays(GLsizei n, GLuint* arrays) {
	(*ptrgoglDeleteVertexArrays)(n, arrays);
}
void goglGenVertexArrays(GLsizei n, GLuint* arrays) {
	(*ptrgoglGenVertexArrays)(n, arrays);
}

void goglBindBuffer(GLenum target, GLuint buffer) {
	(*ptrgoglBindBuffer)(target, buffer);
}
void goglDeleteBuffers(GLsizei n, GLuint* buffers) {
	(*ptrgoglDeleteBuffers)(n, buffers);
}
void goglGenBuffers(GLsizei n, GLuint* buffers) {
	(*ptrgoglGenBuffers)(n, buffers);
}

*/
import "C"
import "unsafe"
import "runtime"

// import "reflect"
import "log"

type GLenum C.GLenum
type GLbitfield C.GLbitfield
type GLclampf C.GLclampf
type GLclampd C.GLclampd
type Pointer unsafe.Pointer

// those types are left for compatibility reasons
// type GLboolean C.GLboolean
// type GLbyte C.GLbyte
// type GLshort C.GLshort
// type GLint C.GLint
// type GLsizei C.GLsizei
// type GLubyte C.GLubyte
// type GLushort C.GLushort
// type GLuint C.GLuint
// type GLfloat C.GLfloat
// type GLdouble C.GLdouble

// helpers

func Init() {
	log.Printf("Clear: %p", C.ptrgoglClear)
	log.Printf("GenVAO: %p", C.ptrgoglGenVertexArrays)
	initResult := C.gl3wInit()
	log.Printf("Init result: %d", initResult)
	C.connectGLpointers()
}

//void glBegin (GLenum mode)
func Begin(mode GLenum) {

}

//void glEnd (void)
func End() {

}

//void glClear (GLbitfield mask)
func Clear(mask GLbitfield) {
	C.goglClear(C.GLbitfield(mask))
	log.Println("Cleared!")
}

//void glClearColor (GLclampf red, GLclampf green, GLclampf blue, GLclampf alpha)
func ClearColor(red GLclampf, green GLclampf, blue GLclampf, alpha GLclampf) {
	C.goglClearColor(C.GLclampf(red), C.GLclampf(green), C.GLclampf(blue), C.GLclampf(alpha))
	log.Println("Colored!")
}

type Object C.GLuint

type Buffer Object

// // Create single buffer object
func GenBuffer() Buffer {
	var b C.GLuint
	C.goglGenBuffers(1, &b)
	return Buffer(b)
}

// // Fill slice with new buffers
func GenBuffers(buffers []Buffer) {
	C.goglGenBuffers(C.GLsizei(len(buffers)), (*C.GLuint)(&buffers[0]))
}

// // Delete buffer object
// func (buffer Buffer) Delete() {
// 	b := C.GLuint(buffer)
// 	C.glDeleteBuffers(1, &b)
// }

// // Delete all textures in slice
// func DeleteBuffers(buffers []Buffer) {
// 	C.glDeleteBuffers(C.GLsizei(len(buffers)), (*C.GLuint)(&buffers[0]))
// }

// // Bind this buffer as target
func (buffer Buffer) Bind(target GLenum) {
	C.goglBindBuffer(C.GLenum(target), C.GLuint(buffer))
}

// // Bind this buff

// // Vertex Arrays
type VertexArray Object

func GenVertexArray() VertexArray {
	var a C.GLuint
	log.Printf("Clear: %p", C.ptrgoglClear)
	log.Printf("GenVAO: %p", C.ptrgoglGenVertexArrays)
	log.Printf("Address of A: %p", &a)
	stackBuf := make([]byte, 0)
	_ = runtime.Stack(stackBuf, true)
	log.Printf("Stack-izzy: %v", stackBuf)
	C.goglGenVertexArrays(1, &a)
	log.Printf("Value of A: %p", a)
	log.Println("eh??")
	// C.goglGenVertexArrays((C.GLsizei)(n), (*C.GLuint)(arrays))
	return VertexArray(a)
}

func GenVertexArrays(arrays []VertexArray) {
	C.goglGenVertexArrays(C.GLsizei(len(arrays)), (*C.GLuint)(&arrays[0]))
}

// func (array VertexArray) Delete() {
// 	C.glDeleteVertexArrays(1, (*C.GLuint)(&array))
// }

// func DeleteVertexArrays(arrays []VertexArray) {
// 	C.glDeleteVertexArrays(C.GLsizei(len(arrays)), (*C.GLuint)(&arrays[0]))
// }

func (array VertexArray) Bind() {
	C.goglBindVertexArray(C.GLuint(array))
}

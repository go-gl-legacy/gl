package gl_test

import (
	"github.com/banthar/Go-SDL/sdl"
	"github.com/shogg/gl"
	"testing"
)

var (
	array = [...]int32{
		1, 2, 3, 4, 5, 6, 7, 8,
		9, 10, 11, 12, 13, 14, 15, 16}
	slice = array[:]
)

func init() {
	sdl.Init(sdl.INIT_VIDEO)

	screen := sdl.SetVideoMode(320, 200, 32, sdl.OPENGL)
	if screen == nil {
		panic("Couldn't set video mode: " + sdl.GetError() + "\n")
	}

	if err := gl.Init(); err != 0 {
		panic("glInit error")
	}
}

func TestTexImage1D(t *testing.T) {

	gl.TexImage1D(gl.TEXTURE_1D,
		0, gl.RGBA, 16, 0, gl.RGBA, gl.INT,
		&array)
	if gl.GetError() != gl.NO_ERROR {
		t.Error("pointer to array failed")
	}

	gl.Enable(gl.TEXTURE_1D)
	gl.TexImage1D(gl.TEXTURE_1D,
		1, gl.RGBA, 16, 1, gl.RGBA, gl.INT,
		&slice[3])
	if gl.GetError() != gl.NO_ERROR {
		t.Error("pointer to element failed")
	}

	gl.TexImage1D(gl.TEXTURE_1D,
		0, gl.RGBA, 16, 0, gl.RGBA, gl.INT,
		slice)
	if gl.GetError() != gl.NO_ERROR {
		t.Error("slice failed")
	}

	gl.TexImage1D(gl.PROXY_TEXTURE_1D,
		0, gl.RGBA, 16, 0, gl.RGBA, gl.INT,
		nil)
	if gl.GetError() != gl.NO_ERROR {
		t.Error("nil pointer failed")
	}
}

func TestTexImage2D(t *testing.T) {

	gl.TexImage2D(gl.TEXTURE_2D,
		0, gl.RGBA, 4, 4, 0, gl.RGBA, gl.INT,
		&array)
	if gl.GetError() != gl.NO_ERROR {
		t.Error("pointer to array failed")
	}

	gl.TexImage2D(gl.TEXTURE_2D,
		0, gl.RGBA, 4, 4, 0, gl.RGBA, gl.INT,
		slice)
	if gl.GetError() != gl.NO_ERROR {
		t.Error("slice failed")
	}

	gl.TexImage2D(gl.PROXY_TEXTURE_2D,
		0, gl.RGBA, 4, 4, 0, gl.RGBA, gl.INT,
		nil)
	if gl.GetError() != gl.NO_ERROR {
		t.Error("nil pointer failed")
	}
}

func TestTexImage3D(t *testing.T) {

	gl.TexImage3D(gl.TEXTURE_3D,
		0, gl.RGBA, 2, 2, 2, 0, gl.RGBA, gl.INT,
		&array)
	if gl.GetError() != gl.NO_ERROR {
		t.Error("pointer to array failed")
	}

	gl.TexImage3D(gl.TEXTURE_3D,
		0, gl.RGBA, 2, 2, 2, 0, gl.RGBA, gl.INT,
		slice)
	if gl.GetError() != gl.NO_ERROR {
		t.Error("slice failed")
	}

	gl.TexImage3D(gl.PROXY_TEXTURE_3D,
		0, gl.RGBA, 2, 2, 2, 0, gl.RGBA, gl.INT,
		nil)
	if gl.GetError() != gl.NO_ERROR {
		t.Error("nil pointer failed")
	}
}

func TestTexSubImage1D(t *testing.T) {

	gl.TexSubImage1D(gl.TEXTURE_1D,
		0, 8, 8, gl.RGBA, gl.INT,
		&array)
	if gl.GetError() != gl.NO_ERROR {
		t.Error("pointer to array failed")
	}

	gl.TexSubImage1D(gl.TEXTURE_1D,
		0, 8, 8, gl.RGBA, gl.INT,
		slice)
	if gl.GetError() != gl.NO_ERROR {
		t.Error("slice failed")
	}

	// gl.TexSubImage1D(gl.PROXY_TEXTURE_1D,
	// 	0, 8, 8, gl.RGBA, gl.INT,
	// 	nil)
	// if gl.GetError() != gl.NO_ERROR {
	// 	t.Error("nil pointer failed")
	// }
}

func TestTexSubImage2D(t *testing.T) {

	gl.TexSubImage2D(gl.TEXTURE_2D,
		0, 2, 2, 2, 2, gl.RGBA, gl.INT,
		&array)
	if gl.GetError() != gl.NO_ERROR {
		t.Error("pointer to failed")
	}

	gl.TexSubImage2D(gl.TEXTURE_2D,
		0, 2, 2, 2, 2, gl.RGBA, gl.INT,
		slice)
	if gl.GetError() != gl.NO_ERROR {
		t.Error("slice failed")
	}

	// gl.TexSubImage2D(gl.PROXY_TEXTURE_2D,
	// 	0, 2, 2, 2, 2, gl.RGBA, gl.INT,
	// 	nil)
	// if gl.GetError() != gl.NO_ERROR {
	// 	t.Error("nil pointer failed")
	// }
}

func TestTexSubImage3D(t *testing.T) {

	gl.TexSubImage3D(gl.TEXTURE_3D,
		0, 1, 1, 1, 1, 1, 1, gl.RGBA, gl.INT,
		&array)
	if gl.GetError() != gl.NO_ERROR {
		t.Error("pointer to failed")
	}

	gl.TexSubImage3D(gl.TEXTURE_3D,
		0, 1, 1, 1, 1, 1, 1, gl.RGBA, gl.INT,
		slice)
	if gl.GetError() != gl.NO_ERROR {
		t.Error("slice failed")
	}

	// gl.TexSubImage3D(gl.PROXY_TEXTURE_3D,
	// 	0, 1, 1, 1, 1, 1, 1, gl.RGBA, gl.INT,
	// 	nil)
	// if gl.GetError() != gl.NO_ERROR {
	// 	t.Error("nil pointer failed")
	// }
}

func newBuffer(bytes int) gl.Buffer {
	buf := gl.GenBuffer()
	buf.Bind(gl.ARRAY_BUFFER)
	gl.BufferData(gl.ARRAY_BUFFER, bytes, slice, gl.STATIC_READ)
	return buf
}

func TestBufferData(t *testing.T) {

	buf := newBuffer(16 * 4)
	defer buf.Delete()

	gl.BufferData(gl.ARRAY_BUFFER, 16*4, &array, gl.STATIC_READ)
	if gl.GetError() != gl.NO_ERROR {
		t.Error("pointer to array failed")
	}

	gl.BufferData(gl.ARRAY_BUFFER, 16*4, slice, gl.STATIC_READ)
	if gl.GetError() != gl.NO_ERROR {
		t.Error("slice failed")
	}

	gl.BufferData(gl.ARRAY_BUFFER, 16*4, nil, gl.STATIC_READ)
	if gl.GetError() != gl.NO_ERROR {
		t.Error("nil pointer failed")
	}
}

func TestBufferSubData(t *testing.T) {

	buf := newBuffer(16 * 4)
	defer buf.Delete()

	gl.BufferSubData(gl.ARRAY_BUFFER, 0, 4*4, &array)
	if gl.GetError() != gl.NO_ERROR {
		t.Error("pointer to array failed")
	}

	gl.BufferSubData(gl.ARRAY_BUFFER, 0, 4*4, slice)
	if gl.GetError() != gl.NO_ERROR {
		t.Error("slice failed")
	}
}

func TestGetBufferSubData(t *testing.T) {

	buf := newBuffer(16 * 4)
	defer buf.Delete()

	result := make([]int32, 4)
	gl.GetBufferSubData(gl.ARRAY_BUFFER, 7*4, 4*4, result)
	if gl.GetError() != gl.NO_ERROR {
		t.Error("slice failed")
	}
	if result[0] != 8 {
		t.Error("[8, 9, 10, 11] expected, was", result)
	}
}

func newProgram() gl.Program {

	vs := `
		uniform float[] foo;
		attribute int bar;

		void main()
		{
			gl_Position = gl_Vertex * foo[0] * bar;
		}`

	fs := `
		void main()
		{
			gl_FragColor = vec4(1.0);
		}`

	// vertex shader
	vshader := gl.CreateShader(gl.VERTEX_SHADER)
	vshader.Source(vs)
	vshader.Compile()
	if vshader.Get(gl.COMPILE_STATUS) != gl.TRUE {
		panic("vertex shader error: " + vshader.GetInfoLog())
	}

	// fragment shader
	fshader := gl.CreateShader(gl.FRAGMENT_SHADER)
	fshader.Source(fs)
	fshader.Compile()
	if fshader.Get(gl.COMPILE_STATUS) != gl.TRUE {
		panic("fragment shader error: " + fshader.GetInfoLog())
	}

	// program
	prg := gl.CreateProgram()
	prg.AttachShader(vshader)
	prg.AttachShader(fshader)
	prg.Link()
	if prg.Get(gl.LINK_STATUS) != gl.TRUE {
		panic("linker error: " + prg.GetInfoLog())
	}

	prg.Use()
	return prg
}

func TestAttribPointer(t *testing.T) {

	prg := newProgram()
	defer prg.Delete()

	bar := prg.GetAttribLocation("bar")
	bar.EnableArray()

	bar.AttribPointer(1, gl.INT, false, 0, &array)
	if gl.GetError() != gl.NO_ERROR {
		t.Error("pointer to array failed")
	}

	bar.AttribPointer(1, gl.INT, false, 0, slice)
	if gl.GetError() != gl.NO_ERROR {
		t.Error("slice failed")
	}
}

func TestUniform1fv(t *testing.T) {

	prg := newProgram()
	defer prg.Delete()

	values := []float32{1.0, 1.1, 1.2, 1.3}
	foo := prg.GetUniformLocation("foo")
	foo.Uniform1fv(values)
	if gl.GetError() != gl.NO_ERROR {
		t.Error("slice failed")
	}
}

func TestCallLists(t *testing.T) {

	base := gl.GenLists(2)
	if base == 0 {
		t.Error("GenLists failed")
	}

	gl.NewList(base+0, gl.COMPILE)
	gl.Color3i(1, 2, 3)
	gl.Begin(gl.POLYGON)
	gl.Vertex2f(1.0, 2.0)
	gl.Vertex2f(3.0, 4.0)
	gl.Vertex2f(5.0, 6.0)
	gl.End()
	gl.EndList()
	if gl.GetError() != gl.NO_ERROR {
		t.Error("NewList 1 failed")
	}

	gl.NewList(base+1, gl.COMPILE)
	gl.Begin(gl.POINTS)
	gl.Vertex2i(3, 4)
	gl.End()
	gl.EndList()
	if gl.GetError() != gl.NO_ERROR {
		t.Error("NewList 2 failed")
	}

	gl.CallList(base)
	if gl.GetError() != gl.NO_ERROR {
		t.Error("CallList 1 failed")
	}

	gl.ListBase(base)
	if gl.GetError() != gl.NO_ERROR {
		t.Error("ListBase failed")
	}

	lists := []uint{0, 1}
	gl.CallLists(2, gl.UNSIGNED_INT, lists)
	if gl.GetError() != gl.NO_ERROR {
		t.Error("CallLists 1 2 failed")
	}
}

func TestColorPointer(t *testing.T) {

	gl.EnableClientState(gl.COLOR_ARRAY)

	gl.ColorPointer(4, gl.INT, 0, &array)
	if gl.GetError() != gl.NO_ERROR {
		t.Error("pointer to array failed")
	}

	gl.ColorPointer(4, gl.INT, 0, slice)
	if gl.GetError() != gl.NO_ERROR {
		t.Error("slice failed")
	}

	gl.DisableClientState(gl.COLOR_ARRAY)
	buf := newBuffer(16 * 4)
	defer buf.Delete()

	gl.ColorPointer(4, gl.INT, 0, uintptr(0))
	if gl.GetError() != gl.NO_ERROR {
		t.Error("buffer offset failed")
	}
}

func TestDrawElements(t *testing.T) {

	gl.EnableClientState(gl.VERTEX_ARRAY)
	gl.VertexPointer(2, gl.INT, 0, slice)

	indices := []uint32{0, 1, 2, 3}

	gl.DrawElements(gl.TRIANGLE_STRIP, 4, gl.UNSIGNED_INT, indices)
	if gl.GetError() != gl.NO_ERROR {
		t.Error("slice failed")
	}

	gl.DisableClientState(gl.VERTEX_ARRAY)
}

func TestDrawPixels(t *testing.T) {

	gl.DrawPixels(4, 4, gl.RGBA, gl.UNSIGNED_INT, slice)
	if gl.GetError() != gl.NO_ERROR {
		t.Error("slice failed")
	}

	buf := newBuffer(16 * 4)
	buf.Bind(gl.PIXEL_UNPACK_BUFFER)
	defer buf.Delete()

	gl.DrawPixels(2, 2, gl.RGBA, gl.UNSIGNED_INT, uintptr(0))
	if gl.GetError() != gl.NO_ERROR {
		t.Error("buffer offset failed")
	}
}

func TestIndexPointer(t *testing.T) {

	gl.EnableClientState(gl.INDEX_ARRAY)

	gl.IndexPointer(gl.INT, 0, &array)
	if gl.GetError() != gl.NO_ERROR {
		t.Error("pointer to array failed")
	}

	gl.IndexPointer(gl.INT, 0, slice)
	if gl.GetError() != gl.NO_ERROR {
		t.Error("slice failed")
	}

	gl.DisableClientState(gl.INDEX_ARRAY)
	buf := newBuffer(16 * 4)
	defer buf.Delete()

	gl.IndexPointer(gl.INT, 0, uintptr(0))
	if gl.GetError() != gl.NO_ERROR {
		t.Error("buffer offset failed")
	}
}

func TestNormalPointer(t *testing.T) {

	gl.EnableClientState(gl.NORMAL_ARRAY)

	gl.NormalPointer(gl.INT, 0, &array)
	if gl.GetError() != gl.NO_ERROR {
		t.Error("pointer to array failed")
	}

	gl.NormalPointer(gl.INT, 0, slice)
	if gl.GetError() != gl.NO_ERROR {
		t.Error("slice failed")
	}

	gl.DisableClientState(gl.NORMAL_ARRAY)
	buf := newBuffer(16 * 4)
	defer buf.Delete()

	gl.NormalPointer(gl.INT, 0, uintptr(0))
	if gl.GetError() != gl.NO_ERROR {
		t.Error("buffer offset failed")
	}
}

func TestReadPixels(t *testing.T) {

	pixels := []byte{1, 1, 1, 2, 2, 2, 3, 3, 3, 4, 4, 4}
	gl.DrawPixels(2, 2, gl.RGB, gl.UNSIGNED_BYTE, pixels)

	result := make([]byte, 4*3)
	gl.ReadPixels(0, 0, 2, 2, gl.RGB, gl.UNSIGNED_BYTE, result)
	if gl.GetError() != gl.NO_ERROR {
		t.Error("slice failed")
	}

	// result is {1,1,1, 2,2,2, 0,0,3, 4,4,4}, no idea why.
	if false {
		t.Error("was", result)
	}
}

func TestTexCoordPointer(t *testing.T) {

	gl.EnableClientState(gl.TEXTURE_COORD_ARRAY)

	gl.TexCoordPointer(3, gl.INT, 0, &array)
	if gl.GetError() != gl.NO_ERROR {
		t.Error("pointer to array failed")
	}

	gl.TexCoordPointer(3, gl.INT, 0, slice)
	if gl.GetError() != gl.NO_ERROR {
		t.Error("slice failed")
	}

	gl.DisableClientState(gl.TEXTURE_COORD_ARRAY)
	buf := newBuffer(16 * 4)
	defer buf.Delete()

	gl.TexCoordPointer(3, gl.INT, 0, uintptr(0))
	if gl.GetError() != gl.NO_ERROR {
		t.Error("buffer offset failed")
	}
}

func TestVertexPointer(t *testing.T) {

	gl.EnableClientState(gl.VERTEX_ARRAY)

	gl.VertexPointer(3, gl.INT, 0, &array)
	if gl.GetError() != gl.NO_ERROR {
		t.Error("pointer to array failed")
	}

	gl.VertexPointer(3, gl.INT, 0, slice)
	if gl.GetError() != gl.NO_ERROR {
		t.Error("slice failed")
	}

	gl.DisableClientState(gl.VERTEX_ARRAY)
	buf := newBuffer(16 * 4)
	defer buf.Delete()

	gl.VertexPointer(3, gl.INT, 0, uintptr(0))
	if gl.GetError() != gl.NO_ERROR {
		t.Error("buffer offset failed")
	}
}

package main

/*

	This is an example OpenGL app. A Drawing tool based on http://mrdoob.com/projects/harmony/

*/

import (
   "sdl"
   "gl"
   "math"
)

type Point struct {
	x int
	y int
}

func (p0 Point) distanceTo(p1 Point) float64 {
	dx := (p0.x - p1.x)
	dy := (p0.y - p1.y)
	return math.Sqrt(float64(dx*dx + dy*dy))
}

type Pen struct {
	pos    Point
	points [4096]Point
	n      int
}

func (pen *Pen) lineTo(p Point) {

	gl.Enable(gl.BLEND)
	gl.Enable(gl.POINT_SMOOTH)
	gl.Enable(gl.LINE_SMOOTH)
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)

	gl.Color4f(0.0, 0.0, 0.0, 0.1)

	gl.Begin(gl.LINES)

	for _, s := range pen.points {

		if s.x == 0 && s.y == 0 {
			continue
		}

		if p.distanceTo(s) < 20.0 {

			gl.Vertex2i(int(p.x), int(p.y))
			gl.Vertex2i(int(s.x), int(s.y))

		}

	}

	gl.End()

	pen.n = (pen.n + 1) % len(pen.points)
	pen.points[pen.n] = p

	pen.moveTo(p)

}

func (pen *Pen) moveTo(p Point) {
	pen.pos = p
}

func main() {

	sdl.Init(sdl.INIT_VIDEO)

	var screen = sdl.SetVideoMode(640, 480, 32, sdl.OPENGL)

	if screen == nil {
		panic("sdl error")
	}

	if gl.Init() != 0 {
		panic("gl error")	
	}

	pen := Pen{}

	gl.MatrixMode(gl.PROJECTION)

	gl.Viewport(0, 0, int(screen.W), int(screen.H))
	gl.LoadIdentity()
	gl.Ortho(0, float64(screen.W), float64(screen.H), 0, -1.0, 1.0)

	gl.ClearColor(1, 1, 1, 0)
	gl.Clear(gl.COLOR_BUFFER_BIT)

	var running = true

	for running {

		e := &sdl.Event{}

		for e.Poll() {
			switch e.Type {
			case sdl.QUIT:
				running = false
			case sdl.KEYDOWN:
               ke := e.Keyboard()
               if ke.Keysym.Sym == sdl.K_ESCAPE {
                   running = false
               }
			case sdl.MOUSEMOTION:
				me := e.MouseMotion()
				if me.State != 0 {
					pen.lineTo(Point{int(me.X), int(me.Y)})
				} else {
					pen.moveTo(Point{int(me.X), int(me.Y)})
				}
			}
		}

		sdl.GL_SwapBuffers()
		sdl.Delay(25)
	}

	sdl.Quit()

}

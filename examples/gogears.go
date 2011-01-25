package main

/*
 * 3-D gear wheels.  This program is in the public domain.
 *
 * Command line options:
 *    -info      print GL implementation information
 *
 *
 * Brian Paul
 */

/* this is go version based on SDL version */

import "sdl"
import "gl/gl20"
import "math"
import "flag"

var printInfo = flag.Bool("info", false, "print GL implementation information")

var T0 uint32 = 0
var Frames uint32 = 0

/**

  Draw a gear wheel.  You'll probably want to call this function when
  building a display list since we do a lot of trig here.

  Input:  inner_radius - radius of hole at center
          outer_radius - radius at center of teeth
          width - width of gear
          teeth - number of teeth
          tooth_depth - depth of tooth

 **/

func gear(inner_radius, outer_radius, width float64, teeth int, tooth_depth float64) {
	var i int
	var r0, r1, r2 float64
	var angle, da float64
	var u, v, len float64

	r0 = inner_radius
	r1 = outer_radius - tooth_depth/2.0
	r2 = outer_radius + tooth_depth/2.0

	da = 2.0 * math.Pi / float64(teeth) / 4.0

	gl.ShadeModel(gl.FLAT)

	gl.Normal3d(0.0, 0.0, 1.0)

	/* draw front face */
	gl.Begin(gl.QUAD_STRIP)
	for i = 0; i <= teeth; i++ {
		angle = float64(i) * 2.0 * math.Pi / float64(teeth)
		gl.Vertex3d(r0*math.Cos(angle), r0*math.Sin(angle), width*0.5)
		gl.Vertex3d(r1*math.Cos(angle), r1*math.Sin(angle), width*0.5)
		if i < teeth {
			gl.Vertex3d(r0*math.Cos(angle), r0*math.Sin(angle), width*0.5)
			gl.Vertex3d(r1*math.Cos(angle+3*da), r1*math.Sin(angle+3*da), width*0.5)
		}
	}
	gl.End()

	/* draw front sides of teeth */
	gl.Begin(gl.QUADS)
	da = 2.0 * math.Pi / float64(teeth) / 4.0
	for i = 0; i < teeth; i++ {
		angle = float64(i) * 2.0 * math.Pi / float64(teeth)

		gl.Vertex3d(r1*math.Cos(angle), r1*math.Sin(angle), width*0.5)
		gl.Vertex3d(r2*math.Cos(angle+da), r2*math.Sin(angle+da), width*0.5)
		gl.Vertex3d(r2*math.Cos(angle+2*da), r2*math.Sin(angle+2*da), width*0.5)
		gl.Vertex3d(r1*math.Cos(angle+3*da), r1*math.Sin(angle+3*da), width*0.5)
	}
	gl.End()

	gl.Normal3d(0.0, 0.0, -1.0)

	/* draw back face */
	gl.Begin(gl.QUAD_STRIP)
	for i = 0; i <= teeth; i++ {
		angle = float64(i) * 2.0 * math.Pi / float64(teeth)
		gl.Vertex3d(r1*math.Cos(angle), r1*math.Sin(angle), -width*0.5)
		gl.Vertex3d(r0*math.Cos(angle), r0*math.Sin(angle), -width*0.5)
		if i < teeth {
			gl.Vertex3d(r1*math.Cos(angle+3*da), r1*math.Sin(angle+3*da), -width*0.5)
			gl.Vertex3d(r0*math.Cos(angle), r0*math.Sin(angle), -width*0.5)
		}
	}
	gl.End()

	/* draw back sides of teeth */
	gl.Begin(gl.QUADS)
	da = 2.0 * math.Pi / float64(teeth) / 4.0
	for i = 0; i < teeth; i++ {
		angle = float64(i) * 2.0 * math.Pi / float64(teeth)

		gl.Vertex3d(r1*math.Cos(angle+3*da), r1*math.Sin(angle+3*da), -width*0.5)
		gl.Vertex3d(r2*math.Cos(angle+2*da), r2*math.Sin(angle+2*da), -width*0.5)
		gl.Vertex3d(r2*math.Cos(angle+da), r2*math.Sin(angle+da), -width*0.5)
		gl.Vertex3d(r1*math.Cos(angle), r1*math.Sin(angle), -width*0.5)
	}
	gl.End()

	/* draw outward faces of teeth */
	gl.Begin(gl.QUAD_STRIP)
	for i = 0; i < teeth; i++ {
		angle = float64(i) * 2.0 * math.Pi / float64(teeth)

		gl.Vertex3d(r1*math.Cos(angle), r1*math.Sin(angle), width*0.5)
		gl.Vertex3d(r1*math.Cos(angle), r1*math.Sin(angle), -width*0.5)
		u = r2*math.Cos(angle+da) - r1*math.Cos(angle)
		v = r2*math.Sin(angle+da) - r1*math.Sin(angle)
		len = math.Sqrt(u*u + v*v)
		u /= len
		v /= len
		gl.Normal3d(v, -u, 0.0)
		gl.Vertex3d(r2*math.Cos(angle+da), r2*math.Sin(angle+da), width*0.5)
		gl.Vertex3d(r2*math.Cos(angle+da), r2*math.Sin(angle+da), -width*0.5)
		gl.Normal3d(math.Cos(angle), math.Sin(angle), 0.0)
		gl.Vertex3d(r2*math.Cos(angle+2*da), r2*math.Sin(angle+2*da), width*0.5)
		gl.Vertex3d(r2*math.Cos(angle+2*da), r2*math.Sin(angle+2*da), -width*0.5)
		u = r1*math.Cos(angle+3*da) - r2*math.Cos(angle+2*da)
		v = r1*math.Sin(angle+3*da) - r2*math.Sin(angle+2*da)
		gl.Normal3d(v, -u, 0.0)
		gl.Vertex3d(r1*math.Cos(angle+3*da), r1*math.Sin(angle+3*da), width*0.5)
		gl.Vertex3d(r1*math.Cos(angle+3*da), r1*math.Sin(angle+3*da), -width*0.5)
		gl.Normal3d(math.Cos(angle), math.Sin(angle), 0.0)
	}

	gl.Vertex3d(r1*math.Cos(0), r1*math.Sin(0), width*0.5)
	gl.Vertex3d(r1*math.Cos(0), r1*math.Sin(0), -width*0.5)

	gl.End()

	gl.ShadeModel(gl.SMOOTH)

	/* draw inside radius cylinder */
	gl.Begin(gl.QUAD_STRIP)
	for i = 0; i <= teeth; i++ {
		angle = float64(i) * 2.0 * math.Pi / float64(teeth)
		gl.Normal3d(-math.Cos(angle), -math.Sin(angle), 0.0)
		gl.Vertex3d(r0*math.Cos(angle), r0*math.Sin(angle), -width*0.5)
		gl.Vertex3d(r0*math.Cos(angle), r0*math.Sin(angle), width*0.5)
	}
	gl.End()

}

var view_rotx float64 = 20.0
var view_roty float64 = 30.0
var view_rotz float64 = 0.0
var gear1, gear2, gear3 uint
var angle float64 = 0.0

func draw() {

	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

	gl.PushMatrix()
	gl.Rotated(view_rotx, 1.0, 0.0, 0.0)
	gl.Rotated(view_roty, 0.0, 1.0, 0.0)
	gl.Rotated(view_rotz, 0.0, 0.0, 1.0)

	gl.PushMatrix()
	gl.Translated(-3.0, -2.0, 0.0)
	gl.Rotated(angle, 0.0, 0.0, 1.0)
	gl.CallList(gear1)
	gl.PopMatrix()

	gl.PushMatrix()
	gl.Translated(3.1, -2.0, 0.0)
	gl.Rotated(-2.0*angle-9.0, 0.0, 0.0, 1.0)
	gl.CallList(gear2)
	gl.PopMatrix()

	gl.PushMatrix()
	gl.Translated(-3.1, 4.2, 0.0)
	gl.Rotated(-2.0*angle-25.0, 0.0, 0.0, 1.0)
	gl.CallList(gear3)
	gl.PopMatrix()

	gl.PopMatrix()

	sdl.GL_SwapBuffers()

	Frames++
	{
		t := sdl.GetTicks()
		if t-T0 >= 5000 {
			seconds := (t - T0) / 1000.0
			fps := Frames / seconds
			print(Frames, " frames in ", seconds, " seconds = ", fps, " FPS\n")
			T0 = t
			Frames = 0
		}
	}
}


func idle() {
	angle += 2.0
}

/* new window size or exposure */
func reshape(width int, height int) {

	h := float64(height) / float64(width)

	gl.Viewport(0, 0, width, height)
	gl.MatrixMode(gl.PROJECTION)
	gl.LoadIdentity()
	gl.Frustum(-1.0, 1.0, -h, h, 5.0, 60.0)
	gl.MatrixMode(gl.MODELVIEW)
	gl.LoadIdentity()
	gl.Translatef(0.0, 0.0, -40.0)
}

func init_() {
	pos := []float32{5.0, 5.0, 10.0, 0.0}
	red := []float32{0.8, 0.1, 0.0, 1.0}
	green := []float32{0.0, 0.8, 0.2, 1.0}
	blue := []float32{0.2, 0.2, 1.0, 1.0}

	gl.Lightfv(gl.LIGHT0, gl.POSITION, pos)
	gl.Enable(gl.CULL_FACE)
	gl.Enable(gl.LIGHTING)
	gl.Enable(gl.LIGHT0)
	gl.Enable(gl.DEPTH_TEST)

	/* make the gears */
	gear1 = gl.GenLists(1)
	gl.NewList(gear1, gl.COMPILE)
	gl.Materialfv(gl.FRONT, gl.AMBIENT_AND_DIFFUSE, red)
	gear(1.0, 4.0, 1.0, 20, 0.7)
	gl.EndList()

	gear2 = gl.GenLists(1)
	gl.NewList(gear2, gl.COMPILE)
	gl.Materialfv(gl.FRONT, gl.AMBIENT_AND_DIFFUSE, green)
	gear(0.5, 2.0, 2.0, 10, 0.7)
	gl.EndList()

	gear3 = gl.GenLists(1)
	gl.NewList(gear3, gl.COMPILE)
	gl.Materialfv(gl.FRONT, gl.AMBIENT_AND_DIFFUSE, blue)
	gear(1.3, 2.0, 0.5, 10, 0.7)
	gl.EndList()

	gl.Enable(gl.NORMALIZE)

	if *printInfo {
		print("GL_RENDERER   = ", gl.GetString(gl.RENDERER), "\n")
		print("GL_VERSION    = ", gl.GetString(gl.VERSION), "\n")
		print("GL_VENDOR     = ", gl.GetString(gl.VENDOR), "\n")
		print("GL_EXTENSIONS = ", gl.GetString(gl.EXTENSIONS), "\n")
	}

}

func main() {

	flag.Parse()

	var done bool
	var keys []uint8

	sdl.Init(sdl.INIT_VIDEO)

	var screen = sdl.SetVideoMode(300, 300, 16, sdl.OPENGL|sdl.RESIZABLE)

	if screen == nil {
		sdl.Quit()
		panic("Couldn't set 300x300 GL video mode: " + sdl.GetError() + "\n")
	}

	sdl.WM_SetCaption("Gears", "gears")

	init_()
	reshape(int(screen.W), int(screen.H))
	done = false
	for !done {
		var event sdl.Event

		idle()
		for event.Poll() {
			switch event.Type {
			case sdl.VIDEORESIZE:
				screen = sdl.SetVideoMode(int(event.Resize().W), int(event.Resize().H), 16,
					sdl.OPENGL|sdl.RESIZABLE)
				if screen != nil {
					reshape(int(screen.W), int(screen.H))
				} else {
					panic("we couldn't set the new video mode??")
				}
				break

			case sdl.QUIT:
				done = true
				break
			}
		}
		keys = sdl.GetKeyState()

		if keys[sdl.K_ESCAPE] != 0 {
			done = true
		}
		if keys[sdl.K_UP] != 0 {
			view_rotx += 5.0
		}
		if keys[sdl.K_DOWN] != 0 {
			view_rotx -= 5.0
		}
		if keys[sdl.K_LEFT] != 0 {
			view_roty += 5.0
		}
		if keys[sdl.K_RIGHT] != 0 {
			view_roty -= 5.0
		}
		if keys[sdl.K_z] != 0 {
			if (sdl.GetModState() & sdl.KMOD_RSHIFT) != 0 {
				view_rotz -= 5.0
			} else {
				view_rotz += 5.0
			}
		}

		draw()
	}
	sdl.Quit()
	return

}

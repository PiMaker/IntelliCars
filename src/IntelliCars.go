package main

import (
	"runtime"
    "time"
    
    "./pkg"

	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.1/glfw"
	"github.com/llgcode/draw2d"
	"github.com/llgcode/draw2d/draw2dgl"
)

var (
	width, height int
	font draw2d.FontData
)

func reshape(window *glfw.Window, w, h int) {
	gl.ClearColor(1, 1, 1, 1)
	gl.Viewport(0, 0, int32(w), int32(h))
	gl.MatrixMode(gl.PROJECTION)
	gl.LoadIdentity()
	gl.Ortho(0, float64(w), 0, float64(h), -1, 1)
	gl.Scalef(1, -1, 1)
	gl.Translatef(0, float32(-h), 0)
	gl.Enable(gl.BLEND)
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)
	gl.Disable(gl.DEPTH_TEST)
	width, height = w, h
    intellicars.Reshape(float64(width), float64(height))
}

func display() {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

	gl.LineWidth(1)
	gc := draw2dgl.NewGraphicContext(width, height)
	gc.SetFontData(draw2d.FontData{
		Name:   "luxi",
		Family: draw2d.FontFamilyMono,
		Style:  draw2d.FontStyleBold | draw2d.FontStyleItalic})

	intellicars.Draw(*gc)

	gl.Flush()
}

func init() {
	runtime.LockOSThread()
}

func main() {
	err := glfw.Init()
	if err != nil {
		panic(err)
	}
	defer glfw.Terminate()
	width, height = 1280, 720
	window, err := glfw.CreateWindow(width, height, "IntelliCars", nil, nil)
	if err != nil {
		panic(err)
	}

	window.MakeContextCurrent()
	window.SetSizeCallback(reshape)
	window.SetKeyCallback(onKey)
	window.SetCharCallback(onChar)

	glfw.SwapInterval(1)

	err = gl.Init()
	if err != nil {
		panic(err)
	}

	reshape(window, width, height)
    
    intellicars.Init(float64(width), float64(height))
    
    go updateLoop(window)
    drawLoop(window)
}

func drawLoop(window *glfw.Window) {
    for !window.ShouldClose() {
        display()
        window.SwapBuffers()
        glfw.PollEvents()
	}
}

func updateLoop(window *glfw.Window) {
    ticker := time.NewTicker(time.Second / 60)
	for {
		intellicars.Update()
		<-ticker.C
	}
}

func onChar(w *glfw.Window, char rune) {
}

func onKey(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
	switch {
        case key == glfw.KeyEscape && action == glfw.Press,
            key == glfw.KeyQ && action == glfw.Press:
            w.SetShouldClose(true)
	}
}
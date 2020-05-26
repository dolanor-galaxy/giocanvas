// confetti -- random shapes
package main

import (
	"flag"
	"image/color"
	"math/rand"

	"gioui.org/app"
	"gioui.org/io/system"
	"gioui.org/unit"
	"github.com/ajstarks/giocanvas"
)

func rn(n int) float32 {
	return float32(rand.Intn(n))
}

func rn8(n int) uint8 {
	return uint8(rand.Intn(n))
}

func main() {
	var w, h, nshapes, maxsize int
	flag.IntVar(&w, "width", 1200, "canvas width")
	flag.IntVar(&h, "height", 900, "canvas height")
	flag.IntVar(&nshapes, "n", 500, "number of shapes")
	flag.IntVar(&maxsize, "size", 10, "max size")
	flag.Parse()
	width := float32(w)
	height := float32(h)
	size := app.Size(unit.Dp(width), unit.Dp(height))
	title := app.Title("Confetti")

	go func() {
		w := app.NewWindow(title, size)
		canvas := giocanvas.NewCanvas(width, height)

		for e := range w.Events() {
			if e, ok := e.(system.FrameEvent); ok {
				canvas.Context.Reset(e.Queue, e.Config, e.Size)
				canvas.CenterRect(50, 50, 100, 100, color.RGBA{0, 0, 0, 255})
				for i := 0; i < nshapes; i++ {
					color := color.RGBA{rn8(255), rn8(255), rn8(255), rn8(255)}
					canvas.CenterRect(rn(100), rn(100), rn(maxsize), rn(maxsize), color)
				}
				e.Frame(canvas.Context.Ops)
			}
		}
	}()
	app.Main()
}

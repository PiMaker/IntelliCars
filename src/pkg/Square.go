package intellicars

import (
	"github.com/llgcode/draw2d/draw2dgl"
    "github.com/llgcode/draw2d/draw2dkit"
	"image/color"
)

/*
Square defines a drawable square.
*/
type Square struct {
    X, Y, Width, Height, Rotation float64
}

func (square Square) icdraw(gc draw2dgl.GraphicContext) {
    gc.Save()
    gc.Rotate(square.Rotation)
    gc.SetFillColor(color.RGBA{0x33, 0x33, 0x33, 0xff})
    gc.SetStrokeColor(color.RGBA{0xaa, 0xaa, 0xaa, 0xff})
    gc.SetLineWidth(5)
    draw2dkit.Rectangle(gc, square.X, square.Y, square.X + square.Width, square.Y + square.Height)
    gc.FillStroke()
    gc.Restore()
}
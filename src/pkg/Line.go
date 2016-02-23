package intellicars

import (
	"github.com/llgcode/draw2d/draw2dgl"
	"image/color"
)

/*
Line defines a drawable square.
*/
type Line struct {
    X1, Y1, X2, Y2 float64
}

func (line Line) icdraw(gc draw2dgl.GraphicContext) {
    gc.Save()
    gc.SetStrokeColor(color.RGBA{0xaa, 0xaa, 0xaa, 0xff})
    gc.SetLineWidth(4)
    gc.BeginPath()
    gc.MoveTo(line.X1, line.X2)
    gc.LineTo(line.X2, line.Y2)
    gc.Stroke()
    gc.Restore()
}
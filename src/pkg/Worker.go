package intellicars

import (
	"github.com/llgcode/draw2d/draw2dgl"
    "math"
)

var (
    currentScroll float64
    width float64
)

func Init(screenWidth float64) {
    currentScroll = screenWidth
    width = screenWidth
    InitTerrain()
}

func Update() {
    UpdateTerrain(currentScroll)
    currentScroll += 1
}

func Draw(gc draw2dgl.GraphicContext) {
    gc.Save()
    gc.Translate(math.Min(0, width - currentScroll), 0)
    DrawTerrain(gc)
    DrawShapes(gc)
    gc.Restore()
}
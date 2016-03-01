package intellicars

import (
    "github.com/llgcode/draw2d/draw2dgl"
    "github.com/aquilax/go-perlin"
	"image/color"
	"time"
)

var (
    lines []Line
    
    perlinGenerator *perlin.Perlin
    
    segment float64 = 30
    minHeight float64 = 500
    randomMultiplier float64 = 250
    perlinDivider float64 = 500
)

func UpdateTerrain(rightmost float64) {
    for lines[len(lines) - 1].X2 < rightmost {
        last := lines[len(lines) - 1]
        line := Line {
            X1: last.X2,
            X2: last.X2 + segment,
            Y1: last.Y2,
            Y2: randomY(rightmost + last.X2)}
        lines = append(lines, line)
        RegisterTerrainLine(line)
    }
}

func InitTerrain() {
    lines = make([]Line, 1)
    lines[0] = Line {
        X1: -segment*2,
        X2: -segment,
        Y1: -minHeight,
        Y2: minHeight }
    RegisterTerrainLine(lines[0])
    
    perlinGenerator = perlin.NewPerlin(2., 2., 3, time.Now().UnixNano())
}

func DrawTerrain(gc draw2dgl.GraphicContext) {
    gc.Save()
    defer gc.Restore()
    
    gc.SetStrokeColor(color.RGBA{0xaa, 0xaa, 0xaa, 0xff})
    gc.SetLineWidth(4)
    gc.BeginPath()
    
    gc.MoveTo(lines[0].X1, lines[0].Y1)

    for _, line := range lines {
        gc.LineTo(line.X2, line.Y2)
    }
    
    gc.Stroke()
}

func randomY(offset float64) float64 {
    return minHeight - perlinGenerator.Noise1D(offset/perlinDivider) * randomMultiplier
}
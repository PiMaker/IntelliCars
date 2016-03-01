package intellicars

import (
	"github.com/llgcode/draw2d/draw2dgl"
	"../../chipmunk"
	"image/color"
	"github.com/llgcode/draw2d/draw2dkit"
    "math"
)

var (
    debugRender = false
)

func DrawShapes(gc draw2dgl.GraphicContext) {
    for _, car := range logicCars {
        if car.framesBehind == -1 {
            drawPhysicsShapes(gc, car.shapes, 0)
        } else {
            drawPhysicsShapes(gc, car.shapes, uint8(math.Max(float64(car.framesBehind/2), 0xaa)))
        }
    }
}

func drawPhysicsShapes(gc draw2dgl.GraphicContext, shapes []*chipmunk.Shape, highlight uint8) {
    gc.Save()
    gc.SetFillColor(color.RGBA{highlight, 0xaa, 0xaa, 0xff})
    gc.SetStrokeColor(color.RGBA{0x66, 0x66, 0x66, 0xff})
    gc.SetLineWidth(2)
    for _, shape := range shapes {
        gc.Save()
        if shape.ShapeType() == chipmunk.ShapeType_Box {
            t := shape.GetAsBox()
            gc.Translate(float64(shape.Body.Position().X), float64(shape.Body.Position().Y))
            gc.Rotate(float64(t.Shape.Body.Angle()))
            gc.Translate(-float64(t.Width)/2, -float64(t.Height)/2)
            draw2dkit.Rectangle(gc,
                float64(0),
                float64(0),
                float64(t.Width),
                float64(t.Height))
            gc.FillStroke()
        } else if shape.ShapeType() == chipmunk.ShapeType_Circle {
            t := shape.GetAsCircle()
            gc.Translate(float64(shape.Body.Position().X), float64(shape.Body.Position().Y))
            gc.Rotate(float64(t.Shape.Body.Angle()))
            draw2dkit.Circle(gc, float64(0), float64(0), float64(t.Radius))
            gc.FillStroke()
            gc.MoveTo(float64(0), float64(0))
            gc.LineTo(float64(0), float64(t.Radius))
            gc.Stroke()
        } else if shape.ShapeType() == chipmunk.ShapeType_Polygon {
            t := shape.GetAsPolygon()
            gc.Translate(float64(shape.Body.Position().X), float64(shape.Body.Position().Y))
            gc.Rotate(float64(t.Shape.Body.Angle()))
            gc.MoveTo(float64(t.Verts[0].X), float64(t.Verts[0].Y))
            for i, vert := range t.Verts {
                if i > 0 {
                    gc.LineTo(float64(vert.X), float64(vert.Y))
                }
            }
            gc.LineTo(float64(t.Verts[0].X), float64(t.Verts[0].Y))
            gc.FillStroke()
        }
        gc.Restore()
        
        if debugRender {
            gc.Save()
            gc.SetStrokeColor(color.RGBA{0xff, 0x00, 0x00, 0xff})
            draw2dkit.Rectangle(gc,
                float64(shape.AABB().Lower.X),
                float64(shape.AABB().Lower.Y),
                float64(shape.AABB().Upper.X),
                float64(shape.AABB().Upper.Y))
            gc.Stroke()
            gc.Restore()
        }
    }
    gc.Restore()
}
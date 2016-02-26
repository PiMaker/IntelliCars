package intellicars

import (
	"github.com/llgcode/draw2d/draw2dgl"
	"github.com/vova616/chipmunk"
	"image/color"
	"github.com/llgcode/draw2d/draw2dkit"
)

/*
Shape defines a drawable shape.
*/
type Shape interface {
    icdraw(gc draw2dgl.GraphicContext)
}

var (
    debugRender = false
    
    shapes []Shape
)

func AddShape(shape Shape)  {
    if shapes == nil {
        shapes = make([]Shape, 0)
    }
    
    shapes = append(shapes, shape)
}

func ClearShapes() {
    shapes = make([]Shape, 0)
}

func GetShapes() []Shape {
    return shapes
}

func DrawShapes(gc draw2dgl.GraphicContext) {
    for _, shape := range shapes {
        shape.icdraw(gc)
    }
}

func DrawPhysicsShapes(gc draw2dgl.GraphicContext, shapes []*chipmunk.Shape) {
    gc.Save()
    gc.SetFillColor(color.RGBA{0x33, 0x33, 0x33, 0xff})
    gc.SetStrokeColor(color.RGBA{0xaa, 0xaa, 0xaa, 0xff})
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
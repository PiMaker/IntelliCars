package intellicars

import (
	"github.com/vova616/chipmunk"
	"github.com/vova616/chipmunk/vect"
    "github.com/llgcode/draw2d/draw2dgl"
	"image/color"
	"github.com/llgcode/draw2d/draw2dkit"
)

type Car struct {
    shapes []*chipmunk.Shape
}

func GenerateCar() Car {
    car := Car {}
    car.shapes = make([]*chipmunk.Shape, 1)
    car.shapes[0] = chipmunk.NewBox(vect.Vect{vect.Float(0), vect.Float(0)}, vect.Float(30), vect.Float(30))
    car.shapes[0].SetElasticity(0.9)
    RegisterPhysicsCar(car)
    AddShape(car)
    return car
}

func (car Car) GetPhysicsShapes() []*chipmunk.Shape {
    return car.shapes
}

func (car Car) icdraw(gc draw2dgl.GraphicContext) {
    gc.Save()
    gc.SetFillColor(color.RGBA{0x33, 0x33, 0x33, 0xff})
    gc.SetStrokeColor(color.RGBA{0xaa, 0xaa, 0xaa, 0xff})
    gc.SetLineWidth(5)
    for _, shape := range car.shapes {
        if shape.ShapeType() == chipmunk.ShapeType_Box {
            t := shape.GetAsBox()
            gc.Translate(float64(shape.Body.Position().X + t.Position.X), float64(shape.Body.Position().Y + t.Position.Y))
            gc.Translate(-float64(t.Width)/2, -float64(t.Height)/2)
            gc.Rotate(float64(t.Shape.Body.Angle()))
            draw2dkit.Rectangle(gc,
                float64(0),
                float64(0),
                float64(t.Width),
                float64(t.Height))
            gc.FillStroke()
        }
    }
    gc.Restore()
}
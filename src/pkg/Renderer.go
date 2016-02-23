package intellicars

import (
	"github.com/llgcode/draw2d/draw2dgl"
)

/*
Shape defines a drawable shape.
*/
type Shape interface {
    icdraw(gc draw2dgl.GraphicContext)
}

var shapes []Shape

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
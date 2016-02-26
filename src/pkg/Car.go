package intellicars

import (
	"github.com/vova616/chipmunk"
    "github.com/vova616/chipmunk/vect"
    "github.com/llgcode/draw2d/draw2dgl"
)

type Car struct {
    shape *chipmunk.Shape
    wheels []Wheel
    shapes []*chipmunk.Shape
}

func GenerateRandomCar() Car {
    car := Car {}
    
    verts := []vect.Vect {
        vect.Vect{vect.Float(0), 0},
        vect.Vect{vect.Float(0), 60},
        vect.Vect{vect.Float(200), 30}}
    car.shape = chipmunk.NewPolygon(verts, vect.Vect{-100,-30})
    
    car.wheels = make([]Wheel, 2)
    
    car.wheels[0] = Wheel{}
    car.wheels[0].center = vect.Vect{vect.Float(200),30}
    car.wheels[0].shape = chipmunk.NewCircle(vect.Vect{0,0}, 30)
    car.wheels[0].shape.SetElasticity(0.2)
    car.wheels[0].shape.SetFriction(1.0)
    
    car.wheels[1] = Wheel{}
    car.wheels[1].center = vect.Vect{vect.Float(0),60}
    car.wheels[1].shape = chipmunk.NewCircle(vect.Vect{0,0}, 22)
    car.wheels[1].shape.SetElasticity(0.2)
    car.wheels[1].shape.SetFriction(1.0)
    
    car.shapes = make([]*chipmunk.Shape, len(car.wheels) + 1)
    car.shapes[0] = car.shape
    for i, wheel := range car.wheels {
        car.shapes[i + 1] = wheel.shape
    }
        
    RegisterPhysicsCar(car)
    AddShape(car)
    return car
}

func (car Car) GetPhysicsWheels() []Wheel {
    return car.wheels
}

func (car Car) GetPhysicsShape() *chipmunk.Shape {
    return car.shape
}

func (car Car) icdraw(gc draw2dgl.GraphicContext) {
    DrawPhysicsShapes(gc, car.shapes)
}
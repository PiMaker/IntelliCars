package intellicars

import (
	"../../chipmunk"
	"../../chipmunk/vect"
	"fmt"
)

var (
    cars []PhysicsCar
    space *chipmunk.Space
)

type PhysicsCar interface {
    GetPhysicsShape() *chipmunk.Shape
    GetPhysicsWheels() []Wheel
}

func InitPhysics() {
    space = chipmunk.NewSpace()
    space.Gravity = vect.Vect{0, 500}
    cars = make([]PhysicsCar, 0)
}

func RegisterPhysicsCar(car PhysicsCar) {
    cars = append(cars, car)
    
    polyshape := car.GetPhysicsShape()
    polyshape.Group = 1;
    
    fmt.Print(polyshape.Moment(16))
    fmt.Print(" => ")
    fmt.Println(chipmunk.PolygonMomentOptimized(polyshape.GetAsPolygon(), 16))
    
    polybody := chipmunk.NewBody(vect.Float(16), polyshape.Moment(float32(16)))
    polybody.AddShape(polyshape)
    space.AddBody(polybody)
    
    for _, wheel := range car.GetPhysicsWheels() {
        shape := wheel.shape
        shape.Group = 1;
        body := chipmunk.NewBody(vect.Float(2), vect.Float(shape.Moment(float32(2))))
        body.SetPosition(wheel.center)
        body.AddShape(shape)
        space.AddBody(body)
        body.SetPosition(vect.Vect{vect.Float(500), vect.Float(10)})
        joint := chipmunk.NewPivotJointAnchor(polybody, body, vect.Vect{wheel.center.X, wheel.center.Y}, vect.Vector_Zero)
        space.AddConstraint(joint)
        motor := chipmunk.NewSimpleMotor(polybody, body, -wheel.shape.GetAsCircle().Radius/10)
        space.AddConstraint(motor)
    }
    
    polybody.SetPosition(vect.Vect{vect.Float(500), vect.Float(10)})
}

func UpdatePhysics() {
    space.Step(vect.Float(1.0 / 60.0))
}

func RegisterTerrainLine(line Line) {
    terrain := chipmunk.NewBodyStatic()
    shape := chipmunk.NewSegment(vect.Vect{vect.Float(line.X1), vect.Float(line.Y1)}, vect.Vect{vect.Float(line.X2), vect.Float(line.Y2)}, 0)
    terrain.AddShape(shape)
    space.AddBody(terrain)
}
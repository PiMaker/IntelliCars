package intellicars

import (
	"github.com/vova616/chipmunk"
	"github.com/vova616/chipmunk/vect"
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
    space.Gravity = vect.Vect{0, 300}
    cars = make([]PhysicsCar, 0)
}

func RegisterPhysicsCar(car PhysicsCar) {
    cars = append(cars, car)
    
    polyshape := car.GetPhysicsShape()
    polyshape.Group = 1;
    polybody := chipmunk.NewBody(vect.Float(10), vect.Float(25000))
    polybody.SetPosition(vect.Vect{vect.Float(500), vect.Float(10)})
    polybody.AddShape(polyshape)
    space.AddBody(polybody)
    
    for _, wheel := range car.GetPhysicsWheels() {
        shape := wheel.shape
        shape.Group = 1;
        body := chipmunk.NewBody(vect.Float(1), vect.Float(shape.Moment(float32(1))))
        body.SetPosition(wheel.center)
        body.AddShape(shape)
        body.UserData = true
        space.AddBody(body)
        //middle := shape.GetAsCircle().Radius
        joint := chipmunk.NewPivotJointAnchor(polybody, body, vect.Vect{body.Position().X/4, body.Position().Y/4}, vect.Vect{0,0})
        space.AddConstraint(joint)
    }
}

func UpdatePhysics() {
    for _, body := range space.Bodies {
        rotate, found := body.UserData.(bool)
        if found && rotate {
            body.AddTorque(0)
        }
    }
    
    space.Step(vect.Float(1.0 / 600.0))
}

func RegisterTerrainLine(line Line) {
    terrain := chipmunk.NewBodyStatic()
    shape := chipmunk.NewSegment(vect.Vect{vect.Float(line.X1), vect.Float(line.Y1)}, vect.Vect{vect.Float(line.X2), vect.Float(line.Y2)}, 0)
    terrain.AddShape(shape)
    space.AddBody(terrain)
}
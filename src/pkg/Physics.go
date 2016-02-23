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
    GetPhysicsShapes() []*chipmunk.Shape
}

func InitPhysics() {
    space = chipmunk.NewSpace()
    space.Gravity = vect.Vect{0, 200}
    cars = make([]PhysicsCar, 0)
}

func RegisterPhysicsCar(car PhysicsCar) {
    cars = append(cars, car)
    
    body := chipmunk.NewBody(vect.Float(1), vect.Float(car.GetPhysicsShapes()[0].Moment(float32(1))))
    body.SetPosition(vect.Vect{vect.Float(500), vect.Float(100)})
    
    for _, shape := range car.GetPhysicsShapes() {
        body.AddShape(shape)
    }
    
    space.AddBody(body)
}

func UpdatePhysics() {
    space.Step(vect.Float(1.0 / 600.0))
}

func RegisterTerrainLine(line Line) {
    terrain := chipmunk.NewBodyStatic()
    shape := chipmunk.NewSegment(vect.Vect{vect.Float(line.X1), vect.Float(line.Y1)}, vect.Vect{vect.Float(line.X2), vect.Float(line.Y2)}, 0)
    terrain.AddShape(shape)
    space.AddBody(terrain)
}
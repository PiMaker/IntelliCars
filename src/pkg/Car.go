package intellicars

import (
	"../../chipmunk"
    "../../chipmunk/vect"
    "math/rand"
    "../../convexhull"
    "math"
)

type Car struct {
    shape *chipmunk.Shape
    wheels []Wheel
    shapes []*chipmunk.Shape
    
    maxDistance float64
    framesBehind int
}

func GenerateFromParent(parent Car) *Car {
    parent.framesBehind = 0
    parent.maxDistance = 0
    
    RegisterPhysicsCar(parent)
    return &parent
}

func GenerateRandomCar() *Car {
    car := Car {}
    var verts chipmunk.Vertices
    
    valid := false
    for !valid {
        verts = make(chipmunk.Vertices, 0)
        points := convexhull.PointList{}
        vcount := rand.Intn(10) + 10
        
        for index := 0; index < vcount; index++ {
            points = append(points,
                convexhull.MakePoint(rand.Float64()*200, rand.Float64()*200))
        }
        
        hull, worked := points.Compute()
        
        if !worked {
            continue
        }
        
        for _, point := range hull {
            verts = append(verts, vect.Vect{vect.Float(point.X), vect.Float(point.Y)})
        }
        
        valid = verts.ValidatePolygon()
    }
    
    car.shape = chipmunk.NewPolygon(verts, vect.Vector_Zero)
    car.shape.SetFriction(0.5)
    car.shape.SetElasticity(0.05)
    
    for _, vert := range verts {
        if rand.Intn(3) == 0 {
            wheel := Wheel{}
            wheel.center = vert
            
            wheel.shape = chipmunk.NewCircle(vect.Vector_Zero, float32(rand.Intn(70) + 10))
            wheel.shape.SetElasticity(0.3)
            wheel.shape.SetFriction(0.99)
            
            valid = true
            
            for _, w := range car.wheels {
                tmp := (math.Pow(float64(wheel.center.X-w.center.X), 2)+math.Pow(float64(wheel.center.Y-w.center.Y), 2))
                valid = !(math.Pow(float64(wheel.shape.GetAsCircle().Radius-w.shape.GetAsCircle().Radius), 2) <= tmp &&
                     tmp <= math.Pow(float64(wheel.shape.GetAsCircle().Radius+w.shape.GetAsCircle().Radius), 2))
                if !valid {
                    break
                }
            }
            
            if valid {
                car.wheels = append(car.wheels, wheel)
            }
        }
    }
    
    car.shapes = make([]*chipmunk.Shape, len(car.wheels) + 1)
    car.shapes[0] = car.shape
    for i, wheel := range car.wheels {
        car.shapes[i + 1] = wheel.shape
    }
    
    car.framesBehind = 0
    car.maxDistance = 0
    
    RegisterPhysicsCar(car)
    return &car
}

func (car Car) GetPhysicsWheels() []Wheel {
    return car.wheels
}

func (car Car) GetPhysicsShape() *chipmunk.Shape {
    return car.shape
}
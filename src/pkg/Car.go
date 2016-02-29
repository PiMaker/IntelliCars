package intellicars

import (
	"../../chipmunk"
    "../../chipmunk/vect"
    "github.com/llgcode/draw2d/draw2dgl"
    "math/rand"
    "../../convexhull"
    "math"
)

type Car struct {
    shape *chipmunk.Shape
    wheels []Wheel
    shapes []*chipmunk.Shape
}

func GenerateTestCar() Car {
    car := Car {}
    
    verts := []vect.Vect {
        vect.Vect{vect.Float(0), 0},
        vect.Vect{vect.Float(0), 60},
        vect.Vect{vect.Float(200), 30}}
    car.shape = chipmunk.NewPolygon(verts, vect.Vector_Zero)
    car.shape.SetFriction(0)
    car.shape.SetElasticity(0.1)
    
    car.wheels = make([]Wheel, 2)
    
    car.wheels[0] = Wheel{}
    car.wheels[0].center = vect.Vect{vect.Float(0),30}
    car.wheels[0].shape = chipmunk.NewCircle(vect.Vect{0,0}, 60)
    car.wheels[0].shape.SetElasticity(0.3)
    car.wheels[0].shape.SetFriction(0.99)
    
    car.wheels[1] = Wheel{}
    car.wheels[1].center = vect.Vect{vect.Float(200),30}
    car.wheels[1].shape = chipmunk.NewCircle(vect.Vect{0,0}, 50)
    car.wheels[1].shape.SetElasticity(0.3)
    car.wheels[1].shape.SetFriction(0.99)
    
    car.shapes = make([]*chipmunk.Shape, len(car.wheels) + 1)
    car.shapes[0] = car.shape
    for i, wheel := range car.wheels {
        car.shapes[i + 1] = wheel.shape
    }
        
    RegisterPhysicsCar(car)
    AddShape(car)
    return car
}

func GenerateRandomCar() Car {
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
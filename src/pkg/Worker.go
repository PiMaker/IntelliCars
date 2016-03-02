package intellicars

import (
	"github.com/llgcode/draw2d/draw2dgl"
    "math"
	"math/rand"
	"time"
    "sort"
    //"strconv"
)

var (
    currentScroll float64
    width float64
    heightTranslation float64
    height float64
    
    logicCars CarList
    
    generation = 0
)

func Init(screenWidth, screenHeight float64) {
    currentScroll = screenWidth
    width = screenWidth
    
    height = screenHeight
    heightTranslation = 0
    
    InitPhysics()
    InitTerrain()
    
    rand.Seed(time.Now().UnixNano())
    
    logicCars = make(CarList, 20)
    
    for i := 0; i < 20; i ++ {
        logicCars[i] = GenerateRandomCar()
    }
}

func Reshape(screenWidth, screenHeight float64) {
    currentScroll = currentScroll - width + screenWidth
    width = screenWidth
    heightTranslation += screenHeight - height
    height = screenHeight
}

func Update() {
    UpdatePhysics()
    
    if (len(space.Bodies) > 0) {
        max := space.Bodies[0]
        for _, car := range logicCars {
            if (car.framesBehind != -1 && float64(car.GetPhysicsShape().Body.Position().X) > float64(max.Position().X)) {
                max = car.GetPhysicsShape().Body
            }
        }
        
        currentScroll = float64(max.Position().X) + (width / 3)
    }
    
    UpdateTerrain(currentScroll + (width / 3))
    
    // Logic
    done := true
    for i, car := range logicCars {
        if car.framesBehind != -1 {
            done = false
            current := float64(car.GetPhysicsShape().Body.Position().X)
            if current > car.maxDistance {
                logicCars[i].maxDistance = current
                logicCars[i].framesBehind = 0
            } else if current <= car.maxDistance {
                logicCars[i].framesBehind++
                if logicCars[i].framesBehind > 500 {
                    logicCars[i].framesBehind = -1
                }
            }
        }
    }
    
    if done {
        //newRound()
    }
}

func Draw(gc draw2dgl.GraphicContext) {
    gc.Save()
    gc.Translate(math.Min(0, width - currentScroll), heightTranslation)
    DrawTerrain(gc)
    DrawShapes(gc)
    gc.Restore()
    
    gc.Save()
    //s := "Generation: " + strconv.Itoa(generation)
    gc.Restore()
}

func newRound() {
    // Reset simulation
    InitPhysics()
    currentScroll = width
    
    // Genetic algorithm
    sort.Sort(logicCars) // Sort by fitness
    newLogicCars := make(CarList, 20)
    newLogicCars[0] = CopyCar(logicCars[0]) // Reuse the first two ones
    newLogicCars[1] = CopyCar(logicCars[1])
    for i := 2; i < 10; i++ { // Derive next 8 from first
        newLogicCars[i] = GenerateFromParent(*logicCars[0])
    }
    for i := 10; i < 12; i++ { // Derive two from second
        newLogicCars[i] = GenerateFromParent(*logicCars[1])
    }
    for i := 12; i < 20; i++ { // Randomly add remaining cars
        newLogicCars[i] = GenerateRandomCar()
    }
    logicCars = newLogicCars
    
    generation++
}
package intellicars

import (
	"github.com/llgcode/draw2d/draw2dgl"
    "math"
	"math/rand"
	"time"
)

var (
    currentScroll float64
    width float64
)

func Init(screenWidth float64) {
    currentScroll = screenWidth
    width = screenWidth
    InitPhysics()
    InitTerrain()
    
    rand.Seed(time.Now().UnixNano())
    
    GenerateRandomCar()
    GenerateRandomCar()
    GenerateRandomCar()
    GenerateRandomCar()
    GenerateRandomCar()
    GenerateRandomCar()
    GenerateRandomCar()
    GenerateRandomCar()
    GenerateRandomCar()
    GenerateRandomCar()
}

func Update() {
    UpdatePhysics()
    
    if (len(space.Bodies) > 0) {
        max := space.Bodies[0]
        for _, body := range space.Bodies {
            if (float64(body.Position().X) > float64(max.Position().X)) {
                max = body
            }
        }
        
        currentScroll = float64(max.Position().X) + (width / 2)
    }
    
    UpdateTerrain(currentScroll + (width / 2))
}

func Draw(gc draw2dgl.GraphicContext) {
    gc.Save()
    gc.Translate(math.Min(0, width - currentScroll), 0)
    DrawTerrain(gc)
    DrawShapes(gc)
    gc.Restore()
}
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
    heightTranslation float64
    height float64
)

func Init(screenWidth, screenHeight float64) {
    currentScroll = screenWidth
    width = screenWidth
    
    height = screenHeight
    heightTranslation = 0
    
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
    gc.Translate(math.Min(0, width - currentScroll), heightTranslation)
    DrawTerrain(gc)
    DrawShapes(gc)
    gc.Restore()
}
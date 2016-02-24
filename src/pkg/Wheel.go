package intellicars

import (
    "github.com/vova616/chipmunk/vect"
    "github.com/vova616/chipmunk"
)

type Wheel struct {
    center vect.Vect
    shape *chipmunk.Shape
}
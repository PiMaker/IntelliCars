package intellicars

import (
    "../../chipmunk/vect"
    "../../chipmunk"
)

type Wheel struct {
    center vect.Vect
    shape *chipmunk.Shape
}
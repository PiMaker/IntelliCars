package intellicars

import (
)

type CarList []*Car

func (c CarList) Len() int {
    return len(c)
}
    
func (c CarList) Less(i, j int) bool {
    return c[i].maxDistance > c[j].maxDistance
}

func (c CarList) Swap(i, j int) {
    t := c[i]
    c[i] = c[j]
    c[j] = t
}
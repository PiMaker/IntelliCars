package chipmunk

import ( 
	"./vect"
    "math"
)

// Wrapper around []vect.Vect.
type Vertices []vect.Vect

// Checks if verts forms a valid polygon.
// The vertices must be convex, winded clockwise and not intersect with itself.
func (verts Vertices) ValidatePolygon() bool {
	numVerts := len(verts)
	for i := 0; i < numVerts; i++ {
		a := verts[i]
		b := verts[(i+1)%numVerts]
		c := verts[(i+2)%numVerts]
        
        for i2 := 0; i2 < numVerts; i2++ {
            a2 := verts[i2]
		    b2 := verts[(i2+1)%numVerts]
            if (a2 == a && b2 == b) ||
                a2 == b || a == b2 {
                continue
            }
            if doIntersect(point{float64(a.X), float64(a.Y)},
                    point{float64(b.X), float64(b.Y)},
                    point{float64(a2.X), float64(a2.Y)},
                    point{float64(b2.X), float64(b2.Y)}) {
                return false
            }
        }
        
		if vect.Cross(vect.Sub(b, a), vect.Sub(c, b)) > 0.0 {
			return false
		}
	}

	return true
}

type point struct {
    x, y float64
}
 
// Given three colinear points p, q, r, the function checks if
// point q lies on line segment 'pr'
func onSegment(p, q, r point) bool {
    return q.x <= math.Max(p.x, r.x) && q.x >= math.Min(p.x, r.x) &&
        q.y <= math.Max(p.y, r.y) && q.y >= math.Min(p.y, r.y)
}
 
// To find orientation of ordered triplet (p, q, r).
// The function returns following values
// 0 --> p, q and r are colinear
// 1 --> Clockwise
// 2 --> Counterclockwise
func orientation(p, q, r point) int {
    // See http://www.geeksforgeeks.org/orientation-3-ordered-points/
    // for details of below formula.
    val := (q.y - p.y) * (r.x - q.x) -
              (q.x - p.x) * (r.y - q.y);
 
    if (val == 0) {
        return 0;  // colinear
    }
    
    if val > 0 {
        return 1
    }
    
    return 2
}
 
// The main function that returns true if line segment 'p1q1'
// and 'p2q2' intersect.
func doIntersect(p1, q1, p2, q2 point) bool {
    // Find the four orientations needed for general and
    // special cases
    o1 := orientation(p1, q1, p2);
    o2 := orientation(p1, q1, q2);
    o3 := orientation(p2, q2, p1);
    o4 := orientation(p2, q2, q1);
 
    // General case
    if (o1 != o2 && o3 != o4) {
        return true;
    }
 
    // Special Cases
    // p1, q1 and p2 are colinear and p2 lies on segment p1q1
    if (o1 == 0 && onSegment(p1, p2, q1)) {
        return true;
    }
 
    // p1, q1 and p2 are colinear and q2 lies on segment p1q1
    if (o2 == 0 && onSegment(p1, q2, q1)) {
        return true;
    }
 
    // p2, q2 and p1 are colinear and p1 lies on segment p2q2
    if (o3 == 0 && onSegment(p2, p1, q2)) {
        return true;
    }
 
     // p2, q2 and q1 are colinear and q1 lies on segment p2q2
    if (o4 == 0 && onSegment(p2, q1, q2)) {
        return true;
    }
 
    return false; // Doesn't fall in any of the above cases
}
package chipmunk

import (
	"./vect"
    "math"
)

type tri struct {
    a, b, c vect.Vect
}

func PolygonMomentOptimized(poly *PolygonShape, mass float32) vect.Float {
    var retval vect.Float
    retval = 0
    
    centroid := centroidForPoly(poly.Verts)
    
    for i := 0; i < poly.NumVerts; i++ {
		v1 := poly.Verts[i]
		v2 := poly.Verts[(i+1)%poly.NumVerts]
        
        t := tri {a: v1, b: v2, c: centroid}
        
        retval += (iForTri(t) + vect.Float(math.Pow(float64(vect.Sub(centroid, centroidForTri(t)).Length()), 2)*float64(mass/float32(poly.NumVerts))))
    }
    
    return retval
}

func centroidForPoly(verts Vertices) vect.Vect {
	var sum vect.Float
    sum = 0
	vsum := vect.Vector_Zero
	
	for i := 0; i < len(verts); i++ {
		v1 := verts[i]
		v2 := verts[(i+1)%len(verts)]
		cross := vect.Cross(v1, v2)
		
		sum += cross
		vsum = vect.Add(vsum, vect.Mult(vect.Add(v1, v2), cross))
	}
	
	return vect.Mult(vsum, 1.0/(3.0*sum))
}

func iForTri(t tri) vect.Float {
    b := float64(math.Max(float64(t.a.Length()), math.Max(float64(t.b.Length()), float64(t.c.Length()))))
    
    h := float64(heightForTri(t, vect.Float(b)))
    
    // a = sqrt(x² - h²)
    var x vect.Float
    if vect.Float(b) == t.a.Length() {
        x = t.b.Length()
    } else {
        x = t.a.Length()
    }
    a := math.Sqrt(math.Pow(float64(x), 2) - math.Pow(h, 2))
    
    return vect.Float((math.Pow(b, 3)*h-math.Pow(b, 2)*h*a+b*h*math.Pow(a, 2)+b*math.Pow(h, 3))/36)
}

func heightForTri(t tri, b vect.Float) vect.Float {
    s := (t.a.Length() + t.b.Length() + t.c.Length()) / 2
    // 1/2bh = sqr(s(s-a)(s-b)(s-c)
    return vect.Float(math.Sqrt(float64(s*(s - t.a.Length())*(s - t.b.Length())*(s - t.c.Length()))))/(b/2.0)
}

func centroidForTri(t tri) vect.Vect {
    return vect.Vect {
        (t.a.X + t.b.X + t.c.X) / 3,
        (t.a.Y + t.b.Y + t.c.Y) / 3 }
}
package fauxgl

type Triangle struct {
	V1, V2, V3 Vertex
}

func NewTriangle(v1, v2, v3 Vertex) *Triangle {
	t := Triangle{v1, v2, v3}
	t.FixNormals()
	return &t
}

func NewTriangleForPoints(p1, p2, p3 Vector) *Triangle {
	v1 := Vertex{Position: p1}
	v2 := Vertex{Position: p2}
	v3 := Vertex{Position: p3}
	return NewTriangle(v1, v2, v3)
}

func (t *Triangle) Normal() Vector {
	e1 := t.V2.Position.Sub(t.V1.Position)
	e2 := t.V3.Position.Sub(t.V1.Position)
	return e1.Cross(e2).Normalize()
}

func (t *Triangle) FixNormals() {
	n := t.Normal()
	zero := Vector{}
	if t.V1.Normal == zero {
		t.V1.Normal = n
	}
	if t.V2.Normal == zero {
		t.V2.Normal = n
	}
	if t.V3.Normal == zero {
		t.V3.Normal = n
	}
}

func (t *Triangle) BoundingBox() Box {
	min := t.V1.Position.Min(t.V2.Position).Min(t.V3.Position)
	max := t.V1.Position.Max(t.V2.Position).Max(t.V3.Position)
	return Box{min, max}
}

func (t *Triangle) Transform(matrix Matrix) {
	t.V1.Position = matrix.MulPosition(t.V1.Position)
	t.V2.Position = matrix.MulPosition(t.V2.Position)
	t.V3.Position = matrix.MulPosition(t.V3.Position)
	t.V1.Normal = matrix.MulDirection(t.V1.Normal)
	t.V2.Normal = matrix.MulDirection(t.V2.Normal)
	t.V3.Normal = matrix.MulDirection(t.V3.Normal)
}

package som

const (
	TopologyLinear      = "linear"      // 1-dimensional space with linear topology
	TopologyRectangular = "rectangular" // 2-dimensional space with rectangular topology
	TopologyHexagonal   = "hexagonal"   // 2-dimensional space with hexagonal topology
)

const (
	defaultPointChanSize = 10
)

type Generator <-chan Point

func makePointChan(limit int) chan Point {
	if limit <= 0 {
		return make(chan Point)
	}
	return make(chan Point, min(defaultPointChanSize, limit))
}

// NewLinearGenerator creates a generator for a linear grid with number of points equals to limit
func NewLinearGenerator(limit int) Generator {
	var (
		g = makePointChan(limit)
	)
	go func() {
		for i := 0; i < limit; i++ {
			g <- Point{float64(i)}
		}
		close(g)
	}()
	return g

}

// GridShape returns the shape of a grid.
// If limit is not provided, the function returns 0, 0.
// If only one limit is provided, the function returns regular grid with X = Y = limit.
// If two limits are provided, the function returns a grid with X = limit[0], Y = limit[1]
func GridShape(limit ...int) (int, int) {
	if len(limit) == 0 {
		return 0, 0
	}
	switch len(limit) {
	case 0:
		return 0, 0
	case 1:
		return limit[0], limit[0]
	default:
		return limit[0], limit[1]
	}
}

// NewRectangularGenerator creates a generator for a rectangular grid with shape equals to shape
// If shape is not provided, the function returns a generator does not produce any point
// If shape is single element, the function returns a generator for a square grid with X = Y = shape[0]
// If shapes is a slice of two elements, the function returns a generator for a grid with X = shape[0], Y = shape[1]
func NewRectangularGenerator(shape ...int) Generator {
	var (
		X, Y = GridShape(shape...)
		g    = makePointChan(X * Y)
	)
	go func() {
		for y := range X {
			for x := range Y {
				g <- Point{float64(x), float64(y)}
			}
		}
		close(g)
	}()
	return g
}

// NewHexagonalGenerator creates a generator for a hexagonal grid with shape equals to shape
// If shape is not provided, the function returns a generator does not produce any point
// If shape is single element, the function returns a generator for a square grid with X = Y = shape[0]
// If shapes is a slice of two elements, the function returns a generator for a grid with X = shape[0], Y = shape[1]
func NewHexagonalGenerator(shape ...int) Generator {
	const (
		horizontalDistance = 1.5
		verticalDistance   = 1.7320508075688772
		verticalSpace      = 0.5 * verticalDistance
	)

	var (
		I, J             = GridShape(shape...)
		g                = makePointChan(I * J)
		vSpacing         = []float64{0, verticalSpace}
		xOffset, yOffset = -1.0, -verticalDistance

		x, y float64
	)
	go func() {
		for range I {
			x = xOffset
			yOffset += verticalDistance
			for j := range J {
				y = yOffset + vSpacing[j%2]
				x += horizontalDistance
				g <- Point{x, y}
			}
		}
		close(g)
	}()
	return g

}

func NewGenerator(topology string, shape ...int) Generator {
	switch topology {
	case TopologyLinear:
		return NewLinearGenerator(shape[0])
	case TopologyRectangular:
		return NewRectangularGenerator(shape...)
	case TopologyHexagonal:
		return NewHexagonalGenerator(shape...)
	}
	return nil
}

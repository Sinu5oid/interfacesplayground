package interfaces

import (
	"fmt"
	"math"
	"strings"
)

type QuadraticEquation struct {
	a float64
	b float64
	c float64
}

func NewQuadraticEquation(a float64, b float64, c float64) QuadraticEquation {
	return QuadraticEquation{
		a: a,
		b: b,
		c: c,
	}
}

func (e QuadraticEquation) Solve() []complex128 {
	result := make([]complex128, 0, 2)
	discriminant := e.Discriminant()

	common := complex(-e.b/(2*e.a), 0)

	if discriminant == 0 {
		result = append(result, common)
	}

	var discriminantPart complex128
	if discriminant > 0 {
		discriminantPart = complex(math.Sqrt(discriminant), 0)
	} else {
		discriminantPart = complex(math.Sqrt(-discriminant), 1)
	}

	discriminantPart = discriminantPart / complex(2*e.a, 0)

	result = append(result, common-discriminantPart)
	result = append(result, common+discriminantPart)

	return result
}

func (e QuadraticEquation) Find(x complex128) complex128 {
	return complex(e.a, 0)*x*x + complex(e.b, 0)*x + complex(e.c, 0)
}

func (e QuadraticEquation) GetMin() (complex128, complex128) {
	if e.a > 0 {
		node := complex(-e.b/(2*e.a), 0)
		return node, e.Find(node)
	}

	inf := complex(math.Inf(-1), 0)
	return inf, inf
}

func (e QuadraticEquation) GetMax() (complex128, complex128) {
	if e.a < 0 {
		node := complex(-e.b/(2*e.a), 0)
		return node, e.Find(node)
	}

	inf := complex(math.Inf(1), 0)
	return inf, inf
}

func (e QuadraticEquation) GetMonotonicIntervals() []MonotonicInterval {
	result := make([]MonotonicInterval, 0, 2)

	if e.a == 0 {
		// trivial case (linear function)
		if e.b == 0 {
			// trivial case (constant function)
			result = append(result, MonotonicInterval{
				leftBound:         complex(math.Inf(-1), 0),
				includeLeftBound:  false,
				rightBound:        complex(math.Inf(1), 0),
				includeRightBound: false,
				monotonicType:     Stable,
			})

			return result
		}

		monotonicType := Ascending
		if e.b < 0 {
			monotonicType = Descending
		}

		result = append(result, MonotonicInterval{
			leftBound:         complex(math.Inf(-1), 0),
			includeLeftBound:  false,
			rightBound:        complex(math.Inf(1), 0),
			includeRightBound: false,
			monotonicType:     monotonicType,
		})

		return result
	}

	if e.a > 0 {
		// parabola branches are heading up
		minX, _ := e.GetMin()

		result = append(result, MonotonicInterval{
			leftBound:         complex(math.Inf(-1), 0),
			includeLeftBound:  false,
			rightBound:        minX,
			includeRightBound: true,
			monotonicType:     Descending,
		})

		result = append(result, MonotonicInterval{
			leftBound:         minX,
			includeLeftBound:  true,
			rightBound:        complex(math.Inf(1), 0),
			includeRightBound: false,
			monotonicType:     Ascending,
		})

		return result
	}

	// parabola branches are heading down
	maxX, _ := e.GetMax()

	result = append(result, MonotonicInterval{
		leftBound:         complex(math.Inf(-1), 0),
		includeLeftBound:  false,
		rightBound:        maxX,
		includeRightBound: true,
		monotonicType:     Ascending,
	})

	result = append(result, MonotonicInterval{
		leftBound:         maxX,
		includeLeftBound:  true,
		rightBound:        complex(math.Inf(1), 0),
		includeRightBound: false,
		monotonicType:     Descending,
	})

	return result
}

func (e QuadraticEquation) String() string {
	equationParts := make([]string, 0, 3)
	if e.a != 0 {
		equationParts = append(equationParts, fmt.Sprintf("%f * x^2", e.a))
	}

	if e.b != 0 {
		equationParts = append(equationParts, fmt.Sprintf("%+f * x", e.b))
	}

	if e.c != 0 {
		if e.a == 0 && e.b == 0 {
			return fmt.Sprintf("%f != 0", e.c)
		}

		equationParts = append(equationParts, fmt.Sprintf("%+f", e.c))
	}

	if len(equationParts) == 0 {
		equationParts = append(equationParts, "0")
	}

	return fmt.Sprintf("%s = 0", strings.Join(equationParts, " "))
}

func (e QuadraticEquation) TypeRepresentative() string {
	if e.a == 0 {
		if e.b == 0 {
			return "type: constant"
		}

		return "type: linear"
	}
	return "type: quadratic"
}

func (e QuadraticEquation) RootsRepresentative() string {
	roots := e.Solve()

	var result string
	if len(roots) == 1 {
		result = fmt.Sprintf("double root: x = %v", Representative(roots[0]))
	} else {
		result = fmt.Sprintf("roots: x1 = %v, x2 = %v", Representative(roots[0]), Representative(roots[1]))
	}

	return result
}

func (e QuadraticEquation) MinRepresentative() string {
	minX, minY := e.GetMin()

	return fmt.Sprintf("min: (%v; %v)", Representative(minX), Representative(minY))
}

func (e QuadraticEquation) MaxRepresentative() string {
	maxX, maxY := e.GetMax()

	return fmt.Sprintf("max: (%v; %v)", Representative(maxX), Representative(maxY))
}

func (e QuadraticEquation) MonotonicIntervalsRepresentative() string {
	intervalsRaw := e.GetMonotonicIntervals()
	intervals := make([]string, 0, 2)

	for _, interval := range intervalsRaw {
		intervals = append(intervals, fmt.Sprintf("%s", interval))
	}

	return fmt.Sprintf("monotonic intervals: %s", strings.Join(intervals, "\t"))
}

func (e QuadraticEquation) Representative() string {
	return fmt.Sprintf(
		"%s\ncanonical view: %s\n%s\n%s\n%s\n%s",
		e.TypeRepresentative(),
		e.String(),
		e.RootsRepresentative(),
		e.MinRepresentative(),
		e.MaxRepresentative(),
		e.MonotonicIntervalsRepresentative(),
	)
}

func (e QuadraticEquation) Discriminant() float64 {
	return math.Pow(e.b, 2) - (4 * e.a * e.c)
}

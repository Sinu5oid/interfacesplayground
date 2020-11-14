package interfaces

import (
	"fmt"
	"math"
)

type ConstantEquation struct {
	a float64
}

func NewConstantEquation(a float64) ConstantEquation {
	return ConstantEquation{a: a}
}

func (e ConstantEquation) Solve() (roots []complex128) {
	if e.a != 0 {
		return []complex128{complex(math.NaN(), 0)}
	}

	return []complex128{
		complex(math.Inf(-1), 0),
		complex(math.Inf(1), 0),
	}
}

func (e ConstantEquation) Find(_ complex128) complex128 {
	return complex(e.a, 0)
}

func (e ConstantEquation) GetMin() (complex128, complex128) {
	return complex(math.Inf(-1), 0), complex(e.a, 0)
}

func (e ConstantEquation) GetMax() (complex128, complex128) {
	return complex(math.Inf(1), 0), complex(e.a, 0)
}

func (e ConstantEquation) GetMonotonicIntervals() []MonotonicInterval {
	return []MonotonicInterval{
		{
			leftBound:         complex(math.Inf(-1), 0),
			includeLeftBound:  false,
			rightBound:        complex(math.Inf(1), 0),
			includeRightBound: false,
			monotonicType:     Stable,
		},
	}
}

func (e ConstantEquation) String() string {
	if e.a != 0 {
		return fmt.Sprintf("%f != 0", e.a)
	}

	return fmt.Sprintf("%f = 0", e.a)
}

func (e ConstantEquation) TypeRepresentative() string {
	return "type: constant"
}

func (e ConstantEquation) RootsRepresentative() string {
	roots := e.Solve()

	if len(roots) == 1 {
		return fmt.Sprintf("root: %s", Representative(roots[0]))
	}

	return fmt.Sprintf("roots: %s, %s", Representative(roots[0]), Representative(roots[1]))
}

func (e ConstantEquation) MinRepresentative() string {
	minX, minY := e.GetMin()

	return fmt.Sprintf("min: (%v; %v)", Representative(minX), Representative(minY))
}

func (e ConstantEquation) MaxRepresentative() string {
	minX, minY := e.GetMax()

	return fmt.Sprintf("max: (%v; %v)", Representative(minX), Representative(minY))
}

func (e ConstantEquation) MonotonicIntervalsRepresentative() string {
	return fmt.Sprintf("monotonic interval: %s", e.GetMonotonicIntervals()[0])
}

func (e ConstantEquation) Representative() string {
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

package interfaces

import (
	"fmt"
	"math"
	"strings"
)

type LinearEquation struct {
	k float64
	b float64
}

func NewLinearEquation(k float64, b float64) LinearEquation {
	return LinearEquation{
		k: k,
		b: b,
	}
}

func (e LinearEquation) Solve() []complex128 {
	result := make([]complex128, 0, 2)
	if e.k == 0 {
		if e.b != 0 {
			result = append(result, complex(math.NaN(), 0))
			return result
		}

		result = append(result, complex(math.Inf(-1), 0))
		result = append(result, complex(math.Inf(1), 0))
		return result
	}

	root := -e.b / e.k
	result = append(result, complex(root, 0))
	return result
}

func (e LinearEquation) Find(x complex128) complex128 {
	return complex(e.k, 0)*x + complex(e.b, 0)
}

func (e LinearEquation) GetMin() (complex128, complex128) {
	if e.k == 0 {
		return complex(math.Inf(-1), 0), complex(e.b, 0)
	}

	if e.k > 0 {
		return complex(math.Inf(-1), 0), complex(math.Inf(-1), 0)
	}

	return complex(math.Inf(1), 0), complex(math.Inf(-1), 0)
}

func (e LinearEquation) GetMax() (complex128, complex128) {
	if e.k == 0 {
		return complex(math.Inf(1), 0), complex(e.b, 0)
	}

	if e.k > 0 {
		return complex(math.Inf(1), 0), complex(math.Inf(1), 0)
	}

	return complex(math.Inf(-1), 0), complex(math.Inf(1), 0)
}

func (e LinearEquation) GetMonotonicIntervals() []MonotonicInterval {
	result := make([]MonotonicInterval, 0, 1)

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

func (e LinearEquation) String() string {
	equationParts := make([]string, 0, 2)

	if e.k != 0 {
		equationParts = append(equationParts, fmt.Sprintf("%+f * x", e.k))
	}

	if e.b != 0 {
		equationParts = append(equationParts, fmt.Sprintf("%+f", e.b))
	}

	if len(equationParts) == 0 {
		equationParts = append(equationParts, "0")
	}

	return fmt.Sprintf("%s = 0", strings.Join(equationParts, " "))
}

func (e LinearEquation) TypeRepresentative() string {
	if e.k == 0 {
		return "type: constant"
	}

	return "type: linear"
}

func (e LinearEquation) RootsRepresentative() string {
	roots := e.Solve()

	if len(roots) == 1 {
		return fmt.Sprintf("root: %s", Representative(roots[0]))
	}

	return fmt.Sprintf("roots: %s, %s", Representative(roots[0]), Representative(roots[1]))
}

func (e LinearEquation) MinRepresentative() string {
	minX, minY := e.GetMin()

	return fmt.Sprintf("min: (%v; %v)", Representative(minX), Representative(minY))
}

func (e LinearEquation) MaxRepresentative() string {
	minX, minY := e.GetMax()

	return fmt.Sprintf("max: (%v; %v)", Representative(minX), Representative(minY))
}

func (e LinearEquation) MonotonicIntervalsRepresentative() string {
	return fmt.Sprintf("monotonic interval: %s", e.GetMonotonicIntervals()[0])
}

func (e LinearEquation) Representative() string {
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

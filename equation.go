package interfaces

import "errors"

type OneArgumentEquation interface {
	Solve() (roots []complex128)
	Find(x complex128) complex128
	GetMin() (complex128, complex128)
	GetMax() (complex128, complex128)
	GetMonotonicIntervals() []MonotonicInterval

	// Representatives
	String() string
	TypeRepresentative() string
	RootsRepresentative() string
	MinRepresentative() string
	MaxRepresentative() string
	MonotonicIntervalsRepresentative() string
	Representative() string
}

func NewOneArgumentEquation(a float64, args ...float64) (OneArgumentEquation, error) {
	switch len(args) {
	case 0:
		return NewConstantEquation(a), nil
	case 1:
		if a == 0 {
			return NewConstantEquation(args[0]), nil
		}
		return NewLinearEquation(a, args[0]), nil
	case 2:
		if a == 0 {
			if args[0] == 0 {
				return NewConstantEquation(args[1]), nil
			}

			return NewLinearEquation(args[0], args[1]), nil
		}
		return NewQuadraticEquation(a, args[0], args[1]), nil
	default:
		return nil, errors.New("not implemented")
	}
}

package interfaces

import "fmt"

type MonotonicType string

const (
	Ascending  MonotonicType = "ascending"
	Descending MonotonicType = "descending"
	Stable     MonotonicType = "stable"
)

type MonotonicInterval struct {
	leftBound         complex128
	includeLeftBound  bool
	rightBound        complex128
	includeRightBound bool
	monotonicType     MonotonicType
}

func (i MonotonicInterval) String() string {
	leftBoundDelimiter := "("
	if i.includeLeftBound {
		leftBoundDelimiter = "["
	}

	rightBoundDelimiter := ")"
	if i.includeRightBound {
		rightBoundDelimiter = "]"
	}

	return fmt.Sprintf(
		"%s%s; %s%s - %s",
		leftBoundDelimiter,
		Representative(i.leftBound),
		Representative(i.rightBound),
		rightBoundDelimiter,
		i.monotonicType,
	)
}

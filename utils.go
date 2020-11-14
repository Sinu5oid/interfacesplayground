package interfaces

import "fmt"

func Representative(n complex128) string {
	if imag(n) == 0 {
		return fmt.Sprintf("%v", real(n))
	}

	return fmt.Sprintf("%v", n)
}

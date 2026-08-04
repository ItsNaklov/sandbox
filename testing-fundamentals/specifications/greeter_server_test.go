package specifications_test

import (
	"testing"

	"testing-fundamentals/specifications"
)

func TestGreeterServer(t *testing.T) {
	specifications.GreetSpecification(t, nil)
}

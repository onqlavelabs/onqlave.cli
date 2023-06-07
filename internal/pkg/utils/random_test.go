package utils

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandom(t *testing.T) {
	var slice = SliceIntRange(0, 100, 5)
	for range slice {
		var slice = SliceIntRange(0, 100, 5)
		fmt.Printf("Randome slice int range: %v\n", slice)
	}

	for range slice {
		var num = IntRange(0, 100)
		fmt.Printf("Randome int range: %v\n", num)
	}

	for range slice {
		var str = RandomString(5, Alphabet)
		fmt.Printf("Randome string: %v\n", str)
	}

	for range slice {
		var str = StringInRange(1, 20, Alphanumeric)
		fmt.Printf("Randome string: %v\n", str)
	}
	assert.Equal(t, len(slice), 5)
}

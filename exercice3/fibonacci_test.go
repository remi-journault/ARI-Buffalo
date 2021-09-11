package exercice3

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFibonacci(t *testing.T) {
	// 0,1,2,3,5,8,13,21
	f := fibonacci()
	for i := 0; i < 10; i++ {
		res := f()
		fmt.Println(res)
		switch i {
		case 0:
			assert.Equal(t, 0, res)
		case 1:
			assert.Equal(t, 1, res)
		case 2:
			assert.Equal(t, 1, res)
		case 3:
			assert.Equal(t, 2, res)
		case 4:
			assert.Equal(t, 3, res)
		case 5:
			assert.Equal(t, 5, res)
		case 6:
			assert.Equal(t, 8, res)
		case 7:
			assert.Equal(t, 13, res)
		case 8:
			assert.Equal(t, 21, res)
		}
	}

	assert.NotNil(t, f)
}

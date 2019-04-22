package hard

import (
	"testing"

	"github.com/stretchr/testify/require"
)

/*
"1 + 1" = 2
" 6-4 / 2 " = 4
"2*(5+5*2)/3+(6/2+8)" = 21
"(2+6* 3+5- (3*14/7+2)*5)+3"=-12
*/
func TestCalculate(t *testing.T) {
	r := require.New(t)
	r.Equal(2, calculate("1 + 1"))
	r.Equal(4, calculate(" 6-4 / 2 "))
	r.Equal(21, calculate("2*(5+5*2)/3+(6/2+8)"))
	r.Equal(-12, calculate("(2+6* 3+5- (3*14/7+2)*5)+3"))
}

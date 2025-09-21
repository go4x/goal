package mathx

import (
	"fmt"
	"testing"

	"github.com/go4x/goal/assert"
)

func TestFormatCommaInt(t *testing.T) {
	d := 123456789
	s := FmtCommaInt(int64(d))
	fmt.Println(s)
	fmt.Println(s == "123,456,789")
	assert.Equals("123,456,789", s)
}

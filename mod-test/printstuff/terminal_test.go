package printstuff

import (
	"testing"

	myassert "github.com/stretchr/testify/assert"
)

func TestSomething(t *testing.T) {

	// assert equality
	myassert.Equal(t, 123, 123, "they should be equal")

	// assert inequality
	myassert.NotEqual(t, 123, 456, "they should not be equal")

	var object *int = nil
	// assert for nil (good for errors)
	myassert.Nil(t, object)

	// assert for not nil (good when you expect something)
	// if myassert.NotNil(t, object) {
	//
	// 	// now we know that object isn't nil, we are safe to make
	// 	// further assertions without causing any errors
	// 	myassert.Equal(t, "Something", object.Value)
	//
	// }
}

package design

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./design/ -run ^TestMonarchy$
func TestMonarchy(t *testing.T) {
	a := assert.New(t)

	var monarchy = NewMonarchy("Jake")
	monarchy.Birth("Catherine", "Jake")
	monarchy.Birth("Tom", "Jake")
	monarchy.Birth("Celine", "Jake")
	monarchy.Birth("Jane", "Catherine")
	monarchy.Birth("Peter", "Celine")
	monarchy.Birth("Farah", "Jane")
	monarchy.Birth("Mark", "Catherine")

	list := monarchy.GetOrderOfSuccession()
	a.Equal([]string{"Jake", "Catherine", "Jane", "Farah", "Mark", "Tom", "Celine", "Peter"}, list)

	monarchy.Death("Jake")
	monarchy.Death("Jane")

	list = monarchy.GetOrderOfSuccession()
	a.Equal([]string{"Catherine", "Farah", "Mark", "Tom", "Celine", "Peter"}, list)
}

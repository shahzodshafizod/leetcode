package strings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./strings/ -run ^TestToGoatLatin$
func TestToGoatLatin(t *testing.T) {
	for _, tc := range []struct {
		sentence string
		goated   string
	}{
		{sentence: "I speak Goat Latin", goated: "Imaa peaksmaaa oatGmaaaa atinLmaaaaa"},
		{sentence: "The quick brown fox jumped over the lazy dog", goated: "heTmaa uickqmaaa rownbmaaaa oxfmaaaaa umpedjmaaaaaa overmaaaaaaa hetmaaaaaaaa azylmaaaaaaaaa ogdmaaaaaaaaaa"},
	} {
		goated := toGoatLatin(tc.sentence)
		assert.Equal(t, tc.goated, goated)
	}
}

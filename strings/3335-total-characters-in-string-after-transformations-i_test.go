package strings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./strings/ -run ^TestLengthAfterTransformations$
func TestLengthAfterTransformations(t *testing.T) {
	for _, tc := range []struct {
		s      string
		t      int
		length int
	}{
		{s: "z", t: 100, length: 16},
		{s: "azbk", t: 1, length: 5},
		{s: "abcyy", t: 2, length: 7},
		{s: "zwzzugozhwsycuocakus", t: 1, length: 24},
		// {s: "jqktcurgdvlibczdsvnsg", t: 7517, length: 79033769},
		// {s: "hrzmawnweztcskakojfahyvnoctsctwsbagyqmmoheldlpzctduxmhfcwqcbvovoyswjtdzvsheoofocknqddfsjwxfuuhvznxry", t: 1000, length: 652815408},
		// {s: "xdzbhxqcmhezajdhljzsgshikospdeyxrnwylcvcuvfppquqqxcfbvmdlwbzkxjkwzvoyvmpnlxuyulexoqgayvxlvofyjhmxshfprpbhjywofbqhhufezuyccasrodkzmxkwzfhcfxhlrpidoklhgidflvyppajzgecuhumjyglgzqzcusdniuqgylpxlhkknwbwehtaabnioerjnpxxjqhmxftsoukxbfkndssniyhwqfmtcoerxrkkdjepiyrhmnepuwaunubwrixahwaoecretfzbqavlhzavdherptjpkhqrkpopdheswffikuxrvqohccbyphcrirhhjddjqihwxlszdehalyoqqzsimaaxepttwbfpbtqgwhidvzoegkjeqdhndszmrtgloqwerpdsvqhdvmfqxwdmkocqcltqiojgpstzainiukaejurbvuvjbyyodruvuliahiscdmrjmnthehnmhovjwakenmuwbxqeoox", t: 100000, length: 571676932},
	} {
		length := lengthAfterTransformations(tc.s, tc.t)
		assert.Equal(t, tc.length, length)
	}
}

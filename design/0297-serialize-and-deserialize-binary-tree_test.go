package design

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./design/ -run ^TestCodec$
func TestCodec(t *testing.T) {
	ser := NewCodec()
	deser := NewCodec()
	for _, tc := range []struct {
		original *TreeNode
		restored *TreeNode
	}{
		{original: MakeTree2(1, 2, 3, nil, nil, 4, 5), restored: MakeTree2(1, 2, 3, nil, nil, 4, 5)},
		{original: MakeTree2(), restored: MakeTree2()},
		{original: MakeTree2(1, 2, 3, nil, nil, 4, 5, 6, 7), restored: MakeTree2(1, 2, 3, nil, nil, 4, 5, 6, 7)},
		{original: MakeTree2(1), restored: MakeTree2(1)},
		{original: MakeTree2(1, -1000), restored: MakeTree2(1, -1000)},
	} {
		data := ser.serialize(tc.original)
		restored := deser.deserialize(data)
		assert.Equal(t, tc.restored, restored)
	}
}

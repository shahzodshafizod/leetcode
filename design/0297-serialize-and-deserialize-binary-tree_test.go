package design

import (
	"testing"

	"github.com/shahzodshafizod/leetcode/pkg"
	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./design/ -run ^TestCodec$
func TestCodec(t *testing.T) {
	ser := NewCodec()
	deser := NewCodec()
	for _, tc := range []struct {
		original *pkg.TreeNode
		restored *pkg.TreeNode
	}{
		{original: pkg.MakeTree2(1, 2, 3, nil, nil, 4, 5), restored: pkg.MakeTree2(1, 2, 3, nil, nil, 4, 5)},
		{original: pkg.MakeTree2(), restored: pkg.MakeTree2()},
		{original: pkg.MakeTree2(1, 2, 3, nil, nil, 4, 5, 6, 7), restored: pkg.MakeTree2(1, 2, 3, nil, nil, 4, 5, 6, 7)},
		{original: pkg.MakeTree2(1), restored: pkg.MakeTree2(1)},
		{original: pkg.MakeTree2(1, -1000), restored: pkg.MakeTree2(1, -1000)},
	} {
		data := ser.serialize(tc.original)
		restored := deser.deserialize(data)
		assert.Equal(t, tc.restored, restored)
	}
}

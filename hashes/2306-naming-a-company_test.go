package hashes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./hashes/ -run ^TestDistinctNames$
func TestDistinctNames(t *testing.T) {
	for _, tc := range []struct {
		ideas []string
		count int64
	}{
		{ideas: []string{"coffee", "donuts", "time", "toffee"}, count: 6},
		{ideas: []string{"lack", "back"}, count: 0},
		{ideas: []string{"a", "b", "baby", "aby", "banana", "apple"}, count: 8},
		// {ideas: []string{"j", "t", "fmucd", "liantvj", "sgizj", "qjkqmnefo", "g", "zaa", "epw", "fz", "ycad", "oocqnz", "bwowse", "bcnealx", "imlwxoxzml", "crmahpv", "vh", "qfwyd", "dycxopdrzb", "hvbje", "f", "qwowse", "lnwgxmqdc", "k", "fmffskrv", "bmlwxoxzml", "iffrcpdsu", "xpqodjcvri", "cmsyrkrrm", "unb", "ln", "hmffskrv", "bihbm", "fv", "eqwg", "p", "vnpooqdpe", "okja", "lmffskrv", "r", "pnwgxmqdc", "gpyamphhda", "azklwrs", "ejwuuqz", "rmxhrnzz", "gggwj", "qh", "siantvj", "pcad", "y", "nteqr", "mi", "ri", "jqwg", "fnpooqdpe", "ukqhjnk", "tyaastc", "xmsyrkrrm", "l", "ypmpilxnl", "lzf", "nkja", "xiantvj", "fvbje", "ooltwf", "fycxopdrzb", "gi", "bmffskrv", "bnihnwsx", "epci", "awowse", "tokjknsyk", "owan", "vz", "qd", "ng", "goebwzhefq", "xafrsm", "sjyfskc", "bfr", "nwan", "mjkj", "ujlr", "mrmahpv"}, count: 5808},
	} {
		count := distinctNames(tc.ideas)
		assert.Equal(t, tc.count, count)
	}
}

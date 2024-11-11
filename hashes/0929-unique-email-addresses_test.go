package hashes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 ./hashes/ -run ^TestNumUniqueEmails$
func TestNumUniqueEmails(t *testing.T) {
	for _, tc := range []struct {
		emails []string
		count  int
	}{
		{emails: []string{"a@leetcode.com", "b@leetcode.com", "c@leetcode.com"}, count: 3},
		{emails: []string{"test.email+alex@leetcode.com", "test.email.leet+alex@code.com"}, count: 2},
		{emails: []string{"test.email+alex@leetcode.com", "test.e.mail+bob.cathy@leetcode.com", "testemail+david@lee.tcode.com"}, count: 2},
		// {emails: []string{"fg.r.u.uzj+o.pw@kziczvh.com", "r.cyo.g+d.h+b.ja@tgsg.z.com", "fg.r.u.uzj+o.f.d@kziczvh.com", "r.cyo.g+ng.r.iq@tgsg.z.com", "fg.r.u.uzj+lp.k@kziczvh.com", "r.cyo.g+n.h.e+n.g@tgsg.z.com", "fg.r.u.uzj+k+p.j@kziczvh.com", "fg.r.u.uzj+w.y+b@kziczvh.com", "r.cyo.g+x+d.c+f.t@tgsg.z.com", "r.cyo.g+x+t.y.l.i@tgsg.z.com", "r.cyo.g+brxxi@tgsg.z.com", "r.cyo.g+z+dr.k.u@tgsg.z.com", "r.cyo.g+d+l.c.n+g@tgsg.z.com", "fg.r.u.uzj+vq.o@kziczvh.com", "fg.r.u.uzj+uzq@kziczvh.com", "fg.r.u.uzj+mvz@kziczvh.com", "fg.r.u.uzj+taj@kziczvh.com", "fg.r.u.uzj+fek@kziczvh.com"}, count: 2},
	} {
		count := numUniqueEmails(tc.emails)
		assert.Equal(t, tc.count, count)
	}
}

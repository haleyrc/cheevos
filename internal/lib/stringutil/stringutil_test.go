package stringutil_test

import (
	"testing"

	"github.com/haleyrc/cheevos/internal/assert"
	"github.com/haleyrc/cheevos/internal/lib/stringutil"
)

func TestMakeSafe(t *testing.T) {
	testcases := map[string]struct {
		input string
		want  string
	}{
		"removes all prefixed whitespace characters": {
			input: "\t\n \rstart",
			want:  "start",
		},
		"removes all postfixed whitespace characters": {
			input: "start\t\n \r",
			want:  "start",
		},
		"doesn't remove any infix whitespace characters": {
			input: "start\t\n \rend",
			want:  "start\t\n \rend",
		},
		"escapes html": {
			input: "<p>name</p>",
			want:  "&lt;p&gt;name&lt;/p&gt;",
		},
		"escapes javascript": {
			input: "alert('name');",
			want:  "alert(&#39;name&#39;);",
		},
	}
	for name, tc := range testcases {
		t.Run(name, func(t *testing.T) {
			got := stringutil.MakeSafe(tc.input)
			assert.String(t, "safe string", got).Equals(tc.want)
		})
	}
}

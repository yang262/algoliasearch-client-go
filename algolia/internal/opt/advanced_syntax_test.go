// Code generated by go generate. DO NOT EDIT.

package opt

import (
	"encoding/json"
	"testing"

	"github.com/algolia/algoliasearch-client-go/algolia/opt"
	"github.com/stretchr/testify/require"
)

func TestAdvancedSyntax(t *testing.T) {
	for _, c := range []struct {
		opts     []interface{}
		expected *opt.AdvancedSyntaxOption
	}{
		{
			opts:     []interface{}{nil},
			expected: opt.AdvancedSyntax(false),
		},
		{
			opts:     []interface{}{opt.AdvancedSyntax(true)},
			expected: opt.AdvancedSyntax(true),
		},
		{
			opts:     []interface{}{opt.AdvancedSyntax(false)},
			expected: opt.AdvancedSyntax(false),
		},
	} {
		var (
			in  = ExtractAdvancedSyntax(c.opts...)
			out opt.AdvancedSyntaxOption
		)
		data, err := json.Marshal(&in)
		require.NoError(t, err)
		err = json.Unmarshal(data, &out)
		require.NoError(t, err)
		require.Equal(t, *c.expected, out)
	}
}

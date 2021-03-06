// Code generated by go generate. DO NOT EDIT.

package opt

import (
	"github.com/algolia/algoliasearch-client-go/algolia/opt"
)

// ExtractUnretrievableAttributes returns the first found UnretrievableAttributesOption from the
// given variadic arguments or nil otherwise.
func ExtractUnretrievableAttributes(opts ...interface{}) *opt.UnretrievableAttributesOption {
	for _, o := range opts {
		if v, ok := o.(*opt.UnretrievableAttributesOption); ok {
			return v
		}
	}
	return nil
}

// Code generated by go generate. DO NOT EDIT.

package opt

import (
	"github.com/algolia/algoliasearch-client-go/algolia/opt"
)

// ExtractUserToken returns the first found UserTokenOption from the
// given variadic arguments or nil otherwise.
func ExtractUserToken(opts ...interface{}) *opt.UserTokenOption {
	for _, o := range opts {
		if v, ok := o.(*opt.UserTokenOption); ok {
			return v
		}
	}
	return nil
}

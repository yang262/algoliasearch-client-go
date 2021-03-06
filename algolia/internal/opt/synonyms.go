// Code generated by go generate. DO NOT EDIT.

package opt

import (
	"github.com/algolia/algoliasearch-client-go/algolia/opt"
)

// ExtractSynonyms returns the first found SynonymsOption from the
// given variadic arguments or nil otherwise.
func ExtractSynonyms(opts ...interface{}) *opt.SynonymsOption {
	for _, o := range opts {
		if v, ok := o.(*opt.SynonymsOption); ok {
			return v
		}
	}
	return nil
}

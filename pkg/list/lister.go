package list

import (
	"context"
	"k8s.io/apimachinery/pkg/runtime"
)

// Lister is an object that can retrieve resources that match the provided field and label criteria.
type Lister interface {
	// NewList returns an empty object that can be used with the List call.
	// This object must be a pointer type for use with Codec.DecodeInto([]byte, runtime.Object)
	NewList() runtime.Object
	// List selects resources in the storage which match to the selector. 'options' can be nil.
	List(ctx context.Context, options *ListOptions) (runtime.Object, error)
}

type ListOptions struct {
	Limit int
}

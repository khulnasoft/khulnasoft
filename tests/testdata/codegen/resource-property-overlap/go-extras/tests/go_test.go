package tests

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/khulnasoft/khulnasoft/sdk/v3/go/common/resource"
	"github.com/khulnasoft/khulnasoft/sdk/v3/go/khulnasoft"

	"resource-property-overlap/example"
)

// Tests that XArray{x}.ToXArrayOutput().Index(khulnasoft.Int(0)) == x.
func TestArrayOutputIndex(t *testing.T) {
	err := khulnasoft.RunErr(func(ctx *khulnasoft.Context) error {

		r1, err := example.NewRec(ctx, "rec1", &example.RecArgs{})
		if err != nil {
			return err
		}

		r1o := r1.ToRecOutput()

		r2o := example.RecArray{r1o}.ToRecArrayOutput().
			Index(khulnasoft.Int(0))

		wg := &sync.WaitGroup{}
		wg.Add(1)

		khulnasoft.All(r1o, r2o).ApplyT(func(xs []interface{}) int {
			assert.Equal(t, xs[0], xs[1])
			wg.Done()
			return 0
		})

		wg.Wait()
		return nil
	}, khulnasoft.WithMocks("project", "stack", mocks(0)))
	assert.NoError(t, err)
}

type mocks int

func (mocks) NewResource(args khulnasoft.MockResourceArgs) (string, resource.PropertyMap, error) {
	return args.Name + "_id", args.Inputs, nil
}

func (mocks) Call(args khulnasoft.MockCallArgs) (resource.PropertyMap, error) {
	return args.Args, nil
}

// Code generated from Pkl module `flyinprogrammer.pkl_go_bananas.Earth`. DO NOT EDIT.
package earth

import (
	"context"

	"github.com/apple/pkl-go/pkl"
)

type Earth struct {
}

// LoadFromPath loads the pkl module at the given path and evaluates it into a Earth
func LoadFromPath(ctx context.Context, path string) (ret *Earth, err error) {
	evaluator, err := pkl.NewEvaluator(ctx, pkl.PreconfiguredOptions)
	if err != nil {
		return nil, err
	}
	defer func() {
		cerr := evaluator.Close()
		if err == nil {
			err = cerr
		}
	}()
	ret, err = Load(ctx, evaluator, pkl.FileSource(path))
	return ret, err
}

// Load loads the pkl module at the given source and evaluates it with the given evaluator into a Earth
func Load(ctx context.Context, evaluator pkl.Evaluator, source *pkl.ModuleSource) (*Earth, error) {
	var ret Earth
	if err := evaluator.EvaluateModule(ctx, source, &ret); err != nil {
		return nil, err
	}
	return &ret, nil
}

package v3

import (
	"fmt"

	"github.com/xeger/semdiff/types"
)

func Diff(a, b *OpenAPI) (summary *types.ChangeSet) {
	ap := indexPaths(a)
	bp := indexPaths(b)

	pathDiffs := diffOperations(ap, bp)

	return types.NewChangeSet(pathDiffs)
}

func diffOperations(ap, bp map[string]*Operation) (diffs []*types.Change) {
	uniqB := make(map[string]*Operation)
	uniqA := make(map[string]*Operation)

	// Find paths in B but not in A
	for k, v := range bp {
		if _, ok := ap[k]; !ok {
			uniqB[k] = v
		}
	}

	// Find paths in A but not in B
	for k, v := range ap {
		if _, ok := bp[k]; !ok {
			uniqA[k] = v
		}
	}

	for _, o := range uniqB {
		diffs = append(diffs, &types.Change{
			Description: fmt.Sprintf("added operation"),
			Minor:       true,
			Name:        o.OperationID,
			Type:        "path",
		})
	}

	for _, o := range uniqA {
		diffs = append(diffs, &types.Change{
			Description: fmt.Sprintf("removed operation"),
			Major:       true,
			Name:        o.OperationID,
			Type:        "path",
		})
	}

	return diffs
}

func indexPaths(spec *OpenAPI) map[string]*Operation {
	index := make(map[string]*Operation)

	for p, pi := range spec.Paths {
		// TODO: options, head, trace, connect, ...
		if pi.Get != nil {
			index[pathKey("GET", p)] = pi.Get
		}
		if pi.Post != nil {
			index[pathKey("POST", p)] = pi.Post
		}
		if pi.Put != nil {
			index[pathKey("PUT", p)] = pi.Put
		}
		if pi.Delete != nil {
			index[pathKey("DELETE", p)] = pi.Delete
		}
		if pi.Patch != nil {
			index[pathKey("PATCH", p)] = pi.Patch
		}
	}

	return index
}

func pathKey(method, path string) string {
	return method + " " + path
}

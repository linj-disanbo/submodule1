package sdk

import "github.com/linj-disanbo/submodule1/tx/types"

// CreateTx CreateTx
func CreateTx(exec string) *types.Tx {
	return &types.Tx{Exec: exec}
}

// SignTx SignTx
func SignTx(tx *types.Tx) {
}

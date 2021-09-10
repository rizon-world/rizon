package types_test

import (
	rizon "github.com/rizon-world/rizon/app"
	"github.com/stretchr/testify/require"
	"testing"

	"github.com/rizon-world/rizon/x/tokenswap/types"
)

func init() {
	rizon.SetRizonConfig()
}

func TestParamsEqual(t *testing.T) {
	p1 := types.DefaultParams()
	acc := p1.Validate()
	require.Equal(t, nil, acc, "Default Params is not validated")
}

package types_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/skip-mev/slinky/x/oracle/types"
)

func TestGenesisValidation(t *testing.T) {
	tcs := []struct {
		name       string
		cpgs       []types.CurrencyPairGenesis
		nextID     uint64
		expectPass bool
	}{
		{
			"if any of the currency-pair geneses are invalid - fail",
			[]types.CurrencyPairGenesis{
				{
					CurrencyPair: types.CurrencyPair{
						Base:  "AA",
						Quote: "BB",
					},
				},
				{
					// invalid CurrencyPairGenesis
					CurrencyPair: types.CurrencyPair{
						Base: "BB",
					},
				},
			},
			0,
			false,
		},
		{
			"if the CurrencyPairPrice is nil, but the nonce is non-zero - fail",
			[]types.CurrencyPairGenesis{
				{
					CurrencyPair: types.CurrencyPair{
						Base:  "AA",
						Quote: "BB",
					},
					Nonce: 10,
				},
			},
			0,
			false,
		},
		{
			"if all of the currency-pair geneses are valid - pass",
			[]types.CurrencyPairGenesis{
				{
					CurrencyPair: types.CurrencyPair{
						Base:  "AA",
						Quote: "BB",
					},
					Id: 0,
				},
				{
					// invalid CurrencyPairGenesis
					CurrencyPair: types.CurrencyPair{
						Base:  "BB",
						Quote: "CC",
					},
					Id: 1,
				},
			},
			2,
			true,
		},
		{
			"if any of the CurrencyPairGenesis ID's are duplicated - fail",
			[]types.CurrencyPairGenesis{
				{
					CurrencyPair: types.CurrencyPair{
						Base:  "AA",
						Quote: "BB",
					},
					Id: 1,
				},
				{
					CurrencyPair: types.CurrencyPair{
						Base:  "BB",
						Quote: "CC",
					},
					Id: 1,
				},
			},
			3,
			false,
		},
		{
			"if any of the CurrencyPairs are repeated - fail",
			[]types.CurrencyPairGenesis{
				{
					CurrencyPair: types.CurrencyPair{
						Base:  "AA",
						Quote: "BB",
					},
					Id: 1,
				},
				{
					CurrencyPair: types.CurrencyPair{
						Base:  "AA",
						Quote: "BB",
					},
					Id: 2,
				},
			},
			3,
			false,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			gs := types.NewGenesisState(tc.cpgs, tc.nextID)
			err := gs.Validate()

			if tc.expectPass {
				assert.Nil(t, err)
			} else {
				assert.NotNil(t, err)
			}
		})
	}
}

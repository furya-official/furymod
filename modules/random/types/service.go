package types

import (
	"github.com/tendermint/tendermint/crypto"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/furya-official/furymod/modules/oracle/types"
	servicetypes "github.com/furya-official/furymod/modules/service/types"
)

const (
	ServiceName          = "random"
	ServiceDesc          = "system service definition of random module"
	ServiceValueJSONPath = "seed"
	AuthorDescription    = "random module account"
	ServiceSchemas       = `
	{
		"input": {
			"$schema": "http://json-schema.org/draft-04/schema#",
			"title": "random-seed-input-body",
			"description": "FURY Hub Random Seed Input Body Schema",
			"type": "object",
			"additionalProperties": false
		},
		"output": {
			"$schema": "http://json-schema.org/draft-04/schema#",
			"title": "random-seed-output-body",
			"description": "FURY Hub Random Seed Output Body Schema",
			"type": "object",
			"properties": {
				"seed": {
					"description": "random seed",
					"type": "string",
					"pattern": "^[0-9a-fA-F]{64}$"
				}
			},
			"additionalProperties": false,
			"required": [
				"seed"
			]
		}
	}`
)

var (
	ServiceTags = []string{types.ModuleName}
	Author      = sdk.AccAddress(crypto.AddressHash([]byte(types.ModuleName)))
)

func GetSvcDefinition() servicetypes.ServiceDefinition {
	return servicetypes.NewServiceDefinition(
		ServiceName,
		ServiceDesc,
		ServiceTags,
		Author,
		AuthorDescription,
		ServiceSchemas,
	)
}

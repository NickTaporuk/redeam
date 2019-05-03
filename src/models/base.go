package models

type (
	// RedeamModel use to identify model
	RedeamModel interface {
		ModelName() string
	}

	RedeamModels []RedeamModel
)

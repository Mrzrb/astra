package petstore

import (
	"github.com/Mrzrb/astra/tests/petstore"
)

type Cat struct {
	petstore.Pet

	Breed         string `json:"breed"`
	IsIndependent bool   `json:"isIndependent"`
}

type Dog struct {
	petstore.Pet

	Breed     string `json:"breed"`
	IsTrained bool   `json:"isTrained"`
}

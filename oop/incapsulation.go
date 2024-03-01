package oop

import (
	"crypto/sha256"

	"encoding/hex"
)

type IncapsulatedData struct {
	internalData string
}

func (d *IncapsulatedData) GetHash() string {
	h := sha256.New()

	h.Write([]byte(d.internalData))

	bs := h.Sum(nil)

	encoded := hex.EncodeToString(bs)

	return encoded
}

func NewIncapsulatedData(internalData string) *IncapsulatedData {
	return &IncapsulatedData{internalData}
}

var Data *IncapsulatedData = NewIncapsulatedData("hello world")

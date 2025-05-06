package models

type FieldType struct {
	Type           string
	AllowedSymbols string
}

func NewFieldType() *FieldType {
	return &FieldType{}
}
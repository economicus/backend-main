package model

type Object interface {
	GetID() uint
	GetOwnerID() uint
	ToMap() map[string]interface{}
	ToMapWithFields([]string) map[string]interface{}
}

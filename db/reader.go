package db

type Reader interface {
	ToModel() any
}

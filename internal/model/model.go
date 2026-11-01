package model

type FormData[T any] struct {
	Judul  string
	Action string
	Data   T
	Errors map[string]string
}

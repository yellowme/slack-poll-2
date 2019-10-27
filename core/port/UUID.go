package port

type UUID interface {
	Generate() string
}

package collection_model

//go:generate go run github.com/objectbox/objectbox-go/cmd/objectbox-gogen

type Collection struct {
	Id       uint64
	Name     string
	Genre    string
	ImageUrl string
	Progress float64
}

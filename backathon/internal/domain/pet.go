package domain

type PetStatus string

const (
	PetStatusAvailable PetStatus = "available"
	PetStatusPending   PetStatus = "pending"
	PetStatusSold      PetStatus = "sold"
)

type Pet struct {
	Id        int64
	Status    PetStatus
	Tags      []Tag
	Name      string
	PhotoUrls []string
	Category  Category
}

type Tag struct {
	Id   int64
	Name string
}

type Category struct {
	Id   int64
	Name string
}

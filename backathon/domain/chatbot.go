package domain

import "github.com/google/uuid"

type Hub struct {
	clients map[uuid.UUID]struct{}
}

const (
	HubSize = 10
)

func NewHub() Hub {
	return Hub{
		clients: make(map[uuid.UUID]struct{}, HubSize),
	}
}

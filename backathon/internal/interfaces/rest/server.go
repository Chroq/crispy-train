package rest

type ServerHandler struct {
	*PetStoreHandler
}

// NewServerHandler creates a new ServerHandler.
func NewServerHandler(
	petStoreHandler *PetStoreHandler,
) *ServerHandler {
	return &ServerHandler{
		PetStoreHandler: petStoreHandler,
	}
}

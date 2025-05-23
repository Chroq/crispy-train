// Package rest provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.4.1 DO NOT EDIT.
package rest

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/oapi-codegen/runtime"
)

const (
	Api_keyScopes       = "api_key.Scopes"
	Petstore_authScopes = "petstore_auth.Scopes"
)

// Defines values for OrderStatus.
const (
	Approved  OrderStatus = "approved"
	Delivered OrderStatus = "delivered"
	Placed    OrderStatus = "placed"
)

// Defines values for PetStatus.
const (
	PetStatusAvailable PetStatus = "available"
	PetStatusPending   PetStatus = "pending"
	PetStatusSold      PetStatus = "sold"
)

// Defines values for FindPetsByStatusParamsStatus.
const (
	FindPetsByStatusParamsStatusAvailable FindPetsByStatusParamsStatus = "available"
	FindPetsByStatusParamsStatusPending   FindPetsByStatusParamsStatus = "pending"
	FindPetsByStatusParamsStatusSold      FindPetsByStatusParamsStatus = "sold"
)

// ApiResponse defines model for ApiResponse.
type ApiResponse struct {
	Code    *int32  `json:"code,omitempty"`
	Message *string `json:"message,omitempty"`
	Type    *string `json:"type,omitempty"`
}

// Category defines model for Category.
type Category struct {
	Id   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// Order defines model for Order.
type Order struct {
	Complete *bool      `json:"complete,omitempty"`
	Id       *int64     `json:"id,omitempty"`
	PetId    *int64     `json:"petId,omitempty"`
	Quantity *int32     `json:"quantity,omitempty"`
	ShipDate *time.Time `json:"shipDate,omitempty"`

	// Status Order Status
	Status *OrderStatus `json:"status,omitempty"`
}

// OrderStatus Order Status
type OrderStatus string

// Pet defines model for Pet.
type Pet struct {
	Category  *Category `json:"category,omitempty"`
	Id        *int64    `json:"id,omitempty"`
	Name      string    `json:"name"`
	PhotoUrls []string  `json:"photoUrls"`

	// Status pet status in the store
	Status *PetStatus `json:"status,omitempty"`
	Tags   *[]Tag     `json:"tags,omitempty"`
}

// PetStatus pet status in the store
type PetStatus string

// Tag defines model for Tag.
type Tag struct {
	Id   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// User defines model for User.
type User struct {
	Email     *string `json:"email,omitempty"`
	FirstName *string `json:"firstName,omitempty"`
	Id        *int64  `json:"id,omitempty"`
	LastName  *string `json:"lastName,omitempty"`
	Password  *string `json:"password,omitempty"`
	Phone     *string `json:"phone,omitempty"`

	// UserStatus User Status
	UserStatus *int32  `json:"userStatus,omitempty"`
	Username   *string `json:"username,omitempty"`
}

// FindPetsByStatusParams defines parameters for FindPetsByStatus.
type FindPetsByStatusParams struct {
	// Status Status values that need to be considered for filter
	Status *FindPetsByStatusParamsStatus `form:"status,omitempty" json:"status,omitempty"`
}

// FindPetsByStatusParamsStatus defines parameters for FindPetsByStatus.
type FindPetsByStatusParamsStatus string

// FindPetsByTagsParams defines parameters for FindPetsByTags.
type FindPetsByTagsParams struct {
	// Tags Tags to filter by
	Tags *[]string `form:"tags,omitempty" json:"tags,omitempty"`
}

// DeletePetParams defines parameters for DeletePet.
type DeletePetParams struct {
	ApiKey *string `json:"api_key,omitempty"`
}

// UpdatePetWithFormParams defines parameters for UpdatePetWithForm.
type UpdatePetWithFormParams struct {
	// Name Name of pet that needs to be updated
	Name *string `form:"name,omitempty" json:"name,omitempty"`

	// Status Status of pet that needs to be updated
	Status *string `form:"status,omitempty" json:"status,omitempty"`
}

// UploadFileParams defines parameters for UploadFile.
type UploadFileParams struct {
	// AdditionalMetadata Additional Metadata
	AdditionalMetadata *string `form:"additionalMetadata,omitempty" json:"additionalMetadata,omitempty"`
}

// CreateUsersWithListInputJSONBody defines parameters for CreateUsersWithListInput.
type CreateUsersWithListInputJSONBody = []User

// LoginUserParams defines parameters for LoginUser.
type LoginUserParams struct {
	// Username The user name for login
	Username *string `form:"username,omitempty" json:"username,omitempty"`

	// Password The password for login in clear text
	Password *string `form:"password,omitempty" json:"password,omitempty"`
}

// AddPetJSONRequestBody defines body for AddPet for application/json ContentType.
type AddPetJSONRequestBody = Pet

// AddPetFormdataRequestBody defines body for AddPet for application/x-www-form-urlencoded ContentType.
type AddPetFormdataRequestBody = Pet

// UpdatePetJSONRequestBody defines body for UpdatePet for application/json ContentType.
type UpdatePetJSONRequestBody = Pet

// UpdatePetFormdataRequestBody defines body for UpdatePet for application/x-www-form-urlencoded ContentType.
type UpdatePetFormdataRequestBody = Pet

// PlaceOrderJSONRequestBody defines body for PlaceOrder for application/json ContentType.
type PlaceOrderJSONRequestBody = Order

// PlaceOrderFormdataRequestBody defines body for PlaceOrder for application/x-www-form-urlencoded ContentType.
type PlaceOrderFormdataRequestBody = Order

// CreateUserJSONRequestBody defines body for CreateUser for application/json ContentType.
type CreateUserJSONRequestBody = User

// CreateUserFormdataRequestBody defines body for CreateUser for application/x-www-form-urlencoded ContentType.
type CreateUserFormdataRequestBody = User

// CreateUsersWithListInputJSONRequestBody defines body for CreateUsersWithListInput for application/json ContentType.
type CreateUsersWithListInputJSONRequestBody = CreateUsersWithListInputJSONBody

// UpdateUserJSONRequestBody defines body for UpdateUser for application/json ContentType.
type UpdateUserJSONRequestBody = User

// UpdateUserFormdataRequestBody defines body for UpdateUser for application/x-www-form-urlencoded ContentType.
type UpdateUserFormdataRequestBody = User

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Add a new pet to the store.
	// (POST /pet)
	AddPet(c *gin.Context)
	// Update an existing pet.
	// (PUT /pet)
	UpdatePet(c *gin.Context)
	// Finds Pets by status.
	// (GET /pet/findByStatus)
	FindPetsByStatus(c *gin.Context, params FindPetsByStatusParams)
	// Finds Pets by tags.
	// (GET /pet/findByTags)
	FindPetsByTags(c *gin.Context, params FindPetsByTagsParams)
	// Deletes a pet.
	// (DELETE /pet/{petId})
	DeletePet(c *gin.Context, petId int64, params DeletePetParams)
	// Find pet by ID.
	// (GET /pet/{petId})
	GetPetById(c *gin.Context, petId int64)
	// Updates a pet in the store with form data.
	// (POST /pet/{petId})
	UpdatePetWithForm(c *gin.Context, petId int64, params UpdatePetWithFormParams)
	// Uploads an image.
	// (POST /pet/{petId}/uploadImage)
	UploadFile(c *gin.Context, petId int64, params UploadFileParams)
	// Returns pet inventories by status.
	// (GET /store/inventory)
	GetInventory(c *gin.Context)
	// Place an order for a pet.
	// (POST /store/order)
	PlaceOrder(c *gin.Context)
	// Delete purchase order by identifier.
	// (DELETE /store/order/{orderId})
	DeleteOrder(c *gin.Context, orderId int64)
	// Find purchase order by ID.
	// (GET /store/order/{orderId})
	GetOrderById(c *gin.Context, orderId int64)
	// Create user.
	// (POST /user)
	CreateUser(c *gin.Context)
	// Creates list of users with given input array.
	// (POST /user/createWithList)
	CreateUsersWithListInput(c *gin.Context)
	// Logs user into the system.
	// (GET /user/login)
	LoginUser(c *gin.Context, params LoginUserParams)
	// Logs out current logged in user session.
	// (GET /user/logout)
	LogoutUser(c *gin.Context)
	// Delete user resource.
	// (DELETE /user/{username})
	DeleteUser(c *gin.Context, username string)
	// Get user by user name.
	// (GET /user/{username})
	GetUserByName(c *gin.Context, username string)
	// Update user resource.
	// (PUT /user/{username})
	UpdateUser(c *gin.Context, username string)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandler       func(*gin.Context, error, int)
}

type MiddlewareFunc func(c *gin.Context)

// AddPet operation middleware
func (siw *ServerInterfaceWrapper) AddPet(c *gin.Context) {

	c.Set(Petstore_authScopes, []string{"write:pets", "read:pets"})

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.AddPet(c)
}

// UpdatePet operation middleware
func (siw *ServerInterfaceWrapper) UpdatePet(c *gin.Context) {

	c.Set(Petstore_authScopes, []string{"write:pets", "read:pets"})

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.UpdatePet(c)
}

// FindPetsByStatus operation middleware
func (siw *ServerInterfaceWrapper) FindPetsByStatus(c *gin.Context) {

	var err error

	c.Set(Petstore_authScopes, []string{"write:pets", "read:pets"})

	// Parameter object where we will unmarshal all parameters from the context
	var params FindPetsByStatusParams

	// ------------- Optional query parameter "status" -------------

	err = runtime.BindQueryParameter("form", true, false, "status", c.Request.URL.Query(), &params.Status)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter status: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.FindPetsByStatus(c, params)
}

// FindPetsByTags operation middleware
func (siw *ServerInterfaceWrapper) FindPetsByTags(c *gin.Context) {

	var err error

	c.Set(Petstore_authScopes, []string{"write:pets", "read:pets"})

	// Parameter object where we will unmarshal all parameters from the context
	var params FindPetsByTagsParams

	// ------------- Optional query parameter "tags" -------------

	err = runtime.BindQueryParameter("form", true, false, "tags", c.Request.URL.Query(), &params.Tags)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter tags: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.FindPetsByTags(c, params)
}

// DeletePet operation middleware
func (siw *ServerInterfaceWrapper) DeletePet(c *gin.Context) {

	var err error

	// ------------- Path parameter "petId" -------------
	var petId int64

	err = runtime.BindStyledParameterWithOptions("simple", "petId", c.Param("petId"), &petId, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter petId: %w", err), http.StatusBadRequest)
		return
	}

	c.Set(Petstore_authScopes, []string{"write:pets", "read:pets"})

	// Parameter object where we will unmarshal all parameters from the context
	var params DeletePetParams

	headers := c.Request.Header

	// ------------- Optional header parameter "api_key" -------------
	if valueList, found := headers[http.CanonicalHeaderKey("api_key")]; found {
		var ApiKey string
		n := len(valueList)
		if n != 1 {
			siw.ErrorHandler(c, fmt.Errorf("Expected one value for api_key, got %d", n), http.StatusBadRequest)
			return
		}

		err = runtime.BindStyledParameterWithOptions("simple", "api_key", valueList[0], &ApiKey, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false})
		if err != nil {
			siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter api_key: %w", err), http.StatusBadRequest)
			return
		}

		params.ApiKey = &ApiKey

	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.DeletePet(c, petId, params)
}

// GetPetById operation middleware
func (siw *ServerInterfaceWrapper) GetPetById(c *gin.Context) {

	var err error

	// ------------- Path parameter "petId" -------------
	var petId int64

	err = runtime.BindStyledParameterWithOptions("simple", "petId", c.Param("petId"), &petId, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter petId: %w", err), http.StatusBadRequest)
		return
	}

	c.Set(Api_keyScopes, []string{})

	c.Set(Petstore_authScopes, []string{"write:pets", "read:pets"})

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetPetById(c, petId)
}

// UpdatePetWithForm operation middleware
func (siw *ServerInterfaceWrapper) UpdatePetWithForm(c *gin.Context) {

	var err error

	// ------------- Path parameter "petId" -------------
	var petId int64

	err = runtime.BindStyledParameterWithOptions("simple", "petId", c.Param("petId"), &petId, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter petId: %w", err), http.StatusBadRequest)
		return
	}

	c.Set(Petstore_authScopes, []string{"write:pets", "read:pets"})

	// Parameter object where we will unmarshal all parameters from the context
	var params UpdatePetWithFormParams

	// ------------- Optional query parameter "name" -------------

	err = runtime.BindQueryParameter("form", true, false, "name", c.Request.URL.Query(), &params.Name)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter name: %w", err), http.StatusBadRequest)
		return
	}

	// ------------- Optional query parameter "status" -------------

	err = runtime.BindQueryParameter("form", true, false, "status", c.Request.URL.Query(), &params.Status)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter status: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.UpdatePetWithForm(c, petId, params)
}

// UploadFile operation middleware
func (siw *ServerInterfaceWrapper) UploadFile(c *gin.Context) {

	var err error

	// ------------- Path parameter "petId" -------------
	var petId int64

	err = runtime.BindStyledParameterWithOptions("simple", "petId", c.Param("petId"), &petId, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter petId: %w", err), http.StatusBadRequest)
		return
	}

	c.Set(Petstore_authScopes, []string{"write:pets", "read:pets"})

	// Parameter object where we will unmarshal all parameters from the context
	var params UploadFileParams

	// ------------- Optional query parameter "additionalMetadata" -------------

	err = runtime.BindQueryParameter("form", true, false, "additionalMetadata", c.Request.URL.Query(), &params.AdditionalMetadata)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter additionalMetadata: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.UploadFile(c, petId, params)
}

// GetInventory operation middleware
func (siw *ServerInterfaceWrapper) GetInventory(c *gin.Context) {

	c.Set(Api_keyScopes, []string{})

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetInventory(c)
}

// PlaceOrder operation middleware
func (siw *ServerInterfaceWrapper) PlaceOrder(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.PlaceOrder(c)
}

// DeleteOrder operation middleware
func (siw *ServerInterfaceWrapper) DeleteOrder(c *gin.Context) {

	var err error

	// ------------- Path parameter "orderId" -------------
	var orderId int64

	err = runtime.BindStyledParameterWithOptions("simple", "orderId", c.Param("orderId"), &orderId, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter orderId: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.DeleteOrder(c, orderId)
}

// GetOrderById operation middleware
func (siw *ServerInterfaceWrapper) GetOrderById(c *gin.Context) {

	var err error

	// ------------- Path parameter "orderId" -------------
	var orderId int64

	err = runtime.BindStyledParameterWithOptions("simple", "orderId", c.Param("orderId"), &orderId, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter orderId: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetOrderById(c, orderId)
}

// CreateUser operation middleware
func (siw *ServerInterfaceWrapper) CreateUser(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.CreateUser(c)
}

// CreateUsersWithListInput operation middleware
func (siw *ServerInterfaceWrapper) CreateUsersWithListInput(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.CreateUsersWithListInput(c)
}

// LoginUser operation middleware
func (siw *ServerInterfaceWrapper) LoginUser(c *gin.Context) {

	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params LoginUserParams

	// ------------- Optional query parameter "username" -------------

	err = runtime.BindQueryParameter("form", true, false, "username", c.Request.URL.Query(), &params.Username)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter username: %w", err), http.StatusBadRequest)
		return
	}

	// ------------- Optional query parameter "password" -------------

	err = runtime.BindQueryParameter("form", true, false, "password", c.Request.URL.Query(), &params.Password)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter password: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.LoginUser(c, params)
}

// LogoutUser operation middleware
func (siw *ServerInterfaceWrapper) LogoutUser(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.LogoutUser(c)
}

// DeleteUser operation middleware
func (siw *ServerInterfaceWrapper) DeleteUser(c *gin.Context) {

	var err error

	// ------------- Path parameter "username" -------------
	var username string

	err = runtime.BindStyledParameterWithOptions("simple", "username", c.Param("username"), &username, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter username: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.DeleteUser(c, username)
}

// GetUserByName operation middleware
func (siw *ServerInterfaceWrapper) GetUserByName(c *gin.Context) {

	var err error

	// ------------- Path parameter "username" -------------
	var username string

	err = runtime.BindStyledParameterWithOptions("simple", "username", c.Param("username"), &username, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter username: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetUserByName(c, username)
}

// UpdateUser operation middleware
func (siw *ServerInterfaceWrapper) UpdateUser(c *gin.Context) {

	var err error

	// ------------- Path parameter "username" -------------
	var username string

	err = runtime.BindStyledParameterWithOptions("simple", "username", c.Param("username"), &username, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter username: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.UpdateUser(c, username)
}

// GinServerOptions provides options for the Gin server.
type GinServerOptions struct {
	BaseURL      string
	Middlewares  []MiddlewareFunc
	ErrorHandler func(*gin.Context, error, int)
}

// RegisterHandlers creates http.Handler with routing matching OpenAPI spec.
func RegisterHandlers(router gin.IRouter, si ServerInterface) {
	RegisterHandlersWithOptions(router, si, GinServerOptions{})
}

// RegisterHandlersWithOptions creates http.Handler with additional options
func RegisterHandlersWithOptions(router gin.IRouter, si ServerInterface, options GinServerOptions) {
	errorHandler := options.ErrorHandler
	if errorHandler == nil {
		errorHandler = func(c *gin.Context, err error, statusCode int) {
			c.JSON(statusCode, gin.H{"msg": err.Error()})
		}
	}

	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandler:       errorHandler,
	}

	router.POST(options.BaseURL+"/pet", wrapper.AddPet)
	router.PUT(options.BaseURL+"/pet", wrapper.UpdatePet)
	router.GET(options.BaseURL+"/pet/findByStatus", wrapper.FindPetsByStatus)
	router.GET(options.BaseURL+"/pet/findByTags", wrapper.FindPetsByTags)
	router.DELETE(options.BaseURL+"/pet/:petId", wrapper.DeletePet)
	router.GET(options.BaseURL+"/pet/:petId", wrapper.GetPetById)
	router.POST(options.BaseURL+"/pet/:petId", wrapper.UpdatePetWithForm)
	router.POST(options.BaseURL+"/pet/:petId/uploadImage", wrapper.UploadFile)
	router.GET(options.BaseURL+"/store/inventory", wrapper.GetInventory)
	router.POST(options.BaseURL+"/store/order", wrapper.PlaceOrder)
	router.DELETE(options.BaseURL+"/store/order/:orderId", wrapper.DeleteOrder)
	router.GET(options.BaseURL+"/store/order/:orderId", wrapper.GetOrderById)
	router.POST(options.BaseURL+"/user", wrapper.CreateUser)
	router.POST(options.BaseURL+"/user/createWithList", wrapper.CreateUsersWithListInput)
	router.GET(options.BaseURL+"/user/login", wrapper.LoginUser)
	router.GET(options.BaseURL+"/user/logout", wrapper.LogoutUser)
	router.DELETE(options.BaseURL+"/user/:username", wrapper.DeleteUser)
	router.GET(options.BaseURL+"/user/:username", wrapper.GetUserByName)
	router.PUT(options.BaseURL+"/user/:username", wrapper.UpdateUser)
}

package main

import (
	"strings"
	"github.com/dghubble/sling"
	"net/http"
	"time"
)

type Operation struct {
	Data *OperationResource `json:"data"`
}

type Operations struct {
	Data []OperationResource `json:"data"`
}

type OperationResource struct {
	Resource
	Attributes *OperationAttributes `json:"attributes"`
}

type OperationAttributes struct {
	Started time.Time `json:"started"`
	Map string `json:"map"`
}

type OperationService struct {
	sling *sling.Sling
}

func newOperationService(sling *sling.Sling) *OperationService {
	return &OperationService {
		sling: sling.Path("operations/"),
	}
}

type OperationListParams struct { }

func (o *OperationService) List(ids []string, params *OperationListParams) (*Operations, *http.Response, error) {
	operations := new(Operations)
	serviceError := new(ServiceError)

	resp, err := o.sling.New().Get(strings.Join(ids, ",")).Receive(operations, serviceError)
	return operations, resp, err
}

func (o *OperationService) ListAll(params *OperationListParams) (*Operations, *http.Response, error) {
	operations := new(Operations)
	serviceError := new(ServiceError)

	resp, err := o.sling.New().Get("").Receive(operations, serviceError)
	return operations, resp, err
}

func (o *OperationService) Create(attributes *OperationAttributes) (*Operation, *http.Response, error) {
	operation := new(Operation)
	serviceError := new(ServiceError)

	resp, err := o.sling.New().Post("").BodyJSON(Operation {
		Data: &OperationResource {
			Attributes: attributes,
		},
	}).Receive(operation, serviceError)

	return operation, resp, prioritize(err, *serviceError)
}

func (o *OperationService) Read(id string) (*Operation, *http.Response, error) {
	operation := new(Operation)
	serviceError := new(ServiceError)

	resp, err := o.sling.New().Get(id).Receive(operation, serviceError)
	return operation, resp, err
}

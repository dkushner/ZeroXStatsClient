package main

import "fmt"

type ServiceError struct {
	Errors []ServiceErrorResource `json:"errors"`
}

type ServiceErrorResource struct {
	Id string `json:"id"`
	Links *ServiceErrorLinks `json:"links"`
	Status string `json:"status"`
	Code string `json:"code"`
	Title string `json:"title"`
	Detail string `json:"detail"`
	Source *ServiceErrorSource `json:"source"`
	Meta interface{} `json:"meta"`
}

type ServiceErrorLinks struct {
	About string `json:"about"`
}

type ServiceErrorSource struct {
	Pointer string `json:"pointer"`
	Parameter string `json:"parameter"`
}

func (e ServiceError) Error() string {
	return fmt.Sprintf("%v: %v", e.Errors[0].Code, e.Errors[0].Detail)
}

func prioritize(httpError error, serviceError ServiceError) error {
	if httpError != nil  {
		return httpError
	}

	return serviceError
}

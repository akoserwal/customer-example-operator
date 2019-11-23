package controller

import (
	"github.com/customer-op/customer-example-operator/pkg/controller/customer"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, customer.Add)
}

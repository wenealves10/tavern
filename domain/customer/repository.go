package customer

import (
	"errors"

	"github.com/google/uuid"
	"github.com/wenealves10/microservice-golang-k8s-ci-cd/aggregate"
)

var (
	ErrCustomerNotFound    = errors.New("the customer was not found in the repository")
	ErrFailedToAddCustomer = errors.New("failed to add customer to the repository")
	ErrUpdateCustomer      = errors.New("failed to update customer in the repository")
)

type CustomerRepository interface {
	Get(uuid.UUID) (aggregate.Customer, error)
	Add(aggregate.Customer) error
	Update(aggregate.Customer) error
}

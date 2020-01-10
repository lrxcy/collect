package main

import "fmt"

type Customer interface {
	Accept(Visitor)
}

type Visitor interface {
	Visit(Customer)
}

type CustomerCol struct {
	customers []Customer
}

func (c *CustomerCol) Add(customer Customer) {
	c.customers = append(c.customers, customer)
}

func (c *CustomerCol) Accept(visitor Visitor) {
	for _, customer := range c.customers {
		customer.Accept(visitor)
	}
}

type EnterpriseCustomer struct {
	name string
}

func (c *EnterpriseCustomer) Accept(visitor Visitor) {
	visitor.Visit(c)
}

func NewEnterpriseCustomer(name string) *EnterpriseCustomer {
	return &EnterpriseCustomer{
		name: name,
	}
}

type IndividualCustomer struct {
	name string
}

func (c *IndividualCustomer) Accept(visitor Visitor) {
	visitor.Visit(c)
}

func NewIndividualCustomer(name string) *IndividualCustomer {
	return &IndividualCustomer{
		name: name,
	}
}

type ServiceRequestVisitor struct{}

func (*ServiceRequestVisitor) Visit(customer Customer) {
	switch c := customer.(type) {
	case *EnterpriseCustomer:
		fmt.Printf("serving enterprise customer %s\n", c.name)
	case *IndividualCustomer:
		fmt.Printf("serving individual customer %s\n", c.name)
	}
}

// only for enterprise
type AnalysisVisitor struct{}

func (*AnalysisVisitor) Visit(customer Customer) {
	switch c := customer.(type) {
	case *EnterpriseCustomer:
		fmt.Printf("analysis enterprise customer %s\n", c.name)
	case *IndividualCustomer:
		fmt.Printf("%s is not an enterprise customer, no analysis would be applied\n", c.name)
	}
}

func main() {
	c1 := &CustomerCol{}
	c1.Add(NewEnterpriseCustomer("A company"))
	c1.Add(NewEnterpriseCustomer("B company"))
	c1.Add(NewIndividualCustomer("bob"))
	c1.Accept(&ServiceRequestVisitor{})
	// Output:
	// serving enterprise customer A company
	// serving enterprise customer B company
	// serving individual customer bob

	c2 := &CustomerCol{}
	c2.Add(NewEnterpriseCustomer("A company"))
	c2.Add(NewIndividualCustomer("bob"))
	c2.Add(NewEnterpriseCustomer("B company"))
	c2.Accept(&AnalysisVisitor{})
	// Output:
	// analysis enterprise customer A company
	// bob is not an enterprise customer, no analysis would be applied
	// analysis enterprise customer B company
}

package main

import (
	"fmt"
	service "test-project/ser"
)

type Producer interface {
	Produce() ([]string, error)
}

type Presenter interface {
	Present(data []string)
}

func main() {
	var p Producer = producer{}

	var pr Presenter = presenter{}

	dataservice := service.NewDataService(p, pr)
	dataservice.Process()

}

type producer struct{}
type presenter struct{}

func (producer) Produce() ([]string, error) {
	return []string{"яблоко", "банан", "апельсин"}, nil

}

func (presenter) Present(data []string) {
	for i, item := range data {
		fmt.Printf("%d. %s\n", i+1, item)

	}
}

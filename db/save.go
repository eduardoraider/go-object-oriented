package db

import "fmt"

type Entity interface {
	GetName() string
	Validate() error
}

func Save(e Entity) error {
	if err := e.Validate(); err != nil {
		return err
	}

	fmt.Printf("entity %s saved!\n", e.GetName())

	return nil
}

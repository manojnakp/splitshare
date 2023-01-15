package db

import "github.com/manojnakp/splitshare/pkg/component"

var _ Collection = &ECollection{}

type ECollection []component.Expense

func (c *ECollection) Insert(doc any) error {
	return nil
}

func (c *ECollection) Find(id string) (any, error) {
	return nil, nil
}

func (c *ECollection) Update(id string, doc any) error {
	return nil
}

func (c *ECollection) Delete(id string) error {
	return nil
}

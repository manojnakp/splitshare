package db

import "github.com/manojnakp/splitshare/pkg/component"

var _ Collection = &ECollection{}

type ECollection []*component.Expense

func (c *ECollection) Insert(doc any) error {
	x, ok := doc.(component.Expense)
	if !ok {
		return ErrInvalidType
	}
	*c = append(*c, &x)
	return nil
}

func (c *ECollection) Find(id string) (any, error) {
	for _, v := range *c {
		if v.Id == id {
			return *v, nil
		}
	}
	return nil, ErrNotFound
}

func (c *ECollection) Update(id string, doc any) error {
	x, ok := doc.(component.Expense)
	if !ok {
		return ErrInvalidType
	}
	for i, v := range *c {
		if v.Id == id {
			c.validateUpdate(i, x)
			return nil
		}
	}
	return ErrNotFound
}

func (c *ECollection) Delete(id string) error {
	n := len(*c)
	for i, v := range *c {
		if v.Id == id {
			(*c)[i], (*c)[n-1] = (*c)[n-1], (*c)[i]
			(*c)[n-1] = nil
			*c = (*c)[0 : n-1]
			return nil
		}
	}
	return ErrNotFound
}

func (c *ECollection) validateUpdate(i int, x component.Expense) {
	if x.Name != "" {
		(*c)[i].Name = x.Name
	}
	if x.By != "" {
		(*c)[i].By = x.By
	}
	if x.Among != nil {
		(*c)[i].Among = x.Among
	}
}

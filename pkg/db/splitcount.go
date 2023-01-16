package db

import (
	"errors"
	"github.com/manojnakp/splitshare/pkg/component"
)

var _ Collection = &SCollection{}

var (
	ErrNotFound    = errors.New("record not found")
	ErrInvalidType = errors.New("invalid splitcount")
)

type SCollection []*component.SplitCount

func (c *SCollection) Insert(doc any) error {
	s, ok := doc.(component.SplitCount)
	if !ok {
		return ErrInvalidType
	}
	*c = append(*c, &s)
	return nil
}

func (c *SCollection) Find(id string) (any, error) {
	for _, v := range *c {
		if v.Id == id {
			return *v, nil
		}
	}
	return nil, ErrNotFound
}

func (c *SCollection) Update(id string, doc any) error {
	s, ok := doc.(component.SplitCount)
	if !ok {
		return ErrInvalidType
	}
	for i, v := range *c {
		if v.Id == id {
			c.validateUpdate(i, s)
			return nil
		}
	}
	return ErrNotFound
}

func (c *SCollection) Delete(id string) error {
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

func (c *SCollection) validateUpdate(i int, s component.SplitCount) {
	if s.Title != "" {
		(*c)[i].Title = s.Title
	}
	if s.Desc != "" {
		(*c)[i].Desc = s.Desc
	}
	if s.By != "" {
		(*c)[i].By = s.By
	}
}

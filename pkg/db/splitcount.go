package db

import "github.com/manojnakp/splitshare/pkg/component"

var _ Collection = &SCollection{}

type SCollection []component.SplitCount

func (c *SCollection) Insert(doc any) error {
	return nil
}

func (c *SCollection) Find(id string) (any, error) {
	return nil, nil
}

func (c *SCollection) Update(id string, doc any) error {
	return nil
}

func (c *SCollection) Delete(id string) error {
	return nil
}

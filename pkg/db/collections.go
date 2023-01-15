package db

type Collection interface {
	Insert(doc any) error
	Find(id string) (any, error)
	Update(id string, doc any) error
	Delete(id string) error
}

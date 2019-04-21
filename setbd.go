package setbd

import ()

type DB interface {
	Close() error
	Contain(v string) (bool, error)
	Insert(v string) error
	Erase(v string) error
}

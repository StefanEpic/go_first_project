package db

import (
	"github.com/lib/pq"
	"gorm.io/datatypes"
)

type ListListStrings []pq.StringArray

func (t *ListListStrings) Scan(src interface{}) error {
	return pq.GenericArray{t}.Scan(src)
}

type ListJsonb []datatypes.JSON

func (t *ListJsonb) Scan(src interface{}) error {
	return pq.GenericArray{t}.Scan(src)
}

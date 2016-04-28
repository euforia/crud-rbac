package crudrbac

import (
	"encoding/json"
	"github.com/boltdb/bolt"
)

const (
	RoleBucketName = "crudrbac.role"
)

type RbacDatastore interface {
	GetRole(id string) (Role, error)
	UpdateRole(role Role) error
	CreateRole(role Role) error
	DeleteRole(role Role) error
}

type BoltdbRbacDatastore struct {
	db *bolt.DB
}

func NewBoltdbRbacDatastore(dbfile string) (b *BoltdbRbacDatastore, err error) {
	b = &BoltdbRbacDatastore{}
	b.db, err = bolt.Open(dbfile, 0600, nil)

	return
}

func (brd *BoltdbRbacDatastore) GetRole(id string) (role Role, err error) {
	err = brd.db.View(func(tx *bolt.Tx) error {

		bucket := tx.Bucket([]byte(RoleBucketName))
		//if e == nil {
		d := bucket.Get([]byte(id))

		return json.Unmarshal(d, &role)
		//}

		//return e
	})
	return
}

func (brd *BoltdbRbacDatastore) UpdateRole(role Role) (err error) {
	err = brd.db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(RoleBucketName))
		//if e == nil {
		//var d []byte
		d, e := json.Marshal(role)
		if e == nil {
			e = bucket.Put([]byte(role.Id), d)
		}
		//}
		return e
	})
	return
}

func (brd *BoltdbRbacDatastore) CreateRole(role Role) (err error) {
	err = brd.db.Update(func(tx *bolt.Tx) error {

		bucket, e := tx.CreateBucketIfNotExists([]byte(RoleBucketName))
		if e == nil {
			var d []byte
			if d, e = json.Marshal(role); e == nil {
				e = bucket.Put([]byte(role.Id), d)
			}
		}

		return e
	})
	return
}

func (brd *BoltdbRbacDatastore) DeleteRole(id string) (err error) {
	err = brd.db.Update(func(tx *bolt.Tx) error {

		bucket := tx.Bucket([]byte(RoleBucketName))
		//if e == nil {
		return bucket.Delete([]byte(id))
		//}

		//return e

	})
	return
}

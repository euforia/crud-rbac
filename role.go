package crudrbac

import (
	"time"
)

type Role struct {
	Id string

	// Only fields that are updateable
	Name     string
	Policies []Policy

	// nanoseconds
	CreatedDate int64
	UpdatedDate int64
	Version     int64
}

func NewRole() *Role {
	return &Role{
		CreatedDate: time.Now().UnixNano(),
		Version:     1,
	}
}

func (r *Role) IsGranted(policy Policy) *Policy {
	for _, p := range r.Policies {
		plcy := p.Match(policy)
		if plcy != nil {
			return plcy
		}
	}
	return nil
}

// Reset values.
func (r *Role) Reset() {
	r.CreatedDate = time.Now().UnixNano()
	r.UpdatedDate = r.CreatedDate
	r.Version = 1
}

func (r *Role) Update(role Role) {

	r.Name = role.Name
	r.Policies = role.Policies
	r.Version += 1
	r.UpdatedDate = time.Now().UnixNano()

}

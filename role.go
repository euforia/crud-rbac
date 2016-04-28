package crudrbac

type Role struct {
	Id       string   `json:"id"`
	Name     string   `json:"name"`
	Policies []Policy `json:"policies"`
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

package crudrbac

import (
	"net/http"
	"strings"
)

type OpType string

const (
	OpCreate OpType = "create"
	OpRead   OpType = "read"
	OpUpdate OpType = "update"
	OpDelete OpType = "delete"
	OpAll    OpType = "*"
)

type Policy struct {
	Op       OpType
	Resource string

	// Optional rules on data.  Each key is an 'AND'
	// while a key with more than 1 element will be treated as an 'OR'
	// on the same key.
	Constraints map[string]interface{}

	Allow bool
}

// Match on Op and resource widlcard
// Return the matched policy
func (p *Policy) Match(policy Policy) *Policy {
	// Exact or all op
	// *resource, resource*, *resource*
	if (p.Op == policy.Op || p.Op == OpAll) &&
		(p.Resource == policy.Resource || p.Resource == "*" ||
			(p.Resource[0] == '*' && strings.HasSuffix(policy.Resource, p.Resource[1:])) ||
			(p.Resource[len(p.Resource)-1] == '*' && strings.HasPrefix(policy.Resource, p.Resource[:len(p.Resource)-1])) ||
			(p.Resource[0] == '*' && p.Resource[len(p.Resource)-1] == '*' && strings.Contains(policy.Resource, p.Resource[1:len(p.Resource)-1]))) {

		return p
	}
	return nil
}

func ParseHttpRequestPolicy(r *http.Request, basePath string) (p Policy) {

	p.Resource = strings.TrimPrefix(r.RequestURI, basePath)

	switch r.Method {
	case "GET":
		p.Op = OpRead
	case "PUT", "PATCH":
		p.Op = OpUpdate
	case "POST":
		p.Op = OpCreate
	case "DELETE":
		p.Op = OpDelete
	}

	return
}

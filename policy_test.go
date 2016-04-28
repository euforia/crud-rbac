package crudrbac

import (
	//"encoding/json"
	"testing"
)

var (
	testP1 = Policy{Op: OpAll, Resource: "/Asset/*"}
	testP2 = Policy{Op: OpAll, Resource: "*Asset*"}
)

func Test_Policy_End_wildcard(t *testing.T) {

	if testP1.Op != OpAll {
		t.Fatal("Op mismatch!")
	}

	tp := Policy{Op: OpCreate, Resource: "/Asset/12345"}

	if testP1.Match(tp) == nil {
		t.Fatal("should match")
	}

	tp.Resource = "/Company"
	if testP1.Match(tp) != nil {
		t.Fatal("should not match")
	}
}

func Test_Policy_Dual_wildcard(t *testing.T) {
	tp := Policy{Op: OpCreate, Resource: "/Asset/12ws"}
	if testP2.Match(tp) == nil {
		t.Fatal("Should match")
	}

	tp.Resource = "/Company/foo"
	if testP2.Match(tp) != nil {
		t.Fatal("Should not match")
	}
}

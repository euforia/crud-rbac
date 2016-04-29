package crudrbac

import (
	"encoding/json"
	"io/ioutil"
	"testing"
)

var (
	testDbfile   = "etc/test.bdb"
	testRoleFile = "etc/roles/admin.json"
	testDs       *BoltdbRbacDatastore
	testRole     Role
)

func Test_BoltdbRbacDatastore(t *testing.T) {
	var err error
	if testDs, err = NewBoltdbRbacDatastore(testDbfile); err != nil {
		t.Fatal(err)
	}

	var b []byte
	if b, err = ioutil.ReadFile(testRoleFile); err == nil {
		json.Unmarshal(b, &testRole)
		//t.Logf("%#v\n", testRole)
		return
	}
	t.Fatal(err)
}

func Test_BoltdbRbacDatastore_CreateRole(t *testing.T) {
	if err := testDs.CreateRole(testRole); err != nil {
		t.Fatal(err)
	}

	r, err := testDs.GetRole("admin")
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Created: %#v\n", r)
}

func Test_BoltdbRbacDatastore_DeleteRole(t *testing.T) {
	if err := testDs.DeleteRole("admin"); err != nil {
		t.Fatal(err)
	}

	_, err := testDs.GetRole("admin")
	if err == nil {
		t.Fatal("Should not exist")
	}
}

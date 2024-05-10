package main

import (
	"fmt"

	"github.com/casbin/casbin/v2"
	mongodbadapter "github.com/casbin/mongodb-adapter/v3"
)

func main() {
	// Initialize a MongoDB adapter and use it in a Casbin enforcer:
	// a. Use the NewAdapter() function to create an adapter.
	// b. Use the NewEnforcer() function to create an enforcer.
	// c. Use the SetAdapter() function to set the adapter.
	// d. Use the LoadPolicy() function to load the policy from the adapter.
	// e. Use the Enforce() function to enforce the policy.
	adapter, err := mongodbadapter.NewAdapter("mongodb://localhost:27017/data-collection-hub")
	if err != nil {
		panic(err)
	}
	e, err := casbin.NewEnforcer("configs/dev/casbin_model.dev.conf", adapter)
	if err != nil {
		panic(err)
	}
	err = e.LoadPolicy()
	if err != nil {
		panic(err)
	}
	enforce, err := e.Enforce("alice1", "data1", "read")
	if err != nil {
		panic(err)
	}
	if enforce {
		fmt.Println("Permission granted")
	} else {
		fmt.Println("Permission denied")
	}
	policy, err := e.AddPolicy("alice", "data2", "write")
	if err != nil {
		panic(err)
	} else if !policy {
		fmt.Println("Policy already exists")
	} else {
		fmt.Println("Policy added")
	}
	err = e.SavePolicy()
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Policy saved")
	}

}

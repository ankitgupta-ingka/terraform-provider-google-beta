// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Nathan is editing this to generate diffs in lots of files.
//     He won't submit this change.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------
package google

import (
	"fmt"
)

type AccessContextManagerOperationWaiter struct {
	Config *Config
	CommonOperationWaiter
}

func (w *AccessContextManagerOperationWaiter) QueryOp() (interface{}, error) {
	if w == nil {
		return nil, fmt.Errorf("Cannot query operation, it's unset or nil.")
	}
	// Returns the proper get.
	url := fmt.Sprintf("https://accesscontextmanager.googleapis.com/v1/%s", w.CommonOperationWaiter.Op.Name)
	return sendRequest(w.Config, "GET", "", url, nil)
}

func accessContextManagerOperationWaitTime(config *Config, op map[string]interface{}, activity string, timeoutMinutes int) error {
	if val, ok := op["name"]; !ok || val == "" {
		// This was a synchronous call - there is no operation to wait for.
		return nil
	}
	w := &AccessContextManagerOperationWaiter{
		Config: config,
	}
	if err := w.CommonOperationWaiter.SetOp(op); err != nil {
		return err
	}
	return OperationWait(w, activity, timeoutMinutes)
}

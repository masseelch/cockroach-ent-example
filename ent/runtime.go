// Code generated by entc, DO NOT EDIT.

package ent

import (
	"github.com/yonidavidson/cockroachent/ent/account"
	"github.com/yonidavidson/cockroachent/ent/schema"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	accountFields := schema.Account{}.Fields()
	_ = accountFields
	// accountDescBalance is the schema descriptor for balance field.
	accountDescBalance := accountFields[0].Descriptor()
	// account.BalanceValidator is a validator for the "balance" field. It is called by the builders before save.
	account.BalanceValidator = accountDescBalance.Validators[0].(func(int) error)
}

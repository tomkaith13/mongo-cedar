package cedar_policy

import (
	"os"

	"github.com/cedar-policy/cedar-go"
)

var PolicySet *cedar.PolicySet

const (
	policyFile = "./policy.cedar"
)

func LoadPolicy() error {
	b, err := os.ReadFile(policyFile)
	if err != nil {
		return err
	}

	PolicySet, err = cedar.NewPolicySetFromBytes(policyFile, b)
	if err != nil {
		return err
	}

	return nil
}

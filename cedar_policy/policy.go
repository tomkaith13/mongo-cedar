package cedar_policy

import (
	"os"

	"github.com/cedar-policy/cedar-go"
)

var Policy cedar.Policy
var PolicySet *cedar.PolicySet

const (
	policyFile = "./policy.cedar"
)

func LoadPolicy() error {
	var policy cedar.Policy

	b, err := os.ReadFile(policyFile)
	if err != nil {
		return err
	}

	err = policy.UnmarshalCedar(b)
	if err != nil {
		return err
	}

	Policy = policy
	PolicySet = cedar.NewPolicySet()
	PolicySet.Add("cg policy", &Policy)

	return nil
}

package cedar_policy

import (
	"os"

	"github.com/cedar-policy/cedar-go"
)

var Policy cedar.Policy

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

	return nil
}

package iam_test

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/dtcookie/dynatrace/api/accounts/iam"
)

func TestPolicyServiceClient(t *testing.T) {
	clientID := os.Getenv("DT_CLIENT_ID")
	accountID := os.Getenv("DT_ACCOUNT_ID")
	clientSecret := os.Getenv("DT_CLIENT_SECRET")
	if len(clientID) == 0 || len(accountID) == 0 || len(clientSecret) == 0 {
		t.Skip("DT_CLIENT_ID, DT_ACCOUNT_ID and DT_CLIENT_SECRET required as environment variables - skipping test")
	}

	policyService := iam.NewBasePolicyService(clientID, accountID, clientSecret)

	policyUUIDs, err := policyService.LIST(iam.PolicyLevels.Account, accountID)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Printf("%d policies found\n", len(policyUUIDs))
	for _, policyUUID := range policyUUIDs {
		var policy *iam.Policy
		fmt.Println("GET", policyUUID)
		if policy, err = policyService.GET(iam.PolicyLevels.Account, accountID, policyUUID); err != nil {
			t.Error(err)
			return
		}
		policyBytes, _ := json.Marshal(policy)
		fmt.Println(string(policyBytes))
	}
}

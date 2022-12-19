package iam_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/dtcookie/dynatrace/api/accounts/iam"
)

func TestUserServiceClient(t *testing.T) {
	t.Skip()
	clientID := os.Getenv("DT_CLIENT_ID")
	accountID := os.Getenv("DT_ACCOUNT_ID")
	clientSecret := os.Getenv("DT_CLIENT_SECRET")
	if len(clientID) == 0 || len(accountID) == 0 || len(clientSecret) == 0 {
		t.Skip("DT_CLIENT_ID, DT_ACCOUNT_ID and DT_CLIENT_SECRET required as environment variables - skipping test")
	}

	groupService := iam.NewGroupService(clientID, accountID, clientSecret)
	groupIDs, err := groupService.LIST()
	if err != nil {
		t.Error(err)
		return
	}
	for _, groupID := range groupIDs {
		fmt.Println(groupID)
	}

	service := iam.NewUserService(clientID, accountID, clientSecret)
	user := iam.User{
		Email:  "josef.hickersberger@gmail.com",
		Groups: groupIDs,
	}
	email, err := service.CREATE(&user)
	if err != nil {
		t.Error(err)
		return
	}
	if email != user.Email {
		t.Error(fmt.Errorf("user %s created email address %s differs", user.Email, email))
	}
	ids, err := service.LIST()
	if err != nil {
		t.Error(err)
	}
	found := false
	for _, id := range ids {
		if id == email {
			found = true
		}
	}
	if !found {
		t.Error(fmt.Errorf("user %s created but not found", email))
	}
	if err = service.DELETE(email); err != nil {
		t.Error(err)
	}
}

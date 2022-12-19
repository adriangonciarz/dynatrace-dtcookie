package iam

import (
	"encoding/json"
	"fmt"
)

type UserServiceClient struct {
	clientID     string
	accountID    string
	clientSecret string
}

func (me *UserServiceClient) ClientID() string {
	return me.clientID
}

func (me *UserServiceClient) AccountID() string {
	return me.accountID
}

func (me *UserServiceClient) ClientSecret() string {
	return me.clientSecret
}

func NewUserService(clientID string, accountID string, clientSecret string) *UserServiceClient {
	return &UserServiceClient{clientID: clientID, accountID: accountID, clientSecret: clientSecret}
}

func (me *UserServiceClient) CREATE(user *User) (string, error) {
	var err error

	client := NewIAMClient(me)
	if _, err = client.POST(fmt.Sprintf("https://api.dynatrace.com/iam/v1/accounts/%s/users", me.AccountID()), user, 201, false); err != nil {
		return "", err
	}

	groups := []string{}
	if len(user.Groups) > 0 {
		groups = user.Groups
	}
	if _, err = client.PUT(fmt.Sprintf("https://api.dynatrace.com/iam/v1/accounts/%s/users/%s/groups", me.AccountID(), user.Email), groups, 200, false); err != nil {
		return "", err
	}

	return user.Email, nil
}

type GroupStub struct {
	GroupName string `json:"groupName"`
	UUID      string `json:"uuid"`
}

type GetUserGroupsResponse struct {
	Groups []*GroupStub
}

func (me *UserServiceClient) GET(email string) (interface{}, error) {
	return me.Get(email)
}

func contains(elems []string, elem string) bool {
	if len(elems) == 0 {
		return false
	}
	for _, el := range elems {
		if el == elem {
			return true
		}
	}
	return false
}

func (me *UserServiceClient) Get(email string) (*User, error) {
	var err error
	var responseBytes []byte

	client := NewIAMClient(me)

	if responseBytes, err = client.GET(fmt.Sprintf("https://api.dynatrace.com/iam/v1/accounts/%s/users/%s", me.AccountID(), email), 200, false); err != nil {
		return nil, err
	}

	var response GetUserGroupsResponse
	if err = json.Unmarshal(responseBytes, &response); err != nil {
		return nil, err
	}
	user := User{
		Email:  email,
		Groups: []string{},
	}
	groupService := NewGroupService(me.clientID, me.accountID, me.clientID)
	var visibleGroupIDs []string
	if visibleGroupIDs, err = groupService.LIST(); err != nil {
		return nil, err
	}
	for _, group := range response.Groups {
		if contains(visibleGroupIDs, group.UUID) {
			user.Groups = append(user.Groups, group.UUID)
		}
	}

	return &user, nil
}

func (me *UserServiceClient) UPDATE(user *User) error {
	var err error

	client := NewIAMClient(me)

	if _, err = client.PUT(fmt.Sprintf("https://api.dynatrace.com/iam/v1/accounts/%s/users/%s/groups", me.AccountID(), user.Email), user.Groups, 200, false); err != nil {
		return err
	}
	groups := []string{}
	if len(user.Groups) > 0 {
		groups = user.Groups
	}
	if _, err = client.PUT(fmt.Sprintf("https://api.dynatrace.com/iam/v1/accounts/%s/users/%s/groups", me.AccountID(), user.Email), groups, 200, false); err != nil {
		return err
	}

	return nil
}

type UserStub struct {
	UID   string `json:"uid"`
	Email string `json:"email"`
}

type ListUsersResponse struct {
	Count int        `json:"count:"`
	Items []UserStub `json:"items"`
}

func (me *UserServiceClient) List() ([]UserStub, error) {
	var err error
	var responseBytes []byte

	if responseBytes, err = NewIAMClient(me).GET(fmt.Sprintf("https://api.dynatrace.com/iam/v1/accounts/%s/users", me.AccountID()), 200, false); err != nil {
		return nil, err
	}

	var response ListUsersResponse
	if err = json.Unmarshal(responseBytes, &response); err != nil {
		return nil, err
	}
	return response.Items, nil
}

func (me *UserServiceClient) LIST() ([]string, error) {
	var err error
	var userStubs []UserStub
	if userStubs, err = me.List(); err != nil {
		return nil, err
	}
	ids := []string{}
	for _, stub := range userStubs {
		ids = append(ids, stub.Email)
	}
	return ids, nil
}

func (me *UserServiceClient) DELETE(email string) error {
	_, err := NewIAMClient(me).DELETE(fmt.Sprintf("https://api.dynatrace.com/iam/v1/accounts/%s/users/%s", me.AccountID(), email), 200, false)
	return err
}

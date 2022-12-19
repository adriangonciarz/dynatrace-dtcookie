package iam

import (
	"encoding/json"
	"fmt"
)

type GroupServiceClient struct {
	clientID     string
	accountID    string
	clientSecret string
}

func (me *GroupServiceClient) ClientID() string {
	return me.clientID
}

func (me *GroupServiceClient) AccountID() string {
	return me.accountID
}

func (me *GroupServiceClient) ClientSecret() string {
	return me.clientSecret
}

func NewGroupService(clientID string, accountID string, clientSecret string) *GroupServiceClient {
	return &GroupServiceClient{clientID: clientID, accountID: accountID, clientSecret: clientSecret}
}

func (me *GroupServiceClient) CREATE(group *Group) (string, error) {
	var err error
	var responseBytes []byte

	client := NewIAMClient(me)
	if responseBytes, err = client.POST(fmt.Sprintf("https://api.dynatrace.com/iam/v1/accounts/%s/groups", me.AccountID()), []*Group{group}, 201, false); err != nil {
		return "", err
	}

	responseGroups := []*ListGroup{}
	if err = json.Unmarshal(responseBytes, &responseGroups); err != nil {
		return "", err
	}
	groupID := responseGroups[0].UUID

	if len(group.Permissions) > 0 {
		if _, err = client.PUT(fmt.Sprintf("https://api.dynatrace.com/iam/v1/accounts/%s/groups/%s/permissions", me.AccountID(), groupID), group.Permissions, 200, false); err != nil {
			return "", err
		}
	}

	return groupID, nil
}

func (me *GroupServiceClient) UPDATE(group *Group, uuid string) error {
	var err error

	client := NewIAMClient(me)
	if _, err = client.POST(fmt.Sprintf("https://api.dynatrace.com/iam/v1/accounts/%s/groups/%s", me.AccountID(), uuid), []*Group{group}, 201, false); err != nil {
		return err
	}

	permissions := []*Permission{}

	if len(group.Permissions) > 0 {
		permissions = group.Permissions
	}
	if _, err = client.PUT(fmt.Sprintf("https://api.dynatrace.com/iam/v1/accounts/%s/groups/%s/permissions", me.AccountID(), uuid), permissions, 200, false); err != nil {
		return err
	}

	return nil
}

type ListGroup struct {
	UUID                     string      `json:"uuid"`
	Name                     string      `json:"name"`
	Description              string      `json:"description"`
	FederatedAttributeValues []string    `json:"federatedAttributeValues"`
	Permissions              Permissions `json:"permissions"`
}

type ListGroupsResponse struct {
	Count int          `json:"count:"`
	Items []*ListGroup `json:"items"`
}

func (me *GroupServiceClient) ListInterface() (interface{}, error) {
	return me.List()
}

func (me *GroupServiceClient) List() ([]*ListGroup, error) {
	var err error
	var responseBytes []byte

	if responseBytes, err = NewIAMClient(me).GET(fmt.Sprintf("https://api.dynatrace.com/iam/v1/accounts/%s/groups", me.AccountID()), 200, false); err != nil {
		return nil, err
	}

	var response ListGroupsResponse
	if err = json.Unmarshal(responseBytes, &response); err != nil {
		return nil, err
	}
	return response.Items, nil
}

func (me *GroupServiceClient) LIST() ([]string, error) {
	var err error
	var stubs []*ListGroup
	if stubs, err = me.List(); err != nil {
		return nil, err
	}
	ids := []string{}
	for _, stub := range stubs {
		ids = append(ids, stub.UUID)
	}
	return ids, nil
}

func (me *GroupServiceClient) GET(id string) (interface{}, error) {
	return me.Get(id)
}

func (me *GroupServiceClient) Get(id string) (*Group, error) {
	var err error
	var groupStubs []*ListGroup

	if groupStubs, err = me.List(); err != nil {
		return nil, err
	}

	for _, groupStub := range groupStubs {
		if groupStub.UUID == id {
			var responseBytes []byte

			if responseBytes, err = NewIAMClient(me).GET(fmt.Sprintf("https://api.dynatrace.com/iam/v1/accounts/%s/groups/%s/permissions", me.AccountID(), id), 200, false); err != nil {
				return nil, err
			}

			if err = json.Unmarshal(responseBytes, &groupStub); err != nil {
				return nil, err
			}
			return &Group{
				Name:                     groupStub.Name,
				Description:              groupStub.Description,
				FederatedAttributeValues: groupStub.FederatedAttributeValues,
				Permissions:              groupStub.Permissions,
			}, nil
		}
	}

	return nil, fmt.Errorf("a group with id %s doesn't exist", id)
}

func (me *GroupServiceClient) Delete(id string) error {
	_, err := NewIAMClient(me).DELETE(fmt.Sprintf("https://api.dynatrace.com/iam/v1/accounts/%s/groups/%s", me.AccountID(), id), 200, false)
	return err
}

func (me *GroupServiceClient) DELETE(id string) error {
	return me.Delete(id)
}

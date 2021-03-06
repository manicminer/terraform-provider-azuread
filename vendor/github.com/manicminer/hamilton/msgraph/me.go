package msgraph

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// MeClient performs operations on the authenticated user.
type MeClient struct {
	BaseClient Client
}

// NewMeClient returns a new MeClient.
func NewMeClient(tenantId string) *MeClient {
	return &MeClient{
		BaseClient: NewClient(VersionBeta, tenantId),
	}
}

// Get retrieves information about the authenticated user.
func (c *MeClient) Get(ctx context.Context) (*Me, int, error) {
	var status int
	resp, status, _, err := c.BaseClient.Get(ctx, GetHttpRequestInput{
		ValidStatusCodes: []int{http.StatusOK},
		Uri: Uri{
			Entity:      "/me",
			HasTenantId: false,
		},
	})
	if err != nil {
		return nil, status, fmt.Errorf("MeClient.BaseClient.Get(): %v", err)
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, status, fmt.Errorf("ioutil.ReadAll(): %v", err)
	}
	var me Me
	if err := json.Unmarshal(respBody, &me); err != nil {
		return nil, status, fmt.Errorf("json.Unmarshal(): %v", err)
	}
	return &me, status, nil
}

// GetProfile retrieves the profile of the authenticated user.
func (c *MeClient) GetProfile(ctx context.Context) (*Me, int, error) {
	var status int
	resp, status, _, err := c.BaseClient.Get(ctx, GetHttpRequestInput{
		ValidStatusCodes: []int{http.StatusOK},
		Uri: Uri{
			Entity:      "/me/profile",
			HasTenantId: false,
		},
	})
	if err != nil {
		return nil, status, fmt.Errorf("MeClient.BaseClient.Get(): %v", err)
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, status, fmt.Errorf("ioutil.ReadAll(): %v", err)
	}
	var me Me
	if err := json.Unmarshal(respBody, &me); err != nil {
		return nil, status, fmt.Errorf("json.Unmarshal(): %v", err)
	}
	return &me, status, nil
}

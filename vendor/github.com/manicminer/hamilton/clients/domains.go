package clients

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/manicminer/hamilton/base"
	"github.com/manicminer/hamilton/models"
)

type DomainsClient struct {
	BaseClient base.Client
}

func NewDomainsClient(tenantId string) *DomainsClient {
	return &DomainsClient{
		BaseClient: base.NewClient(base.DefaultEndpoint, tenantId, base.Version10),
	}
}

func (c *DomainsClient) List(ctx context.Context) (*[]models.Domain, int, error) {
	var status int
	resp, status, err := c.BaseClient.Get(ctx, base.GetHttpRequestInput{
		ValidStatusCodes: []int{http.StatusOK},
		Uri: base.Uri{
			Entity:      "/domains",
			HasTenantId: true,
		},
	})

	if err != nil {
		return nil, status, err
	}

	defer resp.Body.Close()
	respBody, _ := ioutil.ReadAll(resp.Body)

	var data struct {
		Domains []models.Domain `json:"value"`
	}

	if err := json.Unmarshal(respBody, &data); err != nil {
		return nil, status, err
	}

	return &data.Domains, status, nil
}

func (c *DomainsClient) Get(ctx context.Context, id string) (*models.Domain, int, error) {
	var status int
	resp, status, err := c.BaseClient.Get(ctx, base.GetHttpRequestInput{
		ValidStatusCodes: []int{http.StatusOK},
		Uri: base.Uri{
			Entity:      fmt.Sprintf("/domains/%s", id),
			HasTenantId: true,
		},
	})
	if err != nil {
		return nil, status, err
	}
	defer resp.Body.Close()
	respBody, _ := ioutil.ReadAll(resp.Body)
	var domain models.Domain
	if err := json.Unmarshal(respBody, &domain); err != nil {
		return nil, status, err
	}
	return &domain, status, nil
}
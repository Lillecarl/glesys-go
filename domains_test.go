package glesys

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDomainsAdd(t *testing.T) {
	c := &mockClient{body: `{ "response": { "domain":
		{ "domainname": "sysgle.se" }}}`}
	d := DomainService{client: c}

	params := AddDomainParams{
		DomainName: "sysgle.se",
	}

	domain, _ := d.Add(context.Background(), params)

	assert.Equal(t, "POST", c.lastMethod, "method is used correct")
	assert.Equal(t, "domain/add", c.lastPath, "path used is correct")
	assert.Equal(t, "sysgle.se", domain.DomainName, "domain name is correct")
}

func TestDomainsDestroy(t *testing.T) {
	c := &mockClient{}
	d := DomainService{client: c}

	d.Delete(context.Background(), "sysgle.se")

	assert.Equal(t, "POST", c.lastMethod, "method used is correct")
	assert.Equal(t, "domain/delete", c.lastPath, "path used is correct")
}

/*
func TestNetworksDetails(t *testing.T) {
	c := &mockClient{body: `{ "response": { "network": {
		"networkid": "vl123456",
		"description": "My Network",
		"datacenter": "Falkenberg",
		"public": "no"
		} } }`}
	s := NetworkService{client: c}

	network, _ := s.Details(context.Background(), "vl123456")

	assert.Equal(t, "GET", c.lastMethod, "method used is correct")
	assert.Equal(t, "network/details/networkid/vl123456", c.lastPath, "path used is correct")
	assert.Equal(t, "Falkenberg", network.DataCenter, "network DataCenter is correct")
	assert.Equal(t, "My Network", network.Description, "network Description is correct")
	assert.Equal(t, "vl123456", network.ID, "network ID is correct")
	assert.Equal(t, "no", network.Public, "network Public is correct")
}

func TestNetworksEdit(t *testing.T) {
	c := &mockClient{body: `{ "response": { "network":
		{ "datacenter": "Falkenberg", "description": "mynewnetwork", "networkid": "vl123456" }}}`}
	n := NetworkService{client: c}

	params := EditNetworkParams{
		Description: "mynetwork",
	}

	network, _ := n.Edit(context.Background(), "vl123456", params)

	assert.Equal(t, "POST", c.lastMethod, "method is used correct")
	assert.Equal(t, "network/edit", c.lastPath, "path used is correct")
	assert.Equal(t, "vl123456", network.ID, "network ID is correct")
	assert.Equal(t, "Falkenberg", network.DataCenter, "network DataCenter is correct")
	assert.Equal(t, "mynewnetwork", network.Description, "network Description is correct")
}

func TestNetworksIsPublic(t *testing.T) {
	network := Network{Public: "yes"}
	assert.Equal(t, true, network.IsPublic(), "should be public")

	network.Public = "no"
	assert.Equal(t, false, network.IsPublic(), "should not be public")
}

func TestNetworksList(t *testing.T) {
	c := &mockClient{body: `{ "response": { "networks":
	[{ "datacenter": "Falkenberg", "description": "Internet", "networkid": "internet-fbg", "public": "yes"}] } }`}
	n := NetworkService{client: c}

	networks, _ := n.List(context.Background())

	assert.Equal(t, "GET", c.lastMethod, "method used is correct")
	assert.Equal(t, "network/list", c.lastPath, "path used is correct")
	assert.Equal(t, "Falkenberg", (*networks)[0].DataCenter, "network DataCenter is correct")
	assert.Equal(t, "yes", (*networks)[0].Public, "network is public")
	assert.Equal(t, "Internet", (*networks)[0].Description, "network Description is correct")
	assert.Equal(t, "internet-fbg", (*networks)[0].ID, "network ID is correct")
}
*/

package cdncontainers

import (
	"testing"

	"github.com/LunaNode/gophercloud"
	th "github.com/LunaNode/gophercloud/testhelper"
)

const endpoint = "http://localhost:57909/"

func endpointClient() *gophercloud.ServiceClient {
	return &gophercloud.ServiceClient{Endpoint: endpoint}
}

func TestEnableURL(t *testing.T) {
	actual := enableURL(endpointClient(), "foo")
	expected := endpoint + "foo"
	th.CheckEquals(t, expected, actual)
}

package common

import (
	"github.com/LunaNode/gophercloud"
	"github.com/LunaNode/gophercloud/testhelper/client"
)

const TokenID = client.TokenID

func ServiceClient() *gophercloud.ServiceClient {
	return client.ServiceClient()
}

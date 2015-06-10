package image

import "github.com/LunaNode/gophercloud"

func createURL(client *gophercloud.ServiceClient) string {
	return client.ServiceURL("images")
}

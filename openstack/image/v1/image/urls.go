package image

import "github.com/LunaNode/gophercloud"

func createURL(client *gophercloud.ServiceClient) string {
	return client.ServiceURL("images")
}

func imageURL(client *gophercloud.ServiceClient, id string) string {
	return client.ServiceURL("images", id)
}

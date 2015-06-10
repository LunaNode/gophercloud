package cdncontainers

import "github.com/LunaNode/gophercloud"

func enableURL(c *gophercloud.ServiceClient, containerName string) string {
	return c.ServiceURL(containerName)
}

func getURL(c *gophercloud.ServiceClient, container string) string {
	return c.ServiceURL(container)
}

func updateURL(c *gophercloud.ServiceClient, container string) string {
	return getURL(c, container)
}

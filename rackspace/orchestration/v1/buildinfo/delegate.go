package buildinfo

import (
	"github.com/LunaNode/gophercloud"
	os "github.com/LunaNode/gophercloud/openstack/orchestration/v1/buildinfo"
)

// Get retreives build info data for the Heat deployment.
func Get(c *gophercloud.ServiceClient) os.GetResult {
	return os.Get(c)
}

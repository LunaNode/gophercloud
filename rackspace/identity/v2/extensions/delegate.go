package extensions

import (
	"github.com/LunaNode/gophercloud"
	common "github.com/LunaNode/gophercloud/openstack/common/extensions"
	"github.com/LunaNode/gophercloud/pagination"
)

// ExtractExtensions accepts a Page struct, specifically an ExtensionPage struct, and extracts the
// elements into a slice of os.Extension structs.
func ExtractExtensions(page pagination.Page) ([]common.Extension, error) {
	return common.ExtractExtensions(page)
}

// Get retrieves information for a specific extension using its alias.
func Get(c *gophercloud.ServiceClient, alias string) common.GetResult {
	return common.Get(c, alias)
}

// List returns a Pager which allows you to iterate over the full collection of extensions.
// It does not accept query parameters.
func List(c *gophercloud.ServiceClient) pagination.Pager {
	return common.List(c)
}

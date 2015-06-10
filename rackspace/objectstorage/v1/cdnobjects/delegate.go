package cdnobjects

import (
	"github.com/LunaNode/gophercloud"
	os "github.com/LunaNode/gophercloud/openstack/objectstorage/v1/objects"
)

// Delete is a function that deletes an object from the CDN.
func Delete(c *gophercloud.ServiceClient, containerName, objectName string, opts os.DeleteOptsBuilder) os.DeleteResult {
	return os.Delete(c, containerName, objectName, nil)
}

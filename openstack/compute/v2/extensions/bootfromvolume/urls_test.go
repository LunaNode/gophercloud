package bootfromvolume

import (
	"testing"

	th "github.com/LunaNode/gophercloud/testhelper"
	"github.com/LunaNode/gophercloud/testhelper/client"
)

func TestCreateURL(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	c := client.ServiceClient()

	th.CheckEquals(t, c.Endpoint+"os-volumes_boot", createURL(c))
}

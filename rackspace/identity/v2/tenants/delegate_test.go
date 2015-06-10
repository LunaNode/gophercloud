package tenants

import (
	"testing"

	os "github.com/LunaNode/gophercloud/openstack/identity/v2/tenants"
	"github.com/LunaNode/gophercloud/pagination"
	th "github.com/LunaNode/gophercloud/testhelper"
	fake "github.com/LunaNode/gophercloud/testhelper/client"
)

func TestListTenants(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	os.HandleListTenantsSuccessfully(t)

	count := 0
	err := List(fake.ServiceClient(), nil).EachPage(func(page pagination.Page) (bool, error) {
		actual, err := ExtractTenants(page)
		th.AssertNoErr(t, err)
		th.CheckDeepEquals(t, os.ExpectedTenantSlice, actual)

		count++
		return true, nil
	})
	th.AssertNoErr(t, err)
	th.CheckEquals(t, 1, count)
}

package extensions

import (
	"testing"

	common "github.com/LunaNode/gophercloud/openstack/common/extensions"
	"github.com/LunaNode/gophercloud/pagination"
	th "github.com/LunaNode/gophercloud/testhelper"
	fake "github.com/LunaNode/gophercloud/testhelper/client"
)

func TestList(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	common.HandleListExtensionsSuccessfully(t)

	count := 0

	err := List(fake.ServiceClient()).EachPage(func(page pagination.Page) (bool, error) {
		count++
		actual, err := ExtractExtensions(page)
		th.AssertNoErr(t, err)
		th.AssertDeepEquals(t, common.ExpectedExtensions, actual)

		return true, nil
	})
	th.AssertNoErr(t, err)
	th.CheckEquals(t, 1, count)
}

func TestGet(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	common.HandleGetExtensionSuccessfully(t)

	actual, err := Get(fake.ServiceClient(), "agent").Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, common.SingleExtension, actual)
}

package tokens

import (
	"testing"

	"github.com/LunaNode/gophercloud"
	"github.com/LunaNode/gophercloud/testhelper"
)

func TestTokenURL(t *testing.T) {
	testhelper.SetupHTTP()
	defer testhelper.TeardownHTTP()

	client := gophercloud.ServiceClient{Endpoint: testhelper.Endpoint()}

	expected := testhelper.Endpoint() + "auth/tokens"
	actual := tokenURL(&client)
	if actual != expected {
		t.Errorf("Expected URL %s, but was %s", expected, actual)
	}
}

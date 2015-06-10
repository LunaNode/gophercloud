package stacktemplates

import (
	"testing"

	os "github.com/LunaNode/gophercloud/openstack/orchestration/v1/stacktemplates"
	th "github.com/LunaNode/gophercloud/testhelper"
	fake "github.com/LunaNode/gophercloud/testhelper/client"
)

func TestGetTemplate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	os.HandleGetSuccessfully(t, os.GetOutput)

	actual, err := Get(fake.ServiceClient(), "postman_stack", "16ef0584-4458-41eb-87c8-0dc8d5f66c87").Extract()
	th.AssertNoErr(t, err)

	expected := os.GetExpected
	th.AssertDeepEquals(t, expected, actual)
}

func TestValidateTemplate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	os.HandleValidateSuccessfully(t, os.ValidateOutput)

	opts := os.ValidateOpts{
		Template: map[string]interface{}{
			"heat_template_version": "2013-05-23",
			"description":           "Simple template to test heat commands",
			"parameters": map[string]interface{}{
				"flavor": map[string]interface{}{
					"default": "m1.tiny",
					"type":    "string",
				},
			},
			"resources": map[string]interface{}{
				"hello_world": map[string]interface{}{
					"type": "OS::Nova::Server",
					"properties": map[string]interface{}{
						"key_name": "heat_key",
						"flavor": map[string]interface{}{
							"get_param": "flavor",
						},
						"image":     "ad091b52-742f-469e-8f3c-fd81cadf0743",
						"user_data": "#!/bin/bash -xv\necho \"hello world\" &gt; /root/hello-world.txt\n",
					},
				},
			},
		},
	}
	actual, err := Validate(fake.ServiceClient(), opts).Extract()
	th.AssertNoErr(t, err)

	expected := os.ValidateExpected
	th.AssertDeepEquals(t, expected, actual)
}

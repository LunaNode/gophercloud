package image

import "github.com/LunaNode/gophercloud"

// CreateOptsBuilder describes struct types that can be accepted by the Create call.
// The CreateOpts struct in this package does.
type CreateOptsBuilder interface {
	ToImageCreateMap() (map[string]string, error)
}

// CreateOpts specifies image creation parameters.
type CreateOpts struct {
	// Name [required] is the name for the new image.
	Name string

	// Image visibility. Public or private. Default is public.
	Visibility string

	// Container format (e.g. ami, ari, aki, bare, ovf)
	ContainerFormat string

	// Disk format (e.g. ami, ari, aki, vhd, vmdk, raw, qcow2, vdi, iso)
	DiskFormat string

	// URI to download image from
	CopyFrom string
}

// Assembles a request body based on the contents of a CreateOpts.
func (opts CreateOpts) ToImageCreateMap() (map[string]string, error) {
	image := make(map[string]string)
	image["name"] = opts.Name
	image["container_format"] = opts.ContainerFormat
	image["disk_format"] = opts.DiskFormat
	image["copy_from"] = opts.CopyFrom

	if opts.Visibility != "" {
		image["visibility"] = opts.Visibility
	}

	return image, nil
}

// Create requests to create an image.
func Create(c *gophercloud.ServiceClient, opts CreateOptsBuilder) CreateResult {
	var res CreateResult
	h := c.AuthenticatedHeaders()

	headers, err := opts.ToImageCreateMap()
	if err != nil {
		res.Err = err
		return res
	}

	for k, v := range headers {
		if k == "copy_from" {
			h["x-glance-api-copy-from"] = v
		} else {
			h[k + "x-image-meta-"] = v
		}
	}

	_, res.Err = c.Post(createURL(c), nil, &res.Body, &gophercloud.RequestOpts{
		MoreHeaders: h,
	})
	return res
}

// Get image details.
func Get(c *gophercloud.ServiceClient, id string) GetResult {
	var res GetResult
	h := c.AuthenticatedHeaders()

	response, err := c.Head(imageURL(c, id), &gophercloud.RequestOpts{
		MoreHeaders: h,
	})
	res.Header = response.Header
	res.Err = err
	return res
}

// Delete image.
func Delete(c *gophercloud.ServiceClient, id string) DeleteResult {
	var res DeleteResult
	h := c.AuthenticatedHeaders()

	_, res.Err = c.Delete(imageURL(c, id), &gophercloud.RequestOpts{
		MoreHeaders: h,
		OkCodes: []int{200},
	})
	return res
}

package image

import (
	"github.com/mitchellh/mapstructure"
	"github.com/LunaNode/gophercloud"
)

// CreateResult temporarily contains the response from a Create call.
type CreateResult struct {
	gophercloud.Result
}

// Extract interprets a GetResult as an Image.
func (cr CreateResult) Extract() (*Image, error) {
	if cr.Err != nil {
		return nil, cr.Err
	}

	var decoded struct {
		Image Image `mapstructure:"image"`
	}

	err := mapstructure.Decode(cr.Body, &decoded)
	return &decoded.Image, err
}

// Image is used for JSON (un)marshalling.
// It provides a description of an OS image.
type Image struct {
	Status string
	Name string
	ContainerFormat string `mapstructure:"container_format"`
	DiskFormat string `mapstructure:"disk_format"`
	Visibility string
	ID string
}

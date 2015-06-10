package image

import (
	"github.com/mitchellh/mapstructure"
	"github.com/LunaNode/gophercloud"
	"strings"
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
	Size string
	ID string
}

type GetResult struct {
	gophercloud.Result
}

func (r GetResult) Extract() (*Image, error) {
	if r.Err != nil {
		return nil, r.Err
	}

	var image Image
	for k, vArray := range r.Header {
		k = strings.ToLower(k)
		if strings.HasPrefix(k, "x-image-meta-") {
			v := vArray[0]
			k = strings.Split(k, "x-image-meta-")[1]
			if k == "size" {
				image.Size = v
			} else if k == "status" {
				image.Status = v
			} else if k == "name" {
				image.Name = v
			} else if k == "container_format" {
				image.ContainerFormat = v
			} else if k == "disk_format" {
				image.DiskFormat = v
			} else if k == "visibility" {
				image.Visibility = v
			}
		}
	}

	return &image, nil
}

type DeleteResult struct {
	gophercloud.ErrResult
}

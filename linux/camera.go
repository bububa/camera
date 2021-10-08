//go:build linux
// +build linux

package linux

import (
	"fmt"
	"image"

	"github.com/korandiz/v4l"
	"github.com/korandiz/v4l/fmt/mjpeg"

	"github.com/bububa/camera"
	im "github.com/bububa/camera/image"
)

// Camera represents v4l camera
type Camera struct {
	opts   camera.Options
	device *v4l.Device
}

// New returns new Camera for given camera index.
func New(opts camera.Options) (c *Camera, err error) {
	c = &Camera{
		opts: opts,
	}
	devices := v4l.FindDevices()
	if opts.Index >= len(devices) {
		err = fmt.Errorf("camera: no camera at index %d", opts.Index)
		return
	}
	if c.device, err = v4l.Open(devices[opts.Index].Path); err != nil {
		err = fmt.Errorf("camera: %s", err.Error())
		return
	}

	config, err := c.device.GetConfig()
	if err != nil {
		err = fmt.Errorf("camera: %s", err.Error())
		return
	}

	config.Format = mjpeg.FourCC
	config.Width = int(opts.Width)
	config.Height = int(opts.Height)

	err = c.device.SetConfig(config)
	if err != nil {
		err = fmt.Errorf("camera: %s", err.Error())
		return
	}
	return c, nil
}

// Start implement Device interface
func (c *Camera) Start() error {
	return c.device.TurnOn()
}

// Read reads next frame from camera and returns image.
func (c *Camera) Read() (img image.Image, err error) {
	if c.device == nil {
		err = fmt.Errorf("camera: nil device")
		return
	}
	buffer, err := c.device.Capture()
	if err != nil {
		err = fmt.Errorf("camera: can not grab frame: %s", err.Error())
		return
	}

	img, err = im.NewDecoder(buffer).Decode()
	if err != nil {
		err = fmt.Errorf("camera: %s", err.Error())
		return
	}

	return
}

// GetProperty returns the specified camera property.
func (c *Camera) GetProperty(id int) float64 {
	ret, _ := c.device.GetControl(uint32(id))
	return float64(ret)
}

// SetProperty sets a camera property.
func (c *Camera) SetProperty(id int, value float64) {
	c.device.SetControl(uint32(id), int32(value))
}

// Close closes camera.
func (c *Camera) Close() (err error) {
	if c.device == nil {
		err = fmt.Errorf("camera: camera is not opened")
		return
	}

	c.device.TurnOff()
	c.device.Close()
	c.device = nil
	return
}

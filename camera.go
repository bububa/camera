package camera

import "image"

// Camera represents camera
type Camera struct {
	device Device
}

// NewCamera returns new Camera for given device
func NewCamera(device Device) *Camera {
	return &Camera{
		device: device,
	}
}

// Start device start recording
func (c *Camera) Start() error {
	return c.device.Start()
}

// Close to close the camera
func (c *Camera) Close() error {
	return c.device.Close()
}

// Start device start recording
func (c *Camera) Read() (image.Image, error) {
	return c.device.Read()
}

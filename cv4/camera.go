//go:build cv4
// +build cv4

package cv4

import (
	"fmt"
	"image"

	"gocv.io/x/gocv"

	"github.com/bububa/camera"
)

// Camera represents camera.
type Camera struct {
	opts   camera.Options
	device *gocv.VideoCapture
	frame  *gocv.Mat
}

// New returns new Camera for given camera index.
func New(opts camera.Options) (c *Camera, err error) {
	c = &Camera{
		opts: opts,
	}
	mat := gocv.NewMat()
	c.frame = &mat
	c.device, err = gocv.VideoCaptureDevice(opts.Index)
	if err != nil {
		return nil, err
	}
	c.SetProperty(PropFrameWidth, opts.Width)
	c.SetProperty(PropFrameHeight, opts.Height)
	return c, nil
}

// Start implement Device interface
func (c *Camera) Start() error {
	return nil
}

// Read reads next frame from camera and returns image.
func (c *Camera) Read() (img image.Image, err error) {
	if !c.device.IsOpened() {
		err = fmt.Errorf("camera: closed")
		return
	}
	if ok := c.device.Read(c.frame); !ok {
		err = fmt.Errorf("camera: can not grab frame")
		return
	}

	if c.frame == nil {
		err = fmt.Errorf("camera: can not retrieve frame")
		return
	}

	img, e := c.frame.ToImage()
	if e != nil {
		err = fmt.Errorf("camera: %v", e)
		return
	}

	return
}

// GetProperty returns the specified camera property.
func (c *Camera) GetProperty(id int) float64 {
	return c.device.Get(gocv.VideoCaptureProperties(id))
}

// SetProperty sets a camera property.
func (c *Camera) SetProperty(id int, value float64) {
	c.device.Set(gocv.VideoCaptureProperties(id), value)
}

// Close closes camera.
func (c *Camera) Close() (err error) {
	if c.device == nil {
		err = fmt.Errorf("camera: camera is not opened")
		return
	}

	c.frame.Close()
	err = c.device.Close()
	c.device = nil
	return
}

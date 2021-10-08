package camera

import (
	"image"
)

// Device defines device interface
type Device interface {
	Start() error
	Read() (image.Image, error)
	Close() error
	GetProperty(int) float64
	SetProperty(int, float64)
}

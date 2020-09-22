package luxafor

import (
	"fmt"

	"github.com/google/gousb"
)

// Color represents a color state that the Luxafor Flag can be set to.
type Color byte

var (
	flagVendorID  gousb.ID = 0x04d8
	flagProductID gousb.ID = 0xf372

	// Green is the color Green.
	Green Color = 'G'

	// Red is the color Red.
	Red Color = 'R'

	// Blue is the color Blue.
	Blue Color = 'B'

	// Magenta is the color Magenta.
	Magenta Color = 'M'

	// Yellow is the color Yellow.
	Yellow Color = 'Y'

	// Off will turn off the LED. Not strictly a color.
	Off Color = 'O'

	// ErrNoFlag will be returned if no Luxafor Flag can be found.
	ErrNoFlag error = fmt.Errorf("luxafor: No flag found")
)

// Flag represents a handle to a Luxafor Flag.
type Flag struct {
	context *gousb.Context
	closer  func() error
	oe      *gousb.OutEndpoint
}

// Close will close and free all held resources.
func (f Flag) Close() error {
	return f.closer()
}

// OpenFlag will find and open the first Luxafor Flag that is found
// attached to the system.
func OpenFlag() (*Flag, error) {
	ctx := gousb.NewContext()
	dev, err := ctx.OpenDeviceWithVIDPID(flagVendorID, flagProductID)
	if err != nil {
		return nil, err
	}

	if dev == nil {
		return nil, ErrNoFlag
	}

	if err := dev.SetAutoDetach(true); err != nil {
		return nil, err
	}

	iface, closer, err := dev.DefaultInterface()
	if err != nil {
		return nil, err
	}

	oe, err := iface.OutEndpoint(0x01)
	if err != nil {
		return nil, err
	}

	return &Flag{
		context: ctx,
		oe:      oe,
		closer: func() error {
			closer()
			ctx.Close()
			// TODO(paultag): handle error here better
			return nil
		},
	}, nil
}

// SetColor will ask the Luxafor Flag to change colors to the desired color.
func (f Flag) SetColor(color Color) error {
	_, err := f.oe.Write([]byte{0x00, 0x00})
	if err != nil {
		return err
	}

	_, err = f.oe.Write([]byte{0x00, byte(color)})
	return err
}

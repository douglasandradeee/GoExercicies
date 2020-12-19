package model

import "errors"

// Immobile -
type Immobile struct {
	X     int    `json:"coordinate_x"`
	Y     int    `json:"coordinate_y"`
	Name  string `json:"name"`
	value int
}

// ErrInvalidPropertyValue - error presented when the property was assigned a very low value.
var ErrInvalidPropertyValue = errors.New("The informed value is not valid for a property")

// ErrVeryHighPropertyValue - error presented when the property value is too high.
var ErrVeryHighPropertyValue = errors.New("The informed value is too high for a property")

// SetValue - defines the value of the property/Immobile.
func (i *Immobile) SetValue(value int) (err error) {
	err = nil
	if value < 50000 {
		err = ErrInvalidPropertyValue
		return
	} else if value > 10000000 {
		err = ErrVeryHighPropertyValue
		return
	}
	i.value = value
	return
}

// GetValue - returns the value of the property/Immobile.
func (i *Immobile) GetValue() int {
	return i.value
}

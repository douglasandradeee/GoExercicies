package model

// Immobile -
type Immobile struct {
	X     int    `json:"coordinate_x"`
	Y     int    `json:"coordinate_y"`
	Name  string `json:"name"`
	value int
}

// SetValue - defines the value of the property/Immobile.
func (i *Immobile) SetValue(value int) {
	i.value = value
}

// GetValue - returns the value of the property/Immobile.
func (i *Immobile) GetValue() int {
	return i.value
}

package model

// Chicken - It`s a model of bird
type Chicken interface {
	Cackle() string
}

// Duck - It`s a model of bird
type Duck interface {
	Quack() string
}

// Bird -
type Bird struct {
	Name string `json:"name"`
}

// Cackle - It's the characteristic sound of the chicken
func (b Bird) Cackle() string {
	return "Cocoricó..."
}

// Quack - It's the characteristic sound of the duck
func (b Bird) Quack() string {
	return "Quack quack ..."
}

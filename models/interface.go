package models

// Constructor is the interfece for all models to work with the database
type Constructor interface {
	init()
	Create()
}

// Init will take an interface of Constructor and call the internal Init function to set up the database
func Init(i Constructor) {
	i.init()
}

package module

//Module interface defining operations that can be performed on a module
type Module interface {
	//Clone the module, specifying the cache directory as a parameter
	Clone(string) error
	//Update the module, specifying the cache directory as a parameter
	Update(string) error
	//Delete the module, specifying the cache directory as a parameter
	Delete(string) error
}

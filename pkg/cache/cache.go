package cache

//Cache interface defining simple caching mechanisms
type Cache interface {
	//Create cache
	Create() error
	//Update all elements in the cache
	Update() error
	//Delete all elements in the cache
	Delete() error
}

package cache

//Cache interface defining simple caching mechanisms
type Cache interface {
	//Create cache
	Create() error

	//Delete all elements in the cache
	Delete(bool) error
}

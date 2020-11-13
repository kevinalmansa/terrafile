package cache

//InitError is an initialization error
type InitError struct{}

func (e *InitError) Error() string {
	return "LocalModuleCache: Configuration not found"
}

package modules

// Process retrieves and runs input data from user
// GetResult returns result of processed module
type Module interface {
	Process(params string) error
	GetResult() (string, error)

	HandlerName() string
}

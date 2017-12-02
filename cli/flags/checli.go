package flags

// ClientOptions are the options used to configure the client cli
type CliOptions struct {
	ConfigDir string
	Version   bool
}

// NewClientOptions returns a new ClientOptions
func NewCheCliOptions() *CliOptions {
	return &CliOptions{}
}

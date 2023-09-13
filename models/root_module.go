package models

type RootModule struct {
	// name of the module
	Name string

	// description of the module
	Description string

	// a custom hook for convenience to run pre-init and post-init jobs
	// Jobs JobHook `yaml:"jobs, omitempty"`
}

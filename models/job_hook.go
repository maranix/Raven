package models

type JobHook struct {
	// jobs to run before initializing the module
	Pre []string `yaml:"pre"`
	// jobs to run after initializing the module
	Post []string `yaml:"post"`
}

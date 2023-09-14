package models

type JobHook struct {
	// jobs to run before initializing the module
	Pre []Job
	// jobs to run after initializing the module
	Post []Job
}

type Job struct {
	// name of the executable or cli tool
	Exec string
	// rest of the command
	Command string
	// path where this job should run at
	Workdir string `yaml:"workdir,omitempty"`
}

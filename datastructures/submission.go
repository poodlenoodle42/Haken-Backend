package datastructures

type Language struct {
	Name          string `json:"name" yaml:"name"`
	Abbreviation  string `json:"abbreviation" yaml:"abbreviation"`
	PreLaunchTask string `json:"-" yaml:"preLaunchTask"`
	LaunchTask    string `json:"-" yaml:"launchTask"`
}

type LanguageName string

//Submission of a User for a given Task
type Submission struct {
	ID      int    `json:"id"`
	Author  string `json:"-"`
	TaskID  int    `json:"-"`
	GroupID int    `json:"-"`
	//Source Code path is calculated by ID
}

type Result struct {
	Sub  *Submission `json:"submission"`
	Subt *Subtask    `json:"subtask"`
	//Success exit code of program for subtask, -1 indicates success
	Success int `json:"success"`
}

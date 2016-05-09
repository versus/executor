// task.go
package r2d2

import (
	"fmt"
)

type Task struct {
	name    string
	command string `json:"hostname" valid:"alphanum,required"`
	debug   bool
	chdir   string
	shell	string
	error_ignore: bool
	timeout int
}



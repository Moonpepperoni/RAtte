package format

import (
	"fmt"
	"github.com/Moonpepperoni/ratte/dirutil"
)

func RenameSuccess(path string) string {
	return fmt.Sprintf("Renaming file %s was successful!", dirutil.RelativePath(path))
}

func RenameFailure(path string, reason error) string {
	return fmt.Sprintf("Renaming file %s failed:\n%s", dirutil.RelativePath(path), reason.Error())
}

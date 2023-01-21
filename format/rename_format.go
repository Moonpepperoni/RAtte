package format

import "fmt"

func RenameSuccess(path string) string {
	return fmt.Sprintf("Renaming file %s was successful!", path)
}

package dirutil

import (
	"reflect"
	"testing"
)

func TestGetAllSecFiles(t *testing.T) {
	paths := []string{
		"this/path.as",
		"some/path/to/file.asm",
		"some/path/to/file.o",
		"some/path/fileasm.sec",
		"this/otherasm.sec",
		"this/deep/folder/213.asm",
		"works_1asm.sec",
		"/folder/still-valid-asm.sec",
		"/not/in/there",
		"justSomeFile.sec"}
	expected := []string{
		"some/path/fileasm.sec",
		"this/otherasm.sec",
		"works_1asm.sec",
		"/folder/still-valid-asm.sec"}
	secFiles := FilterSecFiles(paths)
	if !reflect.DeepEqual(secFiles, expected) {
		t.Errorf("Expected %v to be equal to %v", secFiles, expected)
	}
}

func TestGetAllAsmFiles(t *testing.T) {
	paths := []string{
		"this/path.as",
		"some/path/to/file.asm",
		"some/path/to/file.o",
		"some/path/file.asm",
		"this/other.asm",
		"this/deep/folder/213.asm",
		"works_1.asm",
		"/folder/still-valid.asm",
		"/not/in/there"}
	expected := []string{
		"some/path/to/file.asm",
		"some/path/file.asm",
		"this/other.asm",
		"this/deep/folder/213.asm",
		"works_1.asm",
		"/folder/still-valid.asm"}
	asmFiles := FilterAsmFiles(paths)
	if !reflect.DeepEqual(asmFiles, expected) {
		t.Errorf("Expected %v to be equal to %v", asmFiles, expected)
	}

}

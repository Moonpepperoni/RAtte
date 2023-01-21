package dirutil

import "testing"

func TestRelativePathSuccess(t *testing.T) {
	absolute := "/this/is/a/path.ext"
	cwd := "/this/is/a"
	relative := relativePath(absolute, cwd)
	if relative != "path.ext" {
		t.Errorf("expected relative to be path.ext but got: %s", relative)
	}
}

func TestRelativePathFolderInRoot(t *testing.T) {
	absolute := "/home"
	cwd := "/"
	relative := relativePath(absolute, cwd)
	if relative != "/home" {
		t.Errorf("expected relative to be /home but got: %s", relative)
	}
}

func TestRelativePathRoot(t *testing.T) {
	absolute := "/"
	cwd := "/"
	relative := relativePath(absolute, cwd)
	if relative != "/" {
		t.Errorf("expected relative to be /  but got: %s", relative)
	}
}

func TestRelativeSlashEndedWorkingDirectory(t *testing.T) {
	absolute := "/this/path/is/absolute/file.asm"
	cwd := "/this/path/is/absolute/"
	relative := relativePath(absolute, cwd)
	if relative != "file.asm" {
		t.Errorf("expected relative to be file.asm but got: %s", relative)
	}
}

func TestRelativePathWorkingDirectoryMismatch(t *testing.T) {
	absolute := "/this/path/is/absolute/file.asm"
	cwd := "/cwd/was/changed"
	relative := relativePath(absolute, cwd)
	if relative != "/this/path/is/absolute/file.asm" {
		t.Errorf("expected relative to be /this/path/is/absolute/file.asm but got: %s", relative)
	}
}

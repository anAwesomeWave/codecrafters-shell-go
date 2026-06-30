package dirscanner

import (
	"errors"
	"fmt"
	"os"
	"slices"
	"strings"

	my_slices "github.com/codecrafters-io/shell-starter-go/app/internal/utils/slices"
)

const (
	EXEC_FILES_PERM = 0o111 // octal
)

var (
	ErrExecNotFound = errors.New("exec not found")
)

func ExecLookup(execName string, paths []string) (string, error) {
	for _, dir := range paths {
		execs, err := ScanExecFiles(dir)
		if err != nil {
			return "", fmt.Errorf("couldn't scan dir %s: %w", dir, err)
		}

		if slices.Contains(my_slices.Map(execs, func(fullPath string) string {
			pathParts := strings.Split(fullPath, "/")

			return pathParts[len(pathParts)-1]
		}), execName) {
			return strings.TrimRight(dir, "/") + "/" + execName, nil
		}
	}

	return "", fmt.Errorf("couldn't find exec %s: %w", execName, ErrExecNotFound)
}

func ScanExecFiles(path string) ([]string, error) {
	execFiles, err := scanFiles(path, EXEC_FILES_PERM)
	if err != nil {
		return nil, fmt.Errorf("scan for exec files: %w", err)
	}

	return execFiles, nil
}

func scanFiles(path string, permFilters uint32) ([]string, error) {
	/*
		read dir and return all of the filenames
		perm filters can be applied
	*/
	files, err := os.ReadDir(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil, nil
		}

		return nil, fmt.Errorf("scand dir %s: %w", path, err)
	}

	if permFilters != 0 {
		files, err = my_slices.FilterE(files, func(f os.DirEntry) (bool, error) {
			mode, err := f.Info()
			if err != nil {
				return false, fmt.Errorf("get file info: %w", err)
			}

			return (uint32(mode.Mode().Perm()) & permFilters) != 0, nil
		})

		if err != nil {
			return nil, fmt.Errorf("couldn't apply permission filters for files: %w", err)
		}
	}

	return my_slices.Map(files, func(f os.DirEntry) string {
		return f.Name()
	}), nil
}

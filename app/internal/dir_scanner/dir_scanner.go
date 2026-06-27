package dirscanner

import (
	"errors"
	"fmt"
	"os"

	"github.com/codecrafters-io/shell-starter-go/app/internal/utils/slices"
)

const (
	EXEC_FILES_PERM = 0o111 // octal
)

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
		files, err = slices.FilterE(files, func(f os.DirEntry) (bool, error) {
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

	return slices.Map(files, func(f os.DirEntry) string {
		return f.Name()
	}), nil
}

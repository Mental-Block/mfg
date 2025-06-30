package file

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

// Exist checks whether a file with filename exists
// return true if exists, else false
func Exist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

// DirExists checks whether a directory path exists
// return true if exists, else false
func DirExists(path string) bool {
	f, err := os.Stat(path)
	return err == nil && f.IsDir()
}

// Parse tries to read json or yaml file
// and transform the content into a struct passed
// in the 2nd argument
// File extension matters, only file with extension
// json, yaml, or yml that is parsable
func Parse(filePath string, v interface{}) error {
	b, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}

	switch filepath.Ext(filePath) {
	case ".json":
		if err := json.Unmarshal(b, v); err != nil {
			return fmt.Errorf("invalid json: %w", err)
		}
	case ".yaml", ".yml":
		if err := yaml.Unmarshal(b, v); err != nil {
			return fmt.Errorf("invalid yaml: %w", err)
		}
	default:
		return errors.New("unsupported file type")
	}

	return nil
}


func CreateFile(path string, perm os.FileMode) (*os.File, error) {
    // 1. Get directory part
    dir := filepath.Dir(path)

    // 2. Ensure parent directory exists (like 'mkdir -p')
    if err := os.MkdirAll(dir, perm); err != nil {
        return nil, fmt.Errorf("creating directories %q: %w", dir, err)
    }

    // 3. Create (or truncate) the file
    f, err := os.Create(path)
	
    if err != nil {
        return nil, fmt.Errorf("creating file %q: %w", path, err)
    }

    return f, nil
}

// MoveFile attempts a fast rename, then falls back to copy+delete if needed.
func MoveFile(src, dst string) error {
    // Ensure destination directory exists
    if err := os.MkdirAll(filepath.Dir(dst), 0755); err != nil {
        return fmt.Errorf("creating destination dir: %w", err)
    }

    err := os.Rename(src, dst)
    if err == nil {
        return nil // moved in same filesystem
    }

    // Fallback for cross-filesystem moves (e.g., "invalid cross-device link")
    if !strings.Contains(err.Error(), "cross-device link") {
        return fmt.Errorf("failed to rename: %w", err)
    }

    // Perform manual copy + delete
    in, err := os.Open(src)
    if err != nil {
        return fmt.Errorf("open source: %w", err)
    }
	
    defer in.Close()

    out, err := os.OpenFile(dst, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
    if err != nil {
        return fmt.Errorf("create destination: %w", err)
    }

    defer func() {
        if cerr := out.Close(); cerr != nil && err == nil {
            err = fmt.Errorf("closing destination: %w", cerr)
        }
    }()

    if _, err := io.Copy(out, in); err != nil {
        return fmt.Errorf("copy data: %w", err)
    }

    // Optionally preserve file permissions
    if info, statErr := os.Stat(src); statErr == nil {
        _ = os.Chmod(dst, info.Mode())
    }

    if err := os.Remove(src); err != nil {
        return fmt.Errorf("remove original: %w", err)
    }

    return nil
}

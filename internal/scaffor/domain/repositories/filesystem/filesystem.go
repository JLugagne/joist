package filesystem

import (
	"context"
)

// FileSystem defines the interface for file system operations.
type FileSystem interface {
	ReadFile(ctx context.Context, path string) ([]byte, error)
	WriteFile(ctx context.Context, path string, data []byte) error
	MkdirAll(ctx context.Context, path string) error
	// ReadDir returns the base names of entries directly under path. When path
	// does not exist it returns os.ErrNotExist so callers can treat optional
	// directories (e.g. _partials/) as absent rather than fatal.
	ReadDir(ctx context.Context, path string) ([]string, error)
}

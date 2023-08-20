package server

import (
	"fmt"
	"os"
)

func getFileModifiedNano(filename string) (int64, error) {
	stat, err := os.Stat(filename)

	if err != nil {
		return 0, fmt.Errorf("os.Stat: %w", err)
	}

	return stat.ModTime().UnixNano(), nil
}

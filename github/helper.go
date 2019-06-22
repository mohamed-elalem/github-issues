package github

import "os"

func createFile(path string) (*os.File, error) {
	file, err := os.OpenFile(path,
		os.O_CREATE|os.O_TRUNC|os.O_RDWR,
		0755)

	if err != nil {
		return nil, err
	}

	return file, nil
}

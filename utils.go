package crypto

import (
	"errors"
	"os"
	"os/exec"
	"path"
)

func readFileIntoBuffer(path string, buffer []byte, expectedLength int) (err error) {
	f, err := os.Open(path)
	if err != nil {
		return
	}

	bytesCopied, err := f.Read(buffer)
	if err != nil {
	    return
	}

	if bytesCopied != expectedLength {
		err = errors.New("partial read occurred")
	}

	err = f.Close()
	if err != nil {
		return
	}

	return
}

func isFileExists(filepath string) bool {
	_, err := os.Stat(filepath)
	if os.IsNotExist(err) {
		return false
	}

	return true
}

func normalizeAndEnsureKeysPoolDir(poolPath string) (normalizedPath string, err error) {
	// By default it is expected, that keys pool path would be absolute.
	normalizedPath = poolPath

	// .. but if not - it os OK too.
	if poolPath[0] != '/' {
		dir, err := os.Getwd()
		if err != nil {
			return "", err
		}

		normalizedPath = path.Join(dir, poolPath)
	}

	// Ensure that keys pool directory exists.
	err = os.MkdirAll(normalizedPath, os.ModePerm)
	return
}

func backendExecutable() (executablePath string, err error) {
	currentDir, err := os.Getwd()
	if err != nil {
		return
	}

	executablePath = path.Join(currentDir, lamportCExecutable)
	return
}

func (scheme *Lamport) execute(args... string) error {
	command := exec.Command(scheme.executablePath, args...)
	command.Dir = scheme.keysPoolPath
	return command.Run()
}

func tryRemove(filePath string) {
	_ = os.Remove(filePath)
}
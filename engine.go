package crypto

import (
	"errors"
	"io/ioutil"
	"os"
	"os/exec"
)

const (
	keysDir = "keys/"
)

var (
	ErrKeypairAlreadyPresent = errors.New("key pair is already present")
)

type Lamport struct {
	dir  string
	name string
}

func NewLamport(operationName string) (scheme *Lamport, err error) {
	dir, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	dir += "/" + keysDir

	scheme = &Lamport{
		dir:  dir,
		name: operationName,
	}

	return
}

func (scheme *Lamport) GenerateKeypair() (err error) {
	pKeyPath := scheme.dir + scheme.name + ".pkey"
	pubKeyPath := scheme.dir + scheme.name + ".pubkey"
	if isFileExists(pKeyPath) || isFileExists(pubKeyPath) {
		err = ErrKeypairAlreadyPresent
		return
	}

	command := exec.Command("./lamportc", "generate", scheme.name)
	command.Dir = scheme.dir

	err = command.Run()
	if err != nil {
		// No files must be left in case of partial command execution.
		_ = os.Remove(pKeyPath)
		_ = os.Remove(pubKeyPath)
	}

	return
}

func (scheme *Lamport) Sign(hash Hash) (sigFilename string, err error) {
	hashFilename := scheme.name + ".bin"
	_ = ioutil.WriteFile(scheme.dir+hashFilename, hash[:], 0644)
	defer os.Remove(scheme.dir + hashFilename)

	pKeyFilename := scheme.name + ".pkey"
	sigFilename = scheme.name

	command := exec.Command("./lamportc", "sign", hashFilename, pKeyFilename, sigFilename)
	command.Dir = scheme.dir

	err = command.Run()
	if err != nil {
		// No files must be left in case of partial command execution.
		_ = os.Remove(scheme.dir + sigFilename)
	}

	return
}

func (scheme *Lamport) Verify(hash Hash) (ok bool, err error) {
	hashFilename := scheme.name + ".bin"
	_ = ioutil.WriteFile(scheme.dir+hashFilename, hash[:], 0644)
	defer os.Remove(scheme.dir + hashFilename)

	pubKeyFilename := scheme.name + ".pubkey"
	sigFilename := scheme.name + ".sig"

	command := exec.Command("./lamportc", "verify", hashFilename, sigFilename, pubKeyFilename)
	command.Dir = scheme.dir

	err = command.Run()
	if err != nil {
		err = nil
		ok = false
	}

	ok = true
	return
}

func isFileExists(filepath string) bool {
	_, err := os.Stat(filepath)
	if os.IsNotExist(err) {
		return false
	}

	return true
}

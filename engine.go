package crypto

import (
	"crypto/rand"
	"io/ioutil"
)

type Lamport struct {
	keysPoolPath   string
	executablePath string
	name           string
}

func NewLamport(operationName string, keysPoolPath string) (scheme *Lamport, err error) {
	normalizedKeysPoolPath, err := normalizeAndEnsureKeysPoolDir(keysPoolPath)
	if err != nil {
		return
	}

	executablePath, err := backendExecutable()
	if err != nil {
		return
	}

	scheme = &Lamport{
		executablePath: executablePath,
		keysPoolPath:   normalizedKeysPoolPath,
		name:           operationName,
	}

	return
}

func (scheme *Lamport) GenerateKeypair() (err error) {
	if isFileExists(scheme.pKeyFilePath()) || isFileExists(scheme.pKeyFilePath()) {
		return ErrKeypairAlreadyPresent
	}

	err = scheme.execute("generate", scheme.name)
	defer func() {
		if err != nil {
			// No files must be left in case of partial command execution.
			tryRemove(scheme.pKeyFilePath())
			tryRemove(scheme.pKeyFilePath())
		}
	}()

	return err
}

func (scheme *Lamport) Sign(hash Hash) (err error) {
	defer tryRemove(scheme.hashKeyFilePath())
	err = ioutil.WriteFile(scheme.hashKeyFilePath(), hash[:], 0644)
	if err != nil {
		return
	}

	err = scheme.execute("sign", scheme.hashFileName(), scheme.pKeyFileName(), scheme.name)
	if err != nil {
		// No files must be left in case of partial command execution.
		tryRemove(scheme.sigKeyFilePath())
		return
	}

	return
}

func (scheme *Lamport) Verify(hash Hash) (ok bool, err error) {
	defer tryRemove(scheme.hashKeyFilePath())
	err = ioutil.WriteFile(scheme.hashKeyFilePath(), hash[:], 0644)
	if err != nil {
		return
	}

	err = scheme.execute("verify", scheme.hashFileName(), scheme.sigFileName(), scheme.pubKeyFileName())
	if err != nil {
		err = nil
		ok = false
	}

	ok = true
	return
}

func GenerateRandomHash() (h Hash, err error) {
	_, err = rand.Read(h[:])
	return
}

func (scheme *Lamport) LoadPubKey() (pubKey *LamportPubKey, err error) {
	pubKey = &LamportPubKey{}
	err = readFileIntoBuffer(scheme.pubKeyFilePath(), pubKey[:], LamportPubKeySize)
	return
}

func (scheme *Lamport) LoadPrivateKey() (pKey *LamportPKey, err error) {
	pKey = &LamportPKey{}
	err = readFileIntoBuffer(scheme.pKeyFilePath(), pKey[:], LamportPKeySize)
	return
}

func (scheme *Lamport) LoadSignature() (sig *LamportSig, err error) {
	sig = &LamportSig{}
	err = readFileIntoBuffer(scheme.sigKeyFilePath(), sig[:], LamportSigSize)
	return
}

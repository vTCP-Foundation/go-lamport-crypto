package crypto

import "path"

func (scheme *Lamport) pKeyFileName() string {
	return scheme.name + ".pkey"
}

func (scheme *Lamport) pubKeyFileName() string {
	return scheme.name + ".pubkey"
}

func (scheme *Lamport) sigFileName() string {
	return scheme.name + ".sig"
}

func (scheme *Lamport) hashFileName() string {
	return scheme.name + ".bin"
}

func (scheme *Lamport) pKeyFilePath() string {
	return path.Join(scheme.keysPoolPath, scheme.pKeyFileName())
}

func (scheme *Lamport) pubKeyFilePath() string {
	return path.Join(scheme.keysPoolPath, scheme.pubKeyFileName())
}

func (scheme *Lamport) sigKeyFilePath() string {
	return path.Join(scheme.keysPoolPath, scheme.sigFileName())
}

func (scheme *Lamport) hashKeyFilePath() string {
	return path.Join(scheme.keysPoolPath, scheme.hashFileName())
}

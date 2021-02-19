package crypto

const (
	LamportPKeySize   = 1024 * 16
	LamportPubKeySize = 1024 * 16
	LamportSigSize    = 1024 * 8
)

type Hash [32]byte

type LamportSig = [LamportSigSize]byte

type LamportPKey = [LamportPKeySize]byte

type LamportPubKey = [LamportPubKeySize]byte

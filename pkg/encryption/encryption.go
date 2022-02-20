package encryption

type (
	Encryption struct {
		aesKey     string
		privateKey string
		publicKey  string
	}
)

func New(aesKey, privKey, pubKey string) Encryption {
	return Encryption{
		aesKey:     aesKey,
		privateKey: privKey,
		publicKey:  pubKey,
	}
}

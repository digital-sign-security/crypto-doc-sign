package generators

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"github.com/crypto-sign/cmd/backend/configuration"
	"golang.org/x/crypto/ssh"
	"log"
)

type KeysGenerator struct {
	masterKey string
	bitSize   int
}

func NewKeysGenerator(cfg configuration.KeyGenerator) *KeysGenerator {
	return &KeysGenerator{
		masterKey: cfg.MasterKey, // TODO: delete it (maybe?) useless
		bitSize:   cfg.BitSize,
	}
}

func (k *KeysGenerator) generatePrivateKey() (*rsa.PrivateKey, error) {
	// Private Key generation
	privateKey, err := rsa.GenerateKey(rand.Reader, k.bitSize)
	if err != nil {
		return nil, err
	}

	// Validate Private Key
	err = privateKey.Validate()
	if err != nil {
		return nil, err
	}

	log.Println("Private Key generated")
	return privateKey, nil
}

func (k *KeysGenerator) generatePublicKey(privatekey *rsa.PublicKey) ([]byte, error) {
	publicRsaKey, err := ssh.NewPublicKey(privatekey)
	if err != nil {
		return nil, err
	}

	pubKeyBytes := ssh.MarshalAuthorizedKey(publicRsaKey)

	log.Println("Public key generated")
	return pubKeyBytes, nil
}

func (k *KeysGenerator) GenerateKeys() ([]byte, []byte) {
	privateKey, err := k.generatePrivateKey()
	if err != nil {
		log.Fatal(err.Error())
	}

	publicKeyBytes, err := k.generatePublicKey(&privateKey.PublicKey)
	if err != nil {
		log.Fatal(err.Error())
	}

	privateKeyBytes := encodePrivateKeyToPEM(privateKey)
	return publicKeyBytes, privateKeyBytes
}

// encodePrivateKeyToPEM encodes Private Key from RSA to PEM format
func encodePrivateKeyToPEM(privateKey *rsa.PrivateKey) []byte {
	// Get ASN.1 DER format
	privDER := x509.MarshalPKCS1PrivateKey(privateKey)

	// pem.Block
	privBlock := pem.Block{
		Type:    "RSA PRIVATE KEY",
		Headers: nil,
		Bytes:   privDER,
	}

	// Private key in PEM format
	privatePEM := pem.EncodeToMemory(&privBlock)

	return privatePEM
}

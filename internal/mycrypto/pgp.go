package mycrypto

import (
	"fmt"
	"github.com/pkg/errors"
	"log"

	"github.com/ProtonMail/gopenpgp/v2/crypto"
)

// combining PGP (Pretty Good Privacy) for key management with asymmetric encryption
// and AES (Advanced Encryption Standard) for symmetric message encryption
// is a common approach to achieve both the efficiency and security

// see https://github.com/ProtonMail/gopenpgp/blob/main/crypto/sessionkey_test.go

// https://pgp.mit.edu/pks/lookup?search=niqdev&op=index
// https://pgp.mit.edu/pks/lookup?op=get&search=0x42A644BE82ADB4F9

// https://www.comparitech.com/blog/information-security/diffie-hellman-key-exchange
// https://www.comparitech.com/blog/information-security/rsa-encryption
// https://proton.me/blog/what-is-pgp-encryption
// https://security.stackexchange.com/questions/406/how-should-i-distribute-my-public-key

func PgpMessage() {

	// generate PGP key pairs for each participant (sender and recipient)
	// keep the private keys secure and distribute the public keys to the participants
	alicePrivateKey, alicePublicKey, err := generateKeyPair("Alice", "alice@example.com")
	if err != nil {
		log.Fatalln("error generating alice keys")
	}
	fmt.Println("Alice PRIVATE key:\n", alicePrivateKey)
	fmt.Println("Alice PUBLIC key:\n", alicePublicKey)

	bobPrivateKey, bobPublicKey, err := generateKeyPair("Bob", "bob@example.com")
	if err != nil {
		log.Fatalln("error generating bob keys")
	}
	fmt.Println("Bob PRIVATE key:\n", bobPrivateKey)
	fmt.Println("Bob PUBLIC key:\n", bobPublicKey)

	encryptedMessage, err := encryptMessage("MyS3cr3t", bobPublicKey)
	if err != nil {
		log.Fatalln("error encrypting message")
	}
	fmt.Println("Encrypted message with Bob's public key:\n", encryptedMessage)

	decryptedMessage, err := decryptMessage(encryptedMessage, bobPrivateKey)
	if err != nil {
		log.Fatalln("error decrypting message")
	}
	fmt.Println("Decrypted message with Bob's private key:\n", decryptedMessage)
}

func generateKeyPair(name string, email string) (string, string, error) {
	// TODO use helper.GenerateKey with passphrase

	// using Curve25519, alternative rsa/2048
	ecPrivateKey, err := crypto.GenerateKey(name, email, "x25519", 0)
	if err != nil {
		return "", "", errors.Wrapf(err, "error generating key")
	}

	// an armored key is a key that has been encoded into a human-readable, text-based format i.e. ASCII
	// each armored key block starts with a "BEGIN" line and ends with an "END" line,
	// with the actual key data encoded in between
	privateKey, err := ecPrivateKey.Armor()
	if err != nil {
		return "", "", errors.Wrapf(err, "error private key")
	}

	// the public key is derived from the private key
	publicKey, err := ecPrivateKey.GetArmoredPublicKey()
	if err != nil {
		return "", "", errors.Wrapf(err, "error public key")
	}

	return privateKey, publicKey, nil
}

func encryptMessage(plaintext string, publicKey string) (string, error) {
	// for each message, generate a random symmetric AES session key
	sessionKey, err := crypto.GenerateSessionKey()
	if err != nil {
		return "", errors.Wrapf(err, "error generating session key")
	}

	// encrypt the message with AES (symmetric) using the session key
	message := crypto.NewPlainMessageFromString(plaintext)
	dataPacket, err := sessionKey.Encrypt(message)
	if err != nil {
		return "", errors.Wrapf(err, "error encrypting message")
	}

	// encrypt the AES key using the recipient's public PGP key
	// to ensures that only the recipient, who possesses the corresponding private PGP key,
	// can decrypt the AES key
	publicKeyObj, err := crypto.NewKeyFromArmored(publicKey)
	if err != nil {
		return "", errors.Wrapf(err, "error reading public key")
	}
	keyRing, err := crypto.NewKeyRing(publicKeyObj)
	if err != nil {
		return "", errors.Wrapf(err, "error creating public key ring")
	}
	keyPacket, err := keyRing.EncryptSessionKey(sessionKey)
	if err != nil {
		return "", errors.Wrapf(err, "error encrypting session key")
	}

	// combine and send the encrypted AES session key and the AES-encrypted message to the recipient
	splitMessage := crypto.NewPGPSplitMessage(keyPacket, dataPacket)
	armored, err := splitMessage.GetArmored()
	if err != nil {
		return "", errors.Wrapf(err, "error to armor pgp message")
	}

	return armored, nil
}

func decryptMessage(encryptedMessage string, privateKey string) (string, error) {
	pgpMessage, err := crypto.NewPGPMessageFromArmored(encryptedMessage)
	if err != nil {
		return "", errors.Wrapf(err, "error to unarmor pgp message")
	}

	// ids, ok := pgpMessage.GetEncryptionKeyIDs()

	// 1) split the message into the encrypted session key and the AES-encrypted message
	// 2) decrypt the session key with the recipient private key
	// 3) decrypt the message with the session key using AES

	privateKeyObj, err := crypto.NewKeyFromArmored(privateKey)
	if err != nil {
		return "", errors.Wrapf(err, "error reading private key")
	}
	keyRing, err := crypto.NewKeyRing(privateKeyObj)
	if err != nil {
		return "", errors.Wrapf(err, "error creating private key ring")
	}
	message, err := keyRing.Decrypt(pgpMessage, nil, 0)
	if err != nil {
		return "", errors.Wrapf(err, "error decrypting message")
	}

	plaintext := message.GetString()

	return plaintext, nil
}

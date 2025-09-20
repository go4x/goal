// Package ciphers provides cryptographic functions for encryption, decryption, and data compression.
//
// This package includes:
//   - AES encryption/decryption with ECB and CBC modes
//   - Base36, Base62, and full character set compression functions
//
// SECURITY WARNINGS:
//   - ECB mode is NOT secure and should NOT be used in production
//   - Always use CBC mode with a random IV for secure encryption
//   - Ensure keys are cryptographically random and properly sized
//
// Example usage:
//
//	// Secure encryption (recommended)
//	encrypted, err := ciphers.AES.Encrypt(data, key, ciphers.CBC, randomIV)
//	decrypted, err := ciphers.AES.Decrypt(encrypted, key, ciphers.CBC, randomIV)
//
//	// Data compression
//	compressed := ciphers.C62(12345) // "3d7"
package ciphers

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"

	"github.com/gophero/goal/errorx"
)

// AES encrypt and decrypt object
//
// https://en.wikipedia.org/wiki/Advanced_Encryption_Standard
var AES = &aeser{}

type aeser struct {
}

// AESMode is the mode of the AES encryption and decryption
type AESMode int

const (
	// ECB mode: https://en.wikipedia.org/wiki/Block_cipher_mode_of_operation#ECB
	// WARNING: ECB mode has serious security vulnerabilities and should NOT be used for production.
	// ECB mode is vulnerable to pattern analysis attacks and does not provide semantic security.
	// See: https://crypto.stackexchange.com/questions/20941/why-shouldnt-i-use-ecb-encryption/20946#20946
	// Use CBC mode with a random IV instead.
	ECB AESMode = iota
	// CBC mode: https://en.wikipedia.org/wiki/Block_cipher_mode_of_operation#CBC
	// CBC mode is more secure than ECB but requires a random IV for each encryption.
	// The IV must be 16 bytes long and should be cryptographically random.
	CBC
)

func (a *aeser) Encrypt(rawBytes []byte, key []byte, mode AESMode, iv []byte) ([]byte, error) {
	// Validate key length
	if len(key) != 16 && len(key) != 24 && len(key) != 32 {
		return nil, errorx.New("invalid key length: %d bytes, must be 16, 24, or 32 bytes", len(key))
	}

	switch mode {
	case ECB:
		return aesEncryptECB(rawBytes, key)
	case CBC:
		// Validate IV length for CBC mode
		if len(iv) != 16 {
			return nil, errorx.New("invalid IV length: %d bytes, must be 16 bytes for AES", len(iv))
		}
		return aesEncryptCBC(rawBytes, key, iv)
	default:
		return nil, errorx.New("unsupported encrypt mode: %v", mode)
	}
}

func (a *aeser) Decrypt(cipherBytes []byte, key []byte, mode AESMode, iv []byte) ([]byte, error) {
	// Validate key length
	if len(key) != 16 && len(key) != 24 && len(key) != 32 {
		return nil, errorx.New("invalid key length: %d bytes, must be 16, 24, or 32 bytes", len(key))
	}

	switch mode {
	case ECB:
		return aesDecryptECB(cipherBytes, key)
	case CBC:
		// Validate IV length for CBC mode
		if len(iv) != 16 {
			return nil, errorx.New("invalid IV length: %d bytes, must be 16 bytes for AES", len(iv))
		}
		return aesDecryptCBC(cipherBytes, key, iv)
	default:
		return nil, errorx.New("unsupported decrypt mode: %v", mode)
	}
}

// pkcs7Padding pkcs7 padding, see: https://en.wikipedia.org/wiki/Padding_(cryptography)
func pkcs7Padding(cipherText []byte, blockSize int) []byte {
	// calculate the padding length, the minimum is 1, the maximum is blockSize
	padding := blockSize - len(cipherText)%blockSize
	// copy padding bytes to the end of the cipherText
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	// append the padding bytes to the end of the cipherText
	return append(cipherText, padText...)
}

// pkcs7UnPadding pkcs7 unpadding
func pkcs7UnPadding(cipherText []byte) ([]byte, error) {
	length := len(cipherText)
	if length == 0 {
		return nil, errorx.New("empty ciphertext")
	}

	// get the padding length, the last byte of the cipherText represents the number of padding bytes
	unPadding := int(cipherText[length-1])

	// validate padding - check for reasonable bounds
	if unPadding == 0 || unPadding > length || unPadding > 16 {
		return nil, errorx.New("invalid padding: padding length %d is invalid", unPadding)
	}

	// check if all padding bytes are correct
	for i := length - unPadding; i < length; i++ {
		if cipherText[i] != byte(unPadding) {
			return nil, errorx.New("invalid padding: padding bytes are inconsistent")
		}
	}

	// remove the padding bytes
	return cipherText[:(length - unPadding)], nil
}

func aesEncryptECB(rawText []byte, key []byte) ([]byte, error) {
	// create cipher, if the key length is not 16, 24, or 32 bytes, it will panic
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	// AES block size is 16 bytes, 128 bits, so blockSize = 16 bytes
	bs := block.BlockSize()
	// use pkcs#7 padding mode
	rawText = pkcs7Padding(rawText, bs)
	// the length of the encrypted bytes array must be a multiple of the block size, that is, 16
	if len(rawText)%bs != 0 {
		return nil, errorx.New("block size padding failed")
	}

	out := make([]byte, len(rawText))
	dst := out
	// encrypt the raw text by blocks
	for len(rawText) > 0 {
		block.Encrypt(dst, rawText[:bs])
		rawText = rawText[bs:]
		dst = dst[bs:]
	}
	return out, nil
}

func aesDecryptECB(cipherText []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	bs := block.BlockSize()
	if len(cipherText)%bs != 0 {
		return nil, errorx.New("illegal ciphertext length")
	}

	out := make([]byte, len(cipherText))
	dst := out
	// decrypt the cipher text by blocks
	for len(cipherText) > 0 {
		block.Decrypt(dst, cipherText[:bs])
		cipherText = cipherText[bs:]
		dst = dst[bs:]
	}
	out, err = pkcs7UnPadding(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func aesEncryptCBC(rawBytes []byte, key []byte, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	rawBytes = pkcs7Padding(rawBytes, blockSize)
	// create CBC encryptor, the length of the initial vector iv must be equal to the block size
	blockMode := cipher.NewCBCEncrypter(block, iv) // the length of the initial vector iv must be equal to the block size
	dst := make([]byte, len(rawBytes))
	blockMode.CryptBlocks(dst, rawBytes)
	return dst, nil
}

func aesDecryptCBC(cipherBytes []byte, key []byte, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockMode := cipher.NewCBCDecrypter(block, iv) // the length of the initial vector iv must be equal to the block size
	rawBytes := make([]byte, len(cipherBytes))
	blockMode.CryptBlocks(rawBytes, cipherBytes)
	rawBytes, err = pkcs7UnPadding(rawBytes)
	if err != nil {
		return nil, err
	}
	return rawBytes, nil
}

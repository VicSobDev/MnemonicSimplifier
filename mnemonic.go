package mnemonic_simplifier

import (
	"encoding/hex"
	"fmt"

	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
)

// Wallet struct holds a slice of accounts.
type Wallet struct {
	Accounts [10]Account
}

// Account struct represents an Ethereum account with its address, public key, and private key.
type Account struct {
	Address    string
	PublicKey  string
	PrivateKey string
}

// NewWallet creates a new Wallet.
// It returns a Wallet and any error encountered.
func NewWallet(mnemonic string) (Wallet, error) {
	wallet := Wallet{}
	err := wallet.init(mnemonic)
	if err != nil {
		return Wallet{}, err
	}
	return wallet, nil
}

// init initializes the Wallet with accounts derived from the given mnemonic.
// It returns an error if any operation fails.
func (w *Wallet) init(mnemonic string) error {
	// Create a new wallet instance from the mnemonic.
	wallet, err := hdwallet.NewFromMnemonic(mnemonic)
	if err != nil {
		return fmt.Errorf("failed to create wallet from mnemonic: %w", err)
	}

	// Derive the accounts.
	for i := 0; i < 10; i++ {
		path := fmt.Sprintf("m/44'/60'/0'/0/%d", i)
		derivationPath, err := hdwallet.ParseDerivationPath(path)
		if err != nil {
			return fmt.Errorf("failed to parse derivation path: %w", err)
		}

		account, err := wallet.Derive(derivationPath, false)
		if err != nil {
			return fmt.Errorf("failed to derive account: %w", err)
		}
		w.Accounts[i].Address = account.Address.Hex()

		publicKeyECDSA, err := wallet.PublicKey(account)
		if err != nil {
			return fmt.Errorf("failed to get public key: %w", err)
		}

		// Determine the prefix based on the Y coordinate's parity.
		var prefix byte
		if publicKeyECDSA.Y.Bit(0) == 0 {
			prefix = 0x02 // Even Y coordinate
		} else {
			prefix = 0x03 // Odd Y coordinate
		}

		// Prepare the compressed public key in hexadecimal format.
		publicKeyBytes := append([]byte{prefix}, publicKeyECDSA.X.Bytes()...)
		w.Accounts[i].PublicKey = hex.EncodeToString(publicKeyBytes)

		privateKey, err := wallet.PrivateKey(account)
		if err != nil {
			return fmt.Errorf("failed to get private key: %w", err)
		}
		w.Accounts[i].PrivateKey = fmt.Sprintf("%x", privateKey.D)
	}
	return nil
}

# Mnemonic Simplifier

## Overview

Mnemonic Simplifier is a Go package designed to facilitate the creation and management of Ethereum wallets based on a mnemonic passphrase. It uses the `go-ethereum-hdwallet` library to generate wallets with ten accounts, including addresses, public keys, and private keys.

## Features

- Generate a wallet from a mnemonic.
- Retrieve Ethereum account details including address, public key, and private key for each account.
- Supports derivation of the first 10 accounts under the standard Ethereum derivation path.

## Installation

To use the Mnemonic Simplifier in your project, ensure you have Go installed, then follow these steps:

1. Clone the repository or download the package:
   ```bash
   go get github.com/yourusername/mnemonic_simplifier
   ```
2. Import the package in your Go project:
   ```go
   import (
       "github.com/yourusername/mnemonic_simplifier"
   )
   ```

## Usage

To create a new wallet from a mnemonic:

```go
package main

import (
    "fmt"
    "github.com/yourusername/mnemonic_simplifier"
)

func main() {
    mnemonic := "abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon about"
    wallet, err := mnemonic_simplifier.NewWallet(mnemonic)
    if err != nil {
        fmt.Printf("Error: %s\n", err)
        return
    }

    for _, account := range wallet.Accounts {
        fmt.Printf("Address: %s\nPublic Key: %s\nPrivate Key: %s\n", account.Address, account.PublicKey, account.PrivateKey)
    }
}
```

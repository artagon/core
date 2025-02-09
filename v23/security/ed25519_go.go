// Copyright 2020 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !openssl

package security

import "crypto/ed25519"

func newInMemoryED25519SignerImpl(key ed25519.PrivateKey, hash Hash) (Signer, error) {
	return newGoStdlibED25519Signer(key, hash)
}

func newED25519PublicKeyImpl(key ed25519.PublicKey, hash Hash) PublicKey {
	return newGoStdlibED25519PublicKey(key, hash)
}

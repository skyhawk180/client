// Copyright 2015 Keybase, Inc. All rights reserved. Use of
// this source code is governed by the included BSD license.

package libkb

import (
	"bytes"
	"strings"
	"testing"

	"github.com/keybase/client/go/saltpack"
)

type outputBuffer struct {
	bytes.Buffer
}

func (ob outputBuffer) Close() error {
	return nil
}

// Encrypt a message, and make sure recipients can decode it, and
// non-recipients can't decode it.
func TestSaltpackEncDec(t *testing.T) {
	senderKP, err := GenerateNaclDHKeyPair()
	if err != nil {
		t.Fatal(err)
	}

	var receiverKPs []NaclDHKeyPair
	var receiverPKs []NaclDHKeyPublic
	for i := 0; i < 12; i++ {
		kp, err := GenerateNaclDHKeyPair()
		if err != nil {
			t.Fatal(err)
		}
		receiverKPs = append(receiverKPs, kp)
		receiverPKs = append(receiverPKs, kp.Public)
	}

	nonReceiverKP, err := GenerateNaclDHKeyPair()
	if err != nil {
		t.Fatal(err)
	}

	message := "The Magic Words are Squeamish Ossifrage"

	var buf outputBuffer

	err = SaltPackEncrypt(
		strings.NewReader(message), &buf, receiverPKs, senderKP)
	if err != nil {
		t.Fatal(err)
	}

	ciphertext := buf.String()
	if !strings.HasPrefix(ciphertext, saltpack.EncryptionArmorHeader) {
		t.Errorf("ciphertext doesn't have header: %s", ciphertext)
	}

	if !strings.HasSuffix(ciphertext, saltpack.EncryptionArmorFooter+".\n") {
		t.Errorf("ciphertext doesn't have footer: %s", ciphertext)
	}

	for _, key := range receiverKPs {
		buf.Reset()
		err = SaltPackDecrypt(
			strings.NewReader(ciphertext),
			&buf, key)
		if err != nil {
			t.Fatal(err)
		}

		plaintext := buf.String()
		if plaintext != message {
			t.Errorf("expected %s, got %s",
				message, plaintext)
		}
	}

	// Sender is a non-recipient, too.
	nonReceiverKPs := []NaclDHKeyPair{nonReceiverKP, senderKP}

	for _, kp := range nonReceiverKPs {
		buf.Reset()
		err = SaltPackDecrypt(
			strings.NewReader(ciphertext), &buf, kp)
		if err != saltpack.ErrNoDecryptionKey {
			t.Fatal(err)
		}
	}
}
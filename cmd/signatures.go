// Copyright 2017-2018 Stratumn SAS. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"context"
	"encoding/base64"

	"github.com/pkg/errors"
	"github.com/stratumn/go-chainscript"
	"github.com/stratumn/go-crypto/keys"
)

// SignaturesTest tests a segment with signatures.
type SignaturesTest struct{}

// NewSignaturesTest creates the test case.
func NewSignaturesTest() TestCase {
	return &SignaturesTest{}
}

// Generate encoded segment bytes.
func (t *SignaturesTest) Generate() string {
	link, err := chainscript.NewLinkBuilder("test_process", "test_map").
		WithAction("ʙᴀᴛᴍᴀɴ").
		Build()
	if err != nil {
		panic(err)
	}

	_, ed25519Key, err := keys.NewEd25519KeyPair()
	if err != nil {
		panic(err)
	}

	ed25519KeyBytes, err := keys.EncodeED25519SecretKey(ed25519Key)
	if err != nil {
		panic(err)
	}

	err = link.Sign(ed25519KeyBytes, "")
	if err != nil {
		panic(err)
	}

	_, rsaKey, err := keys.NewRSAKeyPair()
	if err != nil {
		panic(err)
	}

	rsaKeyBytes, err := keys.EncodeRSASecretKey(rsaKey)
	if err != nil {
		panic(err)
	}

	err = link.Sign(rsaKeyBytes, "[version,meta.mapId]")
	if err != nil {
		panic(err)
	}

	segment, err := link.Segmentify()
	if err != nil {
		panic(err)
	}

	b, err := chainscript.MarshalSegment(segment)
	if err != nil {
		panic(err)
	}

	return base64.StdEncoding.EncodeToString(b)
}

// Validate encoded segment bytes.
func (t *SignaturesTest) Validate(encoded string) error {
	b, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		return err
	}

	segment, err := chainscript.UnmarshalSegment(b)
	if err != nil {
		return err
	}

	err = segment.Validate(context.Background())
	if err != nil {
		return err
	}

	if len(segment.Link.Signatures) != 2 {
		return errors.Errorf("invalid number of signatures: %d", len(segment.Link.Signatures))
	}

	err = segment.Link.Signatures[0].Validate(segment.Link)
	if err != nil {
		return err
	}

	err = segment.Link.Signatures[1].Validate(segment.Link)
	if err != nil {
		return err
	}

	if segment.Link.Signatures[0].PayloadPath != "[version,data,meta]" {
		return errors.Errorf("invalid first signature payload path: %s", segment.Link.Signatures[0].PayloadPath)
	}

	if segment.Link.Signatures[1].PayloadPath != "[version,meta.mapId]" {
		return errors.Errorf("invalid second signature payload path: %s", segment.Link.Signatures[1].PayloadPath)
	}

	return nil
}

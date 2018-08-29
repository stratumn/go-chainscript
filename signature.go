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

package chainscript

import (
	"crypto/sha256"

	json "github.com/gibson042/canonicaljson-go"
	"github.com/jmespath/go-jmespath"
	"github.com/pkg/errors"
	"github.com/stratumn/go-crypto/signatures"
)

const (
	// SignatureVersion1_0_0 is the first version of the link signature.
	// In that version we use canonical JSON to encode the link parts.
	// We use JMESPATH to select what parts of the link need to be signed.
	// We use SHA-256 on the JSON-encoded bytes and sign the resulting hash.
	// We use github.com/stratumn/go-crypto's 1.0.0 release to produce the
	// signature (which uses PEM-encoded private keys).
	SignatureVersion1_0_0 = "1.0.0"

	// SignatureVersion is the version used for new signatures.
	SignatureVersion = SignatureVersion1_0_0
)

// Signature errors.
var (
	ErrUnknownSignatureVersion = errors.New("unknown signature version")
	ErrInvalidSignature        = errors.New("signature is invalid")
)

// Sign configurable parts of the link with the current signature version.
// The payloadPath is used to select what parts of the link need to be signed
// with the given private key. If no payloadPath is provided, the whole link
// is signed.
func (l *Link) Sign(privateKey []byte, payloadPath string) error {
	if len(payloadPath) == 0 {
		payloadPath = "[version,data,meta]"
	}

	payload, err := l.SignedBytes(SignatureVersion, payloadPath)
	if err != nil {
		return err
	}

	sig, err := signatures.Sign(privateKey, payload)
	if err != nil {
		return errors.WithStack(err)
	}

	s := &Signature{
		Version:     SignatureVersion,
		PayloadPath: payloadPath,
		PublicKey:   sig.PublicKey,
		Signature:   sig.Signature,
	}

	l.Signatures = append(l.Signatures, s)
	return nil
}

// SignedBytes computes the bytes that should be signed.
// The signature version impacts how those bytes are computed.
func (l *Link) SignedBytes(sigVersion, payloadPath string) ([]byte, error) {
	switch sigVersion {
	case SignatureVersion1_0_0:
		if len(payloadPath) == 0 {
			payloadPath = "[version,data,meta]"
		}

		payload, err := jmespath.Search(payloadPath, l)
		if err != nil {
			return nil, errors.WithStack(err)
		}

		payloadBytes, err := json.Marshal(payload)
		if err != nil {
			return nil, errors.WithStack(err)
		}

		h := sha256.Sum256(payloadBytes)
		return h[:], nil
	default:
		return nil, ErrUnknownSignatureVersion
	}
}

// Validate the signature.
func (s *Signature) Validate(l *Link) error {
	signedBytes, err := l.SignedBytes(s.Version, s.PayloadPath)
	if err != nil {
		return err
	}

	if err := l.compatible(); err != nil {
		return err
	}

	switch s.Version {
	case SignatureVersion1_0_0:
		sig := signatures.Signature{
			Message:   signedBytes,
			PublicKey: s.PublicKey,
			Signature: s.Signature,
		}

		if err := signatures.Verify(&sig); err != nil {
			return errors.Wrap(ErrInvalidSignature, err.Error())
		}
	default:
		return ErrUnknownSignatureVersion
	}

	return nil
}

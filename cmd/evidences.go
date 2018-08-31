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
	"bytes"
	"context"
	"encoding/base64"
	"fmt"

	"github.com/golang/protobuf/proto"
	"github.com/stratumn/go-chainscript"
)

// EvidencesTest tests a segment with evidences.
type EvidencesTest struct{}

// NewEvidencesTest creates the test case.
func NewEvidencesTest() TestCase {
	return &EvidencesTest{}
}

// Generate encoded segment bytes.
func (t *EvidencesTest) Generate() string {
	link, err := chainscript.NewLinkBuilder("test_process", "test_map").
		Build()
	if err != nil {
		panic(err)
	}

	segment, err := link.Segmentify()
	if err != nil {
		panic(err)
	}

	err = segment.AddEvidence(&chainscript.Evidence{
		Version:  "0.1.0",
		Backend:  "bitcoin",
		Provider: "testnet",
		Proof:    []byte{42},
	})
	if err != nil {
		panic(err)
	}

	err = segment.AddEvidence(&chainscript.Evidence{
		Version:  "1.0.3",
		Backend:  "ethereum",
		Provider: "mainnet",
		Proof:    []byte{24},
	})
	if err != nil {
		panic(err)
	}

	b, err := proto.Marshal(segment)
	if err != nil {
		panic(err)
	}

	return base64.StdEncoding.EncodeToString(b)
}

// Validate encoded segment bytes.
func (t *EvidencesTest) Validate(encoded string) error {
	b, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		return err
	}

	var segment chainscript.Segment
	err = proto.Unmarshal(b, &segment)
	if err != nil {
		return err
	}

	err = segment.Validate(context.Background(), nil)
	if err != nil {
		return err
	}

	if len(segment.Meta.Evidences) != 2 {
		return fmt.Errorf("invalid evidences count: %d", len(segment.Meta.Evidences))
	}

	btc := segment.GetEvidence("bitcoin", "testnet")
	if btc == nil {
		return fmt.Errorf("missing bitcoin evidence")
	}
	if btc.Version != "0.1.0" || btc.Backend != "bitcoin" || btc.Provider != "testnet" || !bytes.Equal(btc.Proof, []byte{42}) {
		return fmt.Errorf("invalid bitcoin evidence: %v", btc)
	}

	eth := segment.GetEvidence("ethereum", "mainnet")
	if eth == nil {
		return fmt.Errorf("missing ethereum evidence")
	}
	if eth.Version != "1.0.3" || eth.Backend != "ethereum" || eth.Provider != "mainnet" || !bytes.Equal(eth.Proof, []byte{24}) {
		return fmt.Errorf("invalid ethereum evidence: %v", eth)
	}

	return nil
}

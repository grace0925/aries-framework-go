/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package edv

import (
	"encoding/json"
)

// StructuredDocument represents a Structured Document
// as defined in https://identity.foundation/secure-data-store/#structureddocument.
type StructuredDocument struct {
	ID      string                 `json:"id"`
	Meta    map[string]interface{} `json:"meta"`
	Content map[string]interface{} `json:"content"`
}

// EncryptedDocument represents an Encrypted Document as defined in
// https://identity.foundation/secure-data-store/#encrypteddocument.
type EncryptedDocument struct {
	ID                          string                       `json:"id"`
	Sequence                    int                          `json:"sequence"`
	IndexedAttributeCollections []IndexedAttributeCollection `json:"indexed,omitempty"`
	JWE                         json.RawMessage              `json:"jwe"`
}

// IndexedAttributeCollection represents a collection of indexed attributes,
// all of which share a common MAC algorithm and key.
// This format is based on https://identity.foundation/secure-data-store/#creating-encrypted-indexes.
type IndexedAttributeCollection struct {
	Sequence          int                `json:"sequence"`
	HMAC              IDTypePair         `json:"hmac"`
	IndexedAttributes []IndexedAttribute `json:"attributes"`
}

// IndexedAttribute represents a single indexed attribute.
type IndexedAttribute struct {
	Name   string `json:"name"`
	Value  string `json:"value"`
	Unique bool   `json:"unique"`
}

// IDTypePair represents an ID+Type pair.
// TODO: #2262 This is a simplified version of the actual EDV query format, which is still not finalized
// in the spec as of writing. See: https://github.com/decentralized-identity/secure-data-store/issues/34.
type IDTypePair struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}

// Query represents a name+value pair that can be used to query the encrypted indices for specific data.
// TODO: #2262 This is a simplified version of the actual EDV query format, which is still not finalized
// in the spec as of writing. See: https://github.com/decentralized-identity/secure-data-store/issues/34.
type Query struct {
	Name  string `json:"index"`
	Value string `json:"equals"`
}

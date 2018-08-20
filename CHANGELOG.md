# Changelog

Each version of the ChainScript implementation makes specific serialization
choices. Those choices are detailed in this document.

## 1.0.0

- Dependencies:
  - [Canonical JSON](https://github.com/gibson042/canonicaljson-go) v1.0.3 (81f5327eb8367be6f106cd62f136b6f25b4c1678)
  - [Protobuf](https://github.com/golang/protobuf) v1.1.0 (b4deda0973fb4c70b50d226b1af49f3da59f5265)
  - [Crypto](https://github.com/stratumn/go-crypto) v1.0.0
- Data bytes (e.g. _link.data_, _link.meta.data_) are encoded from Go objects
  using canonical JSON
- Link hash is calculated by hashing the protobuf-encoded link bytes with
  SHA-256
- All links are created with the following ClientID:
  github.com/stratumn/go-chainscript
- Signatures provide flexibility over what parts of the link are signed using
  the link's canonical JSON representation and JMESPATH. A SHA-256 hash of the
  JSON bytes is signed with github.com/stratumn/go-crypto.

# ChainScript

[![build status](https://travis-ci.org/stratumn/chainscript.svg?branch=master)](https://travis-ci.org/stratumn/chainscript)

ChainScript is an open standard for representing Proof of Process data.

Proof of Process is a protocol that allows partners to follow the execution of
a shared process.
Proof of Process provides immutability and auditability of every step in the
process.

Learn more [here](https://proofofprocess.org/).

## How to use

ChainScript is defined with [protobuf](https://developers.google.com/protocol-buffers/)
to provide a portable encoding.

It is recommended to import this repository as a `git subtree`:

```bash
git subtree add --prefix proto git@github.com:stratumn/chainscript.git master --squash
```

Your application can then compile the protobuf file to types in your language
of choice and add any helpers you need.

Changes to ChainScript should be done in the current repository: clients should
pull the latest changes with `git subtree`.

## Reference Implementations

Stratumn provides opinionated reference implementations that should suit most
projects:

| Language   | Repository                                 | Status      |
| ---------- | ------------------------------------------ | ----------- |
| Golang     | https://github.com/stratumn/go-chainscript | Development |
| Javascript | https://github.com/stratumn/js-chainscript | Development |

## Design Choices

### Stable Encoding

ChainScript needs a stable encoding as it contains signatures and hashes of
parts of the data. If a link didn't have the same hash in different languages,
cross-platform applications wouldn't work.

ChainScript only uses protobuf types that have a canonical encoding.
That means that in particular, protobuf maps can't be used as their ordering is
not specified by the protobuf specification and can vary depending on the
implementation.

Dynamic data (such as a link's state) is represented as bytes. It is your
application's responsibility to use a stable byte encoding for that data.

You have several options to define a stable encoding (not exhaustive):

- [protobuf](https://developers.google.com/protocol-buffers/)
- [canonical JSON](https://gibson042.github.io/canonicaljson-spec/)
- [amino](https://github.com/tendermint/go-amino)
- custom binary encoding

### Versioning

ChainScript should follow [protobuf's guidelines](https://developers.google.com/protocol-buffers/docs/proto3#updating)
for backwards-compatible updates.

ChainScript provides multiple version fields. These fields should be used by
applications to update the dynamic parts of their data (such as a link's state)
in a backwards-compatible way.

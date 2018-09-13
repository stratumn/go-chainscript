# Test Samples

We provide a description of the contents of each serialized segment sample.
ChainScript implementation should test that they're able to deserialize these
segments properly and validate their content.

## 1.0.0

**simple-segment**: a segment containing data and metadata but no references,
evidences or signatures.

Note: we replaced `link.data` and `link.meta.data` with the objects that they
represent in the following JSON snippet.
The real ChainScript contains only serialized bytes in those fields, but it
should be tested that the values can be correctly extracted.

```json
{
  "link": {
    "version": "1.0.0",
    "data": {
      "name": "ʙᴀᴛᴍᴀɴ",
      "age": 42
    },
    "meta": {
      "clientId": "github.com/stratumn/go-chainscript",
      "prevLinkHash": "Kio=",
      "priority": 42,
      "process": { "name": "test_process", "state": "started" },
      "mapId": "test_map",
      "action": "init",
      "step": "setup",
      "tags": ["tag1", "tag2"],
      "data": "bruce wayne"
    }
  },
  "meta": { "linkHash": "/oLiUbxvh/LOhYC+61GM/gHfWQjnRCXsL8tKkTT6sVM=" }
}
```

**segment-references**: a segment with two references to other segments.

```json
{
  "link": {
    "version": "1.0.0",
    "meta": {
      "clientId": "github.com/stratumn/go-chainscript",
      "refs": [
        { "linkHash": "Kg==", "process": "p1" },
        { "linkHash": "GA==", "process": "p2" }
      ],
      "process": { "name": "test_process" },
      "mapId": "test_map"
    }
  },
  "meta": { "linkHash": "hNMct7lGhGO0dKwkcwjZeV++qkdUcznH671kGPiM28I=" }
}
```

**segment-evidences**: a segment with two external evidences.

```json
{
  "link": {
    "version": "1.0.0",
    "meta": {
      "clientId": "github.com/stratumn/go-chainscript",
      "process": { "name": "test_process" },
      "mapId": "test_map"
    }
  },
  "meta": {
    "linkHash": "Sp1W1ORf18KU0ztX641sN7bp50SnWH/C3a4zaF6Edds=",
    "evidences": [
      {
        "version": "0.1.0",
        "backend": "bitcoin",
        "provider": "testnet",
        "proof": "Kg=="
      },
      {
        "version": "1.0.3",
        "backend": "ethereum",
        "provider": "mainnet",
        "proof": "GA=="
      }
    ]
  }
}
```

**segment-signatures**: a segment with two link signatures (RSA and Ed25519).

```json
{
  "link": {
    "version": "1.0.0",
    "meta": {
      "clientId": "github.com/stratumn/go-chainscript",
      "process": { "name": "test_process" },
      "mapId": "test_map"
    },
    "signatures": [
      {
        "version": "1.0.0",
        "payloadPath": "[version,data,meta]",
        "publicKey": "LS0tLS1CRUdJTiBFRDI1NTE5IFBVQkxJQyBLRVktLS0tLQpNQ293QlFZREsyVndBeUVBRk5hL3BYL0ErTFVac3ZxbzFyRDBSSnVTVnZ5TWNsZ1lzZDNZUndrSWc0UT0KLS0tLS1FTkQgRUQyNTUxOSBQVUJMSUMgS0VZLS0tLS0K",
        "signature": "LS0tLS1CRUdJTiBNRVNTQUdFLS0tLS0KSDRtSnNmZkRRVS9ubDAxcjFDQi9tQzhvVW4rZVkvQzd1VVl4OEVEU1BYOCtHRWFES2JsSFMxbjlPOTV0N3VQaQpoSE1ES2pOUkFVekpQNGZzOGszQ0FnPT0KLS0tLS1FTkQgTUVTU0FHRS0tLS0tCg=="
      },
      {
        "version": "1.0.0",
        "payloadPath": "[version,meta.mapId]",
        "publicKey": "LS0tLS1CRUdJTiBSU0EgUFVCTElDIEtFWS0tLS0tCk1JSUJJakFOQmdrcWhraUc5dzBCQVFFRkFBT0NBUThBTUlJQkNnS0NBUUVBem0yWGdyUzRuSGdTM29BczhKRE4KU0grL0hGOW9Ha1R4ZWszRWpiVHZ6d3F2TzNuc0N5SEZDSm5XWjIwcnFUemtXTGpuaFdDRStGUzUzd1NsS24wWAp2TzFtbmVzaDFwNHVXYmlWQktIUzhmOGl6Z2dBY2ZEbHY1TjNEeUx5dFBxQ0pwMmZ2WnUwL09POHVOdmdwOE5VCjBLOTR4dklUR1N4MnNGZGN6clJCZFNLVzJxY0RmbmdSMjFmVVV6dE02ZzRjWVB6MjBSekhwRWNHd282MnJRcVYKNjlBemxhYnNpWlZrTzNnNlZVL0FPbXFMMXFXWDU2aGZ1ZzFkNnpaaXBuSjNUclZZQlBBY3VTSjVmUHdtb1lrNgpycFVpckgvRWdTZVBUQkpEZXVBWWRaM2xPUS9ZTDRicTR4SjdOMVVSU0o0Nk1meElKYlhyZzVBc2VNY1Z3ekIrClF3SURBUUFCCi0tLS0tRU5EIFJTQSBQVUJMSUMgS0VZLS0tLS0K",
        "signature": "LS0tLS1CRUdJTiBNRVNTQUdFLS0tLS0KeDBTM2JDbHh2TlQyV0doanEreDUvUDFrRGlUaGhWVjZJMkczYVVMMDlQbFlNODhSeGprM1V2ZENLY3ZmNFpweApJTmdNeHlRdmNUcWdZOUxOdmRIR3NpdmtDNHdlbVpsb1l2bUNsTmZhWThGSUxmcXNyTG1yOC9TL2g1THhZTUcwCkhWbnpUZFppRlU0RDVzYnRqa3VIUVpOMDRLMkpLeDQ3a2hhWkp6V3BqVnlwMytFQmljZXhLYW9MT2k0bGdGMlgKbS9PR3ZkNVRoR1FwcW4yTmQ3cFU1dTErSUpScGR4VWVnRVd6VEJqa2NUY2xERktDZXpTTVFFMmFQVk9jV1lCago0eUFaU0srQXZmOHJyM1Z0WUxvVDFGekJ4ZFNIOWpWTkw1MXZ0eFNwSXpLaTBGNldFVEJPMlB3TFhwT0k1VzFyCmJ0aEJmcHdETlJzVUpFRGdTNEcyU1E9PQotLS0tLUVORCBNRVNTQUdFLS0tLS0K"
      }
    ]
  },
  "meta": { "linkHash": "d5ufjKXj2mpbtiMIEZzNGv5JEnJneuIUYvR25ai0sio=" }
}
```

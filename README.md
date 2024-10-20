
![CI status](https://github.com/aalbacetef/tofu/actions/workflows/ci.yml/badge.svg)   [![License](https://img.shields.io/badge/License-BSD_3--Clause-blue.svg)](https://opensource.org/licenses/BSD-3-Clause) 


# tofu

## Introduction

This package implements the TOFU (trust on first use) authentication scheme.

To read more about it, check out the following links:
 - [TOFU - wiki](https://en.wikipedia.org/wiki/Trust_on_first_use)
 - [Gemini spec - TLS](https://geminiprotocol.net/docs/specification.gmi#4-tls)


## Usage 

`tofu` provides an interface, `Store`, to allow the library consumer to 
choose how they want to manage their known hosts. 

There are two implementations, a `FileStore` and an `InMemoryStore`.  
When using `FileStore`, the implementation assumes a format similar to the 
`known_hosts` file used by SSH, each line is a comma-separated set of values:

- hash(address)
- fingerprint - hash(data)
- comment (optional)

### Example 

Define a connection verification function.

```go
var (
    ErrInvalidCert = errors.New("invalid certificate")
    ErrNoPeerCerts = errors.New("no peer certificates")
)

func verifyConn(store tofu.Store) verifyFunc {
    return func(state tls.ConnectionState) error {
        // check for peer certificates
        peerCerts := state.PeerCertificates
        if len(peerCerts) == 0 {
            return ErrNoPeerCerts
        }

        leaf := state.PeerCertificates[0]

        // get fingerprint
        host := tofu.Host{
            Address:     state.ServerName,
            Fingerprint: tofu.Fingerprint(leaf),
        }

        // verify host with fingerprint
        valid, err := tofu.Verify(store, host)
        if err != nil {
            return fmt.Errorf("error verifying: %w", err)
        }

        if !valid {
            return ErrInvalidCert
        }

        return nil
    }
}
```

Pass it into a `tls.Config` 

```go
const minTLSVersion = tls.VersionTLS12

config := &tls.Config{
	MinVersion:         minTLSVersion,
	InsecureSkipVerify: true,
	VerifyConnection:   verifyConn(myCertStore),
}
```

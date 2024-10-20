
![CI status](https://github.com/aalbacetef/tofu/actions/workflows/ci.yml/badge.svg)   [![License](https://img.shields.io/badge/License-BSD_3--Clause-blue.svg)](https://opensource.org/licenses/BSD-3-Clause) 


# Introduction

This package implements the TOFU (trust on first use) authentication scheme.

To read more about it, check out the following links:
 - [TOFU - wiki](https://en.wikipedia.org/wiki/Trust_on_first_use)
 - [Gemini spec - TLS](https://geminiprotocol.net/docs/specification.gmi#4-tls)


# Usage 

The package provides an interface, `Store`, to allow the library consumer to 
choose however they want to handle known hosts. 

There are two implementations, a `FileStore` and a `InMemoryStore`.  
When using `FileStore`, the implementation assumes a format similar to the 
`known_hosts` file used by SSH, that is, each line is a comma-separated set of values:

- hash(address)
- fingerprint - hash(data)
- comment (optional)



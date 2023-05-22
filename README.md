ardconv
=======

[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)
[![Latest Release](https://img.shields.io/github/release/tliron/ardconv.svg)](https://github.com/tliron/ardconv/releases/latest)
[![Go Report Card](https://goreportcard.com/badge/github.com/tliron/ardconv)](https://goreportcard.com/report/github.com/tliron/ardconv)

Utility to convert between these [ARD (Agnostic Raw Data)](https://github.com/tliron/ard) formats:

* YAML
* Lossy JSON
* JSON with an ARD-compatible schema
* XML with an ARD-compatible schema
* CBOR
* MessagePack

Can also be used to validate ARD and prettify/colorize some formats.

Can work on files or via piping stdout to stdin.

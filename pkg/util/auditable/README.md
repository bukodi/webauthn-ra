This package provides a simple interface for a key-value map with history and auditing functions. 
It is intended to be used to store configuration of system in a auditable way.

## Data model
The MapEntry should contain the key and value of the entry.
The whole entry should support marshal and unmarshal to and from a byte slice.
The key should be retrieved as a byte slice.

## Features
- [x] Set data structure with operations: add, delete, update, get, list
- [x] Master hash, what represents the contents of all active entries in the set
- [x] Pluggable storage backend
- [x] Transaction support
- [x] Ability to authorize and/or sign transactions
- [x] History for individual entries
- [x] History for the whole set
- [x] Export proofs of a specific entry or for whole set
- [x] Checkpoint support
- [x] History pruning
- [x] Anonymization historical entries 
- [x] Manage archives with Bloom-filter

This package provides a simple interface for a key-value map with history and auditing functions. 
It is intended to be used to store configuration of system in a auditable way.

## Data model
The MapEntry should contain the key and value of the entry.
The whole entry should support marshal and unmarshal to and from a byte slice.
The key should be retrieved as a byte slice.

# guuid

This package is a wrapper for most common used UUID package:
https://github.com/google/uuid

It supports the following UUID versions:
1. Version 1, based on timestamp and MAC address (RFC-4122)
1. Version 2, based on timestamp, MAC address and POSIX UID/GID (DCE 1.1)
1. Version 3, based on MD5 hashing of a named value (RFC-4122)
1. Version 4, based on random numbers (RFC-4122)
1. Version 5, based on SHA-1 hashing of a named value (RFC-4122)
# SHA-miner

SHA-miner is a CLI tool for finding the smallest possible SHA256 hash value for a given string by appending to it a randomly generated nonce.

## Usage

Compile with

```go
go build
```

Run with

```go
miner "some content" --threads 16
```

The program starts to try different nonces. As soon as it finds a nonce with a hash value smaller than the previous smallest value, it will be printed to stdout in the following format: `<timestamp> <nonce> <hash>`.

After every 10000000 tries a notice will appear to stderr.

Thus it is a decent UNIX program: the actual answers go into stdout, while other messages land in stderr.


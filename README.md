# SHA-miner

SHA-miner is a CLI tool for finding the smallest possible SHA256 hash value for a given string by appending to it a randomly generated nonce.

Although it does a decent job, it is not designed for industry usage, but is more like a showcase of how parallelism and sharing data between goroutines works in Go.

## Usage

Compile with

```go
go build
```

Run with

```go
miner "some content" --threads 16
```

The program starts to try different nonces. As soon as it finds a nonce with a hash value smaller than the previous smallest value, it will be printed to stdout in the following format: `<timestamp> <nonce> <hash>`. After every 10000000 tries a notice will appear to stderr.
This it is a decent UNIX program: the actual answers go into stdout, while other messages land in stderr.


# [argon](https://github.com/AlessandroBellati/argon)
Go library providing a simple and efficient way to generate **Argon2id hashes**, which are widely recognized for their security and performance.
The library offers flexibility with configurable parameters for:
- **salt length**
- **number of iterations**
- **memory usage**
- **number of threads**
- **key length**.

## Installation
```bash
go get github.com/AlessandroBellati/argon
```

## Usage
```go
package main

import (
    "fmt"
    "github.com/AlessandroBellati/argon"
)

func init(){
    // set custom parameters
    argon.SaltLen = 64
    argon.Time = 1
    argon.Memory = 64 * 1024
    argon.Threads = 4
    argon.KeyLen = 64
}

func main() {
    password := "password"
    hash, salt, err := argon.Argon2idSalt(password)
    
    if err != nil {
        // handle error
    }

    // store hash and salt in database or somewhere else
}
```

## [Security for time-bounded defenders](https://datatracker.ietf.org/doc/html/draft-irtf-cfrg-argon2-03#section-9.3)
A bottleneck in a system employing the password-hashing function is
often the function latency rather than memory costs.
A rational defender would then maximize the bruteforce costs for the attacker equipped with a list of hashes, salts, and timing information, for
fixed computing time on the defender's machine.
The attack cost estimates from [AB16](https://datatracker.ietf.org/doc/html/draft-irtf-cfrg-argon2-03#ref-AB16) imply that for Argon2i, 3 passes is almost optimal for the most of reasonable memory sizes, and that for Argon2d and Argon2id, 1 pass maximizes the attack costs for the constant
defender time.

### Recommendations
The Argon2id variant with t=1 and maximum available memory is
recommended as a default setting for all environments.
This setting is secure against side-channel attacks and maximizes adversarial costs on dedicated bruteforce hardware.

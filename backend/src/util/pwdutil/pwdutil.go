package pwdutil

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"src/common/ctype"
	"src/util/errutil"
	"src/util/i18nmsg"
	"strings"

	"golang.org/x/crypto/argon2"
)

type params struct {
	memory      uint32
	iterations  uint32
	parallelism uint8
	saltLength  uint32
	keyLength   uint32
}

func MakePwd(pwd string) string {
	if pwd == "" {
		return ""
	}

	// Recommended parameters for Argon2id
	p := &params{
		memory:      64 * 1024, // 64MB
		iterations:  3,         // 3 iterations
		parallelism: 2,         // 2 parallel threads
		saltLength:  16,        // 16-byte salt
		keyLength:   32,        // 32-byte hash
	}

	// Generate a random salt
	salt := make([]byte, p.saltLength)
	if _, err := rand.Read(salt); err != nil {
		return ""
	}

	// Generate the hash
	hash := argon2.IDKey(
		[]byte(pwd),
		salt,
		p.iterations,
		p.memory,
		p.parallelism,
		p.keyLength,
	)

	// Base64 encode the salt and hash
	b64Salt := base64.RawStdEncoding.EncodeToString(salt)
	b64Hash := base64.RawStdEncoding.EncodeToString(hash)

	// Format: $argon2id$v=19$m=memory,t=iterations,p=parallelism$salt$hash
	encodedHash := fmt.Sprintf(
		"$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s",
		argon2.Version, p.memory, p.iterations, p.parallelism, b64Salt, b64Hash,
	)

	return encodedHash
}

func CheckPwd(pwd, encodedHash string) error {
	// Extract the parameters, salt, and derived key from the encoded hash
	parts := strings.Split(encodedHash, "$")
	if len(parts) != 6 {
		return errutil.New(i18nmsg.InvalidHashFormat)
	}

	// Check the algorithm
	if parts[1] != "argon2id" {
		return errutil.NewWithArgs(
			i18nmsg.UnsupportedAlgorithm,
			ctype.Dict{
				"Value": parts[1],
			},
		)
	}

	// Parse the parameters
	var version int
	p := &params{}
	_, err := fmt.Sscanf(parts[2], "v=%d", &version)
	if err != nil {
		return err
	}

	_, err = fmt.Sscanf(parts[3], "m=%d,t=%d,p=%d", &p.memory, &p.iterations, &p.parallelism)
	if err != nil {
		return err
	}

	// Decode the salt and hash
	salt, err := base64.RawStdEncoding.DecodeString(parts[4])
	if err != nil {
		return err
	}
	p.saltLength = uint32(len(salt))

	hash, err := base64.RawStdEncoding.DecodeString(parts[5])
	if err != nil {
		return err
	}
	p.keyLength = uint32(len(hash))

	// Compute the hash of the provided pwd
	computedHash := argon2.IDKey(
		[]byte(pwd),
		salt,
		p.iterations,
		p.memory,
		p.parallelism,
		p.keyLength,
	)

	// Compare the computed hash with the stored hash
	if base64.RawStdEncoding.EncodeToString(computedHash) == parts[5] {
		return nil
	}

	return errutil.New(i18nmsg.InvalidUsernameOrPwd)
}

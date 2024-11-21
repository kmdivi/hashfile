package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	var (
		hash_algorithm = flag.String("h", "sha256", "string flag")
		filename       = flag.String("f", "default", "string flag")
	)
	flag.Parse()

	hash_algorithmStr := strings.ToLower(*hash_algorithm)
	hash_algorithm = &hash_algorithmStr

	hash_value := ""

	fmt.Println(*hash_algorithm)

	switch *hash_algorithm {
	case "sha256":
		hash_value = get_sha256(*filename)
	case "sha1":
		hash_value = get_sha1(*filename)
	case "md5":
		hash_value = get_md5(*filename)
	default:
		hash_value = get_sha256(*filename)
	}

	fmt.Printf("%s", hash_value)
}

func get_sha256(filename string) string {
	r, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	hash := sha256.New()
	if _, err := io.Copy(hash, r); err != nil {
	}
	hash_value := fmt.Sprintf("%x", hash.Sum(nil))

	return hash_value
}

func get_sha1(filename string) string {
	r, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	hash := sha1.New()
	if _, err := io.Copy(hash, r); err != nil {
	}
	hash_value := fmt.Sprintf("%x", hash.Sum(nil))

	return hash_value
}

func get_md5(filename string) string {
	r, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	hash := md5.New()
	if _, err := io.Copy(hash, r); err != nil {
	}
	hash_value := fmt.Sprintf("%x", hash.Sum(nil))

	return hash_value
}

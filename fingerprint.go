package tofu

import (
	"crypto/md5"
	"crypto/x509"
	"fmt"
	"strings"
)

func hashFunc(data []byte) string {
	hash := md5.Sum(data)

	return string(hash[:])
}

// Fingerprint returns the md5 hash of the DER encoded bytes.
func Fingerprint(cert *x509.Certificate) string {
	hash := hashFunc(cert.Raw)
	n := len(hash)
	bdr := &strings.Builder{}

	for _, h := range hash[:n-1] {
		fmt.Fprintf(bdr, "%02X:", h)
	}

	fmt.Fprintf(bdr, "%02X", hash[n-1])

	return bdr.String()
}

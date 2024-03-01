// This module sign an OCI image with Cosign (https://github.com/sigstore/cosign)
package main

import (
	"context"
	"fmt"
)

type SignOci struct{}

// Sign OCI image with Cosign 🖊
func (t *SignOci) SignImage(
	ctx context.Context,
	// +default="ttl.sh/dagger-build-and-push-oci-image-example:latest"
	imageAddr string,
	cosignPrivateKey *Secret,
	cosignPassword *Secret,
) (string, error) {
	cosignKey, err := cosignPrivateKey.Plaintext(ctx)
	if err != nil {
		return "", err
	}
	cosignPwd, err := cosignPassword.Plaintext(ctx)
	if err != nil {
		return "", err
	}
	cosignCmd := fmt.Sprintf("cosign sign --yes --key env://COSIGN_PRIVATE_KEY %s", imageAddr)
	return dag.Container().
		From("bitnami/cosign:2.2.3@sha256:3fd113db0d9c9c7b5d7db53ceb5ef32b09f0c46947dcb752e93578faec15d2a9").
		WithEnvVariable("COSIGN_PRIVATE_KEY", cosignKey).
		WithEnvVariable("COSIGN_PASSWORD", cosignPwd).
		// WithEnvVariable("REGISTRY_PASSWORD", os.Getenv("REGISTRY_PASSWORD")).
		WithEntrypoint([]string{"sh", "-c"}).
		WithExec([]string{cosignCmd}).
		Stderr(ctx)
}

// This module scan container images for vulnerabilities by leveraging trivy (https://github.com/aquasecurity/trivy)
package main

import (
	"context"
	"strconv"
)

type Trivy struct{}

// Scan the specified OCI image with Trivy 🔍
func (t *Trivy) ScanImage(
	ctx context.Context,
	imageRef string,
	// +optional
	// +default="HIGH,CRITICAL"
	severity string,
	// +optional
	// +default="os,library"
	vulnType string,
	// +optional
	// +default=1
	exitCode int,
	// +optional
	// +default="table"
	format string,
) (string, error) {
	return dag.
		Container().
		From("aquasec/trivy:0.49.1@sha256:91494b87ddc64f62860d52997532643956c24eeee0d0dda317d563c28c8581bc").
		WithExec([]string{
			"image",
			"--quiet",
			"--severity", severity,
			"--vuln-type", vulnType,
			"--exit-code", strconv.Itoa(exitCode),
			"--format", format,
			imageRef,
		}).
		Stdout(ctx)
}

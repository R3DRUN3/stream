// This module builds and push an OCI image (starting from a Dockerfile)
package main

import (
	"context"
)

type BuildAndPushOciImage struct{}

// Builds and push an OCI image
func (t *BuildAndPushOciImage) BuildAndPushOciImage(
	ctx context.Context,
	// +default="ttl.sh/dagger-build-and-push-oci-image-example:latest"
	tag string,
	dockerfilePath string,
) error {
	imageContainer := dag.CurrentModule().Source().Directory(dockerfilePath).DockerBuild()
	_, err := imageContainer.Publish(ctx, tag)
	if err != nil {
		panic(err)
	}
	return nil
}

# BUILD
Custom Dagger modules for *builds* (binaries, OCI images etc.)


## Example Usage  

### Build and Push OCI image
See function help:  
```console
cd build-and-push-oci-image && dagger call build-and-push-oci-image --help
```  

Output:  
```console
Builds and push an OCI image

Usage:
  dagger call build-and-push-oci-image [flags]

Flags:
      --dockerfile-path string   
      --tag string                (default "ttl.sh/dagger-build-and-push-oci-image-example:latest")  
```  
Call the function to test it against a local dockerfile:  
```console
dagger call build-and-push-oci-image --dockerfile-path test-module
```  

Now try to pull the image from the registry and run it:  
```console
docker pull ttl.sh/dagger-build-and-push-oci-image-example:latest \
&& docker run -it --rm ttl.sh/dagger-build-and-push-oci-image-example:latest
```  

Output:  
```console
latest: Pulling from dagger-build-and-push-oci-image-example
4abcf2066143: Pull complete 
bbd588c47b24: Pull complete 
7892c7c1459f: Pull complete 
e0fbf7dc7f57: Pull complete 
Digest: sha256:bc6a1093c186f6089c4e4b8727c4fb1372cdb27f964e5470c053108b5d3063af
Status: Downloaded newer image for ttl.sh/dagger-build-and-push-oci-image-example:latest
ttl.sh/dagger-build-and-push-oci-image-example:latest

What's Next?
  View a summary of image vulnerabilities and recommendations → docker scout quickview ttl.sh/dagger-build-and-push-oci-image-example:latest
Hello from this Dagger Custom module! 👋
```  







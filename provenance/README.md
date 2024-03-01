# Provenance
Custom Dagger modules for artifact *provenance* (signature, attestations etc.)


## Example Usage  

### Sign OCI image
See function help:  
```console
cd oci-signature && dagger call sign-image --help
```  

Output:  
```console
Sign OCI image with Cosign 🖊

Usage:
  dagger call sign-image [flags]

Flags:
      --cosign-password Secret       (default :)
      --cosign-private-key Secret    (default :)
      --image-addr string            (default "ttl.sh/dagger-build-and-push-oci-image-example:latest")
```  

In order to test this generate a pair of cosign keys inside the `test-module` directory:  
```console
cd test-module && cosign generate-key-pair
```  
Now run the module with the following command (after exporting the env vars):  
```console
dagger call sign-image --cosign-private-key=env:COSIGN_PRIVATE_KEY --cosign-password=env:COSIGN_PASSWORD
```  




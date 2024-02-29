# SAST
Custom Dagger modules for *Static Application Security Testing*.


## Example Usage  

### Scan OCI image for vulnerabilities with Trivy
See function help:  
```console
cd trivy && dagger call scan-image --help
```  

Output:  
```console
Pull the official Trivy image and send the trivy CLI an image ref to scan

Usage:
  dagger call scan-image [flags]

Flags:
      --exit-code int
      --format string       (default "table")
      --image-ref string
      --severity string     (default "UNKNOWN,LOW,MEDIUM,HIGH,CRITICAL")
      --vuln-type string  
```  
Call the function against python:3.4 (this should fail):  
```console
dagger call scan-image --vuln-type os --image-ref python:3.4
```  

Call the function against alpine:latest (this should not fail):  
```console
dagger call scan-image --vuln-type os --image-ref alpine:latest
```  






# googleiapclient

A golang library and reference command line tool which provides a way to use a
service account to programmatically access resources behind Google Cloud's
Identity Aware Proxy.

### Note:

This is a simplified fork of (https://github.com/ryanchapman/googleiapclient). I
put this fork together to simplify building and running of the utility.  For
full documention, you should reference
(https://godoc.org/github.com/ryanchapman/googleiapclient) 

## Docker Compose building and usage.

### Building

```
docker-compose -f deployments/docker-compose.yml build iapclient
```

### Running


#### Usage:

```
    docker run -it iapclient:latest --help
```

#### Full example
##### Service Account:

To operate the iapclient, you need to have a google cloud service account with
the role of 'IAP-secured Web App User'.  The key from the service account is
then used to auth.

##### IAP Client ID:

You can get the OAuth Client ID by looking at the returned location header. Look
for the client_id query parameter.

Or you can use awk.

```
CLIENT_ID=$(curl -I https://iap.protetected-domain.com | awk '/^location/ {split($NF, a, /[=&]/); print a[2]}')
```

```
    GC_CREDS=$(cat svcaccnt.json | base64)
    AUTH_HEADER=$(docker run -it iapclient:latest \
      --oauth-client-id=$CLIENT_ID \
      --google-credentials $CREDS \
      --requested-expiration 1h)

$ curl -D- -s -H "$AUTH_HEADER" https://iap.protetected-domain.com
HTTP/2 200
[...]
```

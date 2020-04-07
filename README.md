# googleiapclient

A golang library and reference command line tool which provides a way to use a service account to programmatically access
resources behind Google Cloud's Identity Aware Proxy.

### Note:

This is a simplified fork of https://github.com/ryanchapman/googleiapclient. I put this fork together to simplify building
and running of the utility.  For full documention, you should reference (https://godoc.org/github.com/ryanchapman/googleiapclient) 

## Docker Compose building and usage.

### Building

```
docker-compose -f deployments/docker-compose.yml build iapclient
```

### Running

Usage:

```
    docker run -it iapclient:latest --help
```

Full example:

```
    GC_CREDS=$(cat svcaccnt.json | base64)
    AUTH_HEADER=$(docker run -it iapclient:latest \ 
      --oauth-client-id=823926513327-pr0714rqtdb223bahl0nq2jcd4ur79ec.apps.googleusercontent.com 
      --google-credentials $CREDS \
      --requested-expiration 1h)

$ curl -D- -s -H "$AUTH_HEADER" https://test.initech.com
HTTP/2 200
[...]
```

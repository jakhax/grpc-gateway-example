# Secure the communication with SSL

```bash
mkdir cert
# first we generate an rsa key, server.key
openssl genrsa -out cert/server.key 2048
# next a self signed certificate, server.crt
openssl req -new -x509 -sha256 -key cert/server.key -out cert/server.crt -days 3650

openssl req -new -sha256 -key cert/server.key -out cert/server.csr

openssl x509 -req -sha256 -in cert/server.csr -signkey cert/server.key -out cert/server.crt -days 3650
```
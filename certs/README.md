openssl genrsa -out server.key 2048
openssl genrsa -aes256 -out ca.key 4096
openssl req -new -x509 -sha256 -days 730 -key ca.key -out ca.crt

openssl req -new -x509 -days 3650 -subj "/C=CN/L=Hangzhou/O=bzd111/CN=localhost" -sha256 -key server.key -out server.csr
openssl x509 -req -days 365 -sha256 -in server.csr -CA ca.crt -CAkey ca.key -set_serial 1 -out server.crt

openssl genrsa -out client.key 2048

<!-- openssl req -new -key client.key -out client.csr -->

openssl req -new -x509 -days 3650 -subj "/C=CN/L=Hangzhou/O=bzd111/CN=localhost" -key client.key -out client.csr
openssl x509 -req -days 365 -sha256 -in client.csr -CA ca.crt -CAkey ca.key -set_serial 2 -out client.crt

<!-- openssl req -new -x509 -days 3650 \ -->
<!--     -subj "/C=GB/L=China/O=grpc-client/CN=client.grpc.io" \ -->
<!--     -key client.key -out client.crt -->

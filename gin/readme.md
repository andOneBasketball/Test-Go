## 生成私钥
openssl genrsa -out ./assets/server.key 2048
## 生成证书签名请求
openssl req -new -key ./assets/server.key -out ./assets/server.csr
## 使用私钥签署证书 (server.pem)
openssl x509 -req -days 365 -in ./assets/server.csr -signkey ./assets/server.key -out ./assets/server.pem
## 检查生成的证书和私钥
openssl x509 -in ./assets/server.pem -text -noout
openssl rsa -in ./assets/server.key -check

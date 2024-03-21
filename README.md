# Redirects experiment

### Update hosts

```
127.0.0.1 cel.sergii.org
````

### Start redirects server

```sh
cd cel-go-redirects/serve-redirects
./install.sh
# Needs root to serve on port 443
./serve.sh
```

The server is up now, in verify in another shell

```sh
cel-go-redirects/serve-redirects/verify.sh
```

### Attemp go build


```sh
cd cel-go-redirects/go-use-redirects
./build.sh
```

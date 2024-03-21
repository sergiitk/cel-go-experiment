import http.server
import ssl

server_address = ("cel.wtf", 443)
DIRECTORY = "web"

class Handler(http.server.SimpleHTTPRequestHandler):
    def __init__(self, *args, **kwargs):
        super().__init__(*args, directory=DIRECTORY, **kwargs)

httpd = http.server.HTTPServer(server_address, Handler)
sslctx = ssl.SSLContext(ssl.PROTOCOL_TLS_SERVER)
sslctx.check_hostname = False
sslctx.load_cert_chain(certfile='certs/server.pem', keyfile="certs/key.pem")
httpd.socket = sslctx.wrap_socket(httpd.socket, server_side=True)
print(f"Serving on: {httpd.socket}")
httpd.serve_forever()

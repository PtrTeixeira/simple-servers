import http.server
import os
import socketserver

PORT = int(os.getenv("PORT", "3000"))

class RequestHandler(http.server.BaseHTTPRequestHandler):
    def do_GET(self):
        if self.path != "/":
            self.send_response(404)
            return
        self.send_response(200)
        self.send_header('Content-type', 'text/plain')
        self.end_headers()
        self.wfile.write(bytes("Hello World!\n", 'utf-8'))
        return

Handler = RequestHandler

with socketserver.TCPServer(("", PORT), Handler) as httpd:
    print("Listening on port", PORT)
    httpd.serve_forever()

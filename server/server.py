import http.server
import socketserver
import socket
from http import HTTPStatus
import argparse

# Store the current IP address
current_ip = ""

# Define a custom request handler to log the IP address
class MyHandler(http.server.SimpleHTTPRequestHandler):
    def do_GET(self):
        global current_ip
        if self.path == '/get-ip':
            self.send_response(HTTPStatus.OK)
            self.send_header('Content-type', 'text/plain')
            self.end_headers()
            self.wfile.write(current_ip.encode('utf-8'))
        else:
            super().do_GET()

def get_current_ip():
    # Get the current IP address
    global current_ip
    current_ip = socket.gethostbyname(socket.gethostname())

def main():
    parser = argparse.ArgumentParser(description='Local server with dynamic IP display')
    parser.add_argument('--port', type=int, default=8000, help='Port number for the local server')
    args = parser.parse_args()

    get_current_ip()

    # Start the HTTP server on the specified port
    with socketserver.TCPServer(("", args.port), MyHandler) as httpd:
        print(f"Local server started at http://localhost:{args.port}")
        httpd.serve_forever()

if __name__ == "__main__":
    main()

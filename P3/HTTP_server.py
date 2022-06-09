import http.server
import json

port = 8000

class MyHttp(http.server.BaseHTTPRequestHandler):
    status = "OK" #changes with posting
    def do_GET(self):
        if self.path == "/api/v1/status":
            if MyHttp.status == "OK":
                self.send_response(200)
            else:
                self.send_response(201)
            
            self.send_header('Content-type','text/json')
            self.end_headers()
            massage = json.dumps({"status":MyHttp.status})

            # Send the html message
            self.wfile.write(bytes(massage, 'utf-8'))
        return
    
    def do_POST(self):
        if self.path == "/api/v1/status":
            print("post path ok")
        return
    

if __name__ == '__main__':
    from http.server import HTTPServer
    server = HTTPServer(('localhost', port), MyHttp)

    print('Starting server @port '+str(port)+' , use <Ctrl-C> to stop')
    try:
        server.serve_forever()
    except KeyboardInterrupt:
        pass
    server.server_close()
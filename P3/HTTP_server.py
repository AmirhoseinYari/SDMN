import http.server
import json

port = 8000

class MyHttp(http.server.BaseHTTPRequestHandler):
    status = "OK" #changes with posting

    # curl localhost:8000/api/v1/status
    def do_GET(self):
        if self.path == "/api/v1/status":
            if MyHttp.status == "OK":
                self.send_response(200)
            else:
                self.send_response(201)
            
            self.send_header('Content-Type','application/json')
            self.end_headers()
            massage = json.dumps({"status":MyHttp.status})

            # Send the html message
            self.wfile.write(bytes(massage, 'utf-8'))
        return
    
    # curl -X POST localhost:8000/api/v1/status -H "Content-Type: application/json" -d '{"status": "not OK"}'
    def do_POST(self):
        if self.path == "/api/v1/status":
            #print("post path ok")
            try:
                datalen = int(self.headers['Content-Length'])
                data = json.loads(self.rfile.read(datalen).decode("utf-8"))
            except:
                data = {"status":MyHttp.status} # last status
            
            #print(data,data['status'],type(data))
            if data['status'] != MyHttp.status:
                self.send_response(201)
                self.send_header('Content-Type','application/json')
                self.end_headers()
                MyHttp.status = data['status']
                massage = json.dumps({"status":MyHttp.status})
                self.wfile.write(bytes(massage, 'utf-8'))
            else:
                self.send_response(200)
                print('noting changed!')
        return
    

if __name__ == '__main__':
    from http.server import HTTPServer
    server = HTTPServer(('0.0.0.0', port), MyHttp)

    print('Starting server @port '+str(port)+' , use <Ctrl-C> to stop')
    try:
        server.serve_forever()
    except KeyboardInterrupt:
        pass
    server.server_close()
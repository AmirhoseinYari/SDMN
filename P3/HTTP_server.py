import http.server

port = 8000

class MyHttp(http.server.BaseHTTPRequestHandler):

    def do_Get(self):

        return
    

if __name__ == '__main__':
    from http.server import HTTPServer
    server = HTTPServer(('localhost', port), MyHttp)
    print('Starting server @port '+str(port)+' , use <Ctrl-C> to stop')
    server.serve_forever()
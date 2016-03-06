lets analyze this program, starting with the main method.

http.HandleFunc("/", hello)
func HandleFunc(pattern string, handler func(ResponseWriter, *Request))
- registers the handler for the given pattern. In this case, it registers the handler 'hello' for the "/" pattern
- this doesn't have to implement ServeHTTP

http.Handle("/hello", http.RedirectHandler("/", http.StatusFound))
func Handle(pattern string, handler Handler)
- this is the simpler version of HandleFunc, where you serve the requests matching the pattern with a Handler. 
- "Handler" is a type that requires you to implement the ServeHTTP method
- in this case, we use the "/hello" string to match, and use a handler that redirects a user to the home page, setting a 302 "Found" status on the response.

log.Fatalln(http.ListenAndServe(":8080", &WrapHTTPHandler{http.DefaultServeMux}))
- Fatalln is equivalent to Println() followed by a call to os.Exit(1).
func ListenAndServe(addr string, handler Handler) error
- listens to requests on the port and calls ListenAndServe() on a server object created from the address and handler provided. 
- usually you just leave Handler as nil, which means that DefaultServeMux is used as the handler. In this case however, we are providing our own interface instead, called WrapHTTPHandler
- WrapHTTPHandler has a field in it called "m" that is actually an http.Handler, and we are setting this handler to be http.DefaultServeMux
- we are also overriding the ServeHTTP method of DefaultServeMux (which is of type ServeMux). This overridden method is providing us with the ability to wrap the Handlers that we use in this web server, which means we can add things to each Handler call like logging. 

type WrapHTTPHandler struct

func (h *WrapHTTPHandler) ServeHTTP(w http.ResponseWriter, r *http.Request)

type loggedResponse struct

func (l *loggedResponse) WriteHeader(status int)

func hello(w http.ResponseWriter, r *http.Request)

func goodbye(w http.ResponseWriter, r *http.Request)
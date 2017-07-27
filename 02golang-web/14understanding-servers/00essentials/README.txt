To understand web dev with Go,
you must understand these core concepts:

INTERFACE

REQUEST PESPONSE pattern

http.Handler
-----------------------------------
type Handler interface{
	ServeHTTP(ResponseWritter, *Request)
}
-----------------------------------
A HANDLER will handle a request by providing a response.

web programming synonymous terms:
router
request router
multiplexer
mux
servemux
server
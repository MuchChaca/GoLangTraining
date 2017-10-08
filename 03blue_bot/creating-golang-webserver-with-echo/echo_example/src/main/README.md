Middlewares
// From video
We could say that http request go to a pipeline. At the end we have the handler method that process respond with data to the client.
Instead of doing an authentification each an every handler method, we have middlewares.
Inside the pipeline we can use methods (middlewares) that take the http request, manipulate it and 
either pass it to the next middleware or reject it and give it back to the client. It's done until the http
request will be rejected or will get to the handler method

// From official website of echo
Middleware is a function chained in the HTTP request-response cycle with access to Echo#Context which it uses to perform a specific action, 
for example, logging every request or limiting the number of requests.

Handler is processed in the end after all middleware are finished executing.
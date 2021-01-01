# hello-gRPC-go

BigtThanks to [Tech School's FREE course on dev.to](https://dev.to/techschoolguru/series/7311) which helped me greatly to learn gRPC and answer the following questions.

### 1. Why use gRPC instead of rest?

The first most important answer to this question is the usage of HTTP/2 in gRPC. REST can use HTTP/2 or the much more common HTTP/1.1. HTTP/2 has some great features over its predecessor. Which include, binary protocol instead of text-based protocol, compressed headers, multiplexing on same TCP connection, and server-push. So I think it's a fairer comparison to not use this attribute when comparing the two.

#### So HTTP/2 aside why else would I use gRPC?

Polyglot auto-code generation with [protocol buffers](https://developers.google.com/protocol-buffers/docs/overview). You can create request and payload functionality in one language [protocol buffers](https://developers.google.com/protocol-buffers/docs/overview) and automatically generate code for many popular languages: Java, Ruby, etc. Not only does this reduce code needed written by a developer, it creates a universal contract that all services can share and use. This is great because as platforms grow contracts can be strictly maintained between services.

So in summary you can deliver more performant code faster that is more maintainable. Awesome!

#### But what is the downside?

* It looks like http/2 is supported in most recent versions of web browsers, but older versions don't support.
* I think also there is less talent and support for gRPC, but given that it is easy to learn, I would say onramp time is probably not so bad, especially considering automatic code generation.
* Probably better suited for backend applications although I do see web and mobile client solutions that uses it.

### 2. how is middleware handled

You use functions called interceptors, they serve the niche that middleware functions have. Interceptors in gRPC are used both on server and client.

[Example from techschool of server interceptor](https://dev.to/techschoolguru/use-grpc-interceptor-for-authorization-with-jwt-1c5h)

````go

func unaryInterceptor(
    ctx context.Context,
    req interface{},
    info *grpc.UnaryServerInfo,
    handler grpc.UnaryHandler,
) (interface{}, error) {
    log.Println("--> unary interceptor: ", info.FullMethod)
    return handler(ctx, req)
}

func main() {
    ...
    grpcServer := grpc.NewServer(
       grpc.UnaryInterceptor(unaryInterceptor),
    )
    ...
}

````

### 3. How is authentication handled?

You would use the same code you use with REST middleware but instead of REST middleware you would use gRPC interceptors to invoke those functions. On the client side you would have a functionality to login and interceptor(s) to handle the token passing, and maybe refreshing of token. On the server side you would have method for logging in and creation of tokens, and interceptors for the parsing and validation of the token.

### 4. How is tls handled?

Since gRPC uses HTTP/2 under the hood it would just use the same TLS or even mTLS (mutual TLS). There are explicit APIs for defining this in server and client creation.

server

````go

  grpcServer := grpc.NewServer(
        grpc.Creds(tlsCredentials),
        ....
    )

````

client

````go

 cc2, err := grpc.Dial(
        *serverAddress,
        grpc.WithTransportCredentials(tlsCredentials),
        ....
    )

````

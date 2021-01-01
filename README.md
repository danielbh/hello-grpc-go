# hello-gRPC-go

### Goals:

BigtThanks to [Tech School's FREE course on dev.to](https://dev.to/techschoolguru/series/7311) which helped me greatly to learn gRPC and answer the following questions.

1. Why use gRPC instead of rest?

The first most important answer to this question is the usage of HTTP/2 in gRPC. REST can use HTTP/2 or the much more common HTTP/1.1. HTTP/2 has some great features over its predecessor. Which include, binary protocol instead of text-based, compressed headers, multiplexing on same TCP connection, and server-push. So I think it's a fairer comparison to not use this quality when comparing the two.

So HTTP/2 aside why else would I use gRPC?

Polyglot auto-code generation with [protocol buffers](https://developers.google.com/protocol-buffers/docs/overview). You can create request and payload functionality in one language [protocol buffers](https://developers.google.com/protocol-buffers/docs/overview) and automatically generate code for many popular languages: Java, Ruby, etc. Not only does this reduce code needed written by a developer, it creates a universal contract that all services can share and use. Without this it can create suprising complexity as software platforms grow.

So in summary you can deliver more performant code faster that is more maintainable. Awesome!

But what is the downside?

* One could definitively say that gRPC has limited web support

2. how is middleware handled
3. How is authentication handled?
4. How is tls handled?

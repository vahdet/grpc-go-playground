# grpc-go-playground

Just basic, tiny gRPC tries w/ Go. Based on [the Udemy course by Clement Jean](https://www.udemy.com/course/grpc-golang/).

## How do we start a gRPC server?

1. Create the `.proto` file(s), preferrably in the `api/proto/v1/` directory (see [the related Go standard project layout issue](https://github.com/golang-standards/project-layout/issues/23#issuecomment-498260932))

2. Add the line `go_package` option like:

    ```protobuf
    option go_package="github.com/vahdet/grpc-go-playground/api/proto/v1";
    ```

    Its value is project-specific and should point to the `.proto` file(s) directory.

3. Add messages (i.e. types/schemas in protocol buffers) and services (i.e. methods in the same context) to the `.proto` file(s). 

    > **Note:** The RPC methods cannot have scalar/primitive types as input or output: They have to be messages: https://stackoverflow.com/a/28919914/4636715

    > One way to work this around is to use [`wrappers`](https://github.com/protocolbuffers/protobuf/blob/main/src/google/protobuf/wrappers.proto) like:

    ```protobuf
    import "google/protobuf/wrappers.proto";

    message MyMessage {
        string value = 1;
    }

    service Service {
        rpc request (MyMessage) returns (.google.protobuf.BoolValue); 
    }
    ```

4. To generate the code, we use the `protoc` command. And we may specify the output directory, `--proto_path`, `--go_out` and `--go-grpc_out` flags to get set to generate the code in [pkg](https://github.com/golang-standards/project-layout/tree/master/pkg) folder, more specifically directly in `pkg/api/...`:

    ```shell
    $ protoc --proto_path=api/proto/v1 \
        --go_out=pkg/api --go_opt=paths=source_relative \
        --go-grpc_out=pkg/api --go-grpc_opt=paths=source_relative \
        person.proto
    ```

    We placed a script for it in the `/scripts` directory, again to comply with the community standard project layout: https://github.com/golang-standards/project-layout/tree/master/scripts. And added a command in the root level `Makefile` to run it:

    ```shell
    make codegen
    ```

    The output directory should be created before that `protoc` command is run.

    > To make the generated code right beside the `.proto` file, change the flag values as `--go_out=.` and `--go-grpc_out=.`

## See Also

- [gRPC Basics in Go](https://grpc.io/docs/languages/go/basics/)

- [Protocol Buffers Message and Field Naming Conventions](https://developers.google.com/protocol-buffers/docs/style#message_and_field_names)

- Even not implemented here for simplicity, see a couple of resources on graceful shutdown:

  - [Server.GracefulStop](https://pkg.go.dev/google.golang.org/grpc#Server.GracefulStop) method of the regular grpc package for Go.

  - SO question to use it properly: [Behavior of server.GracefulStop() in golang](https://stackoverflow.com/q/55797865/4636715)

- To have your APIs in both gRPC and RESTful style at the same time, [gRPC-Gateway](https://github.com/grpc-ecosystem/grpc-gateway) can help generate a reverse proxy in JSON format.

- And, for sure, check [Awesome gRPC](https://github.com/grpc-ecosystem/awesome-grpc) projects out there.

- The [routeguide](https://github.com/grpc/grpc-go/tree/master/examples/route_guide/routeguide) example and [server-side](https://grpc.io/docs/languages/go/basics/#server-side-streaming-rpc), [client-side](https://grpc.io/docs/languages/go/basics/#client-side-streaming-rpc) and [bidirectional streaming](https://grpc.io/docs/languages/go/basics/#bidirectional-streaming-rpc) tutorials on it.

- [Me seeking for hexagonal architecture structure on top of standard project layout](https://github.com/golang-standards/project-layout/issues/170)

- Check Google's [PubSub proto file](https://github.com/googleapis/googleapis/blob/master/google/pubsub/v1/pubsub.proto) as a helpful real-world example.

- _Out of Context_: The correct order for `defer` and checking `err != nil`: https://stackoverflow.com/a/39287459/4636715
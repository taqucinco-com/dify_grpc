/*
 * Copyright 2024, gRPC Authors All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

import ArgumentParser
import GRPCNIOTransportHTTP2
//import GRPCProtobuf

struct GreeterService: AsyncParsableCommand {
    static let port: Int = 50051
    
    func runSayHello(name: String) async throws -> String {
        try await withThrowingDiscardingTaskGroup { group in
            let client = GRPCClient(
                transport: try .http2NIOPosix(
                    target: .ipv4(host: "127.0.0.1", port: GreeterService.port),
                    config: .defaults(transportSecurity: .plaintext)
                )
            )

            group.addTask {
                try await client.run()
            }

            defer {
                client.beginGracefulShutdown()
            }

            let greeter = Helloworld_GreeterClient(wrapping: client)
            let reply = try await greeter.sayHello(.with { $0.name = name })
            print(reply.message)
        
            return reply.message
        }
    }
    
    func runSayHelloAgain(name: String) async throws -> AsyncStream<String> {
        return AsyncStream { continuation in
            Task {
                try await withThrowingDiscardingTaskGroup { group in
                    let client = GRPCClient(
                        transport: try .http2NIOPosix(
                            target: .ipv4(host: "127.0.0.1", port: GreeterService.port),
                            config: .defaults(transportSecurity: .plaintext)
                        )
                    )
        
                    group.addTask {
                        try await client.run()
                    }
        
                    defer {
                        client.beginGracefulShutdown()
                    }
        
                    let greeter = Helloworld_GreeterClient(wrapping: client)
                    var requrest = Helloworld_HelloRequest()
                    requrest.name = name
        
//                    func converter () async throws -> ClientResponse.Stream<Helloworld_HelloResponse> {
//                        try await withCheckedThrowingContinuation { continuation in
//                            Task {
//                                try await greeter.sayHelloAgain(requrest) { response in
//                                    continuation.resume(returning: response)
//                                }
//                            }
//                        }
//                    }
        
                    try await greeter.sayHelloAgain(requrest) { response in
                        for try await reply in response.messages {
                            print(reply.message)
                            continuation.yield(reply.message)
                        }
                    }
                    continuation.finish()
                }
            }
        }
    }
}

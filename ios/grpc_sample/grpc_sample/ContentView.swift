//
//  ContentView.swift
//  grpc_sample
//
//  Created by takuya.sudo on 2024/09/29.
//

import SwiftUI

struct ContentView: View {
    var body: some View {
        VStack {
            Image(systemName: "globe")
                .imageScale(.large)
                .foregroundStyle(.tint)
            Text("Hello, world!")
            Button {
                Task {
                    let service = GreeterService()
                    let reply = try await service.runSayHello(name: "taro")
                    print("run Say Hello result: \(reply)")
                }
            } label: {
                Text("Say Hello")
            }
            Button {
                Task {
                    let service = GreeterService()
                    for try await reply in try await service.runSayHelloAgain(name: "taro") {
                        print("run Say Hello Again result: \(reply)")
                    }
                    print("run Say Hello Again Recieved")
                }
            } label: {
                Text("Say Hello Again")
            }
        }
        .padding()
    }
}

#Preview {
    ContentView()
}

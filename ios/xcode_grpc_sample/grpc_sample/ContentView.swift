//
//  ContentView.swift
//  grpc_sample
//
//  Created by takuya.sudo on 2024/09/29.
//

import SwiftUI

struct ContentView: View {
    @State private var replay: String = "Text.."
    @State var name = ""
    @FocusState var isFocused: Bool
    
    var body: some View {
        VStack {
            Text(replay)
            Spacer()
            TextField("名前を入力してください", text: $name)
                .focused($isFocused)
                .padding(.all, 8.0)
            Button {
                Task {
                    let service = GreeterService()
                    let message = try await service.runSayHello(name: name)
                    replay = message
                }
            } label: {
                Text("Say Hello With Unary")
            }
            .padding(.all, 8.0)
            Button {
                Task {
                    replay = ""
                    let service = GreeterService()
                    for try await message in try await service.runSayHelloAgain(name: name) {
                        replay = replay + "\(message)"
                    }
                }
            } label: {
                Text("Say Hello With Server Streaming")
            }
            .padding(.all, 8.0)
        }
        .padding()
    }
}

#Preview {
    ContentView()
}

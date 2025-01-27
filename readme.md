# Micro-service Driven Chat App

## Overview

This project is a collection of micro-services to serve a chat application. Each service handles a specific part of the chat system, making it easier to maintain and scale independently.

## Built With

- GoLang (1.18+)
- Python (3.11)
- TypeScript (20)

## Deployable Services

- [chat_service_go](chat_service_go/readme.md): Handles user requests, talking to other services, and streaming responses
- [chat_service_python](chat_service_python/readme.md): Manages the database and chat history
- [chat_service_typescript](chat_service_typescript/readme.md): The frontend where users interact

## Why?

#### At A High Level...
I just discovered how great GoLang is for servicing apps that you intend to scale. The way it handles concurrent requests and its performance benefits made me rethink how I structure my applications.

#### From A Place Of Honesty...
This evening, everything clicked as I was digging through the back-end of Ollama. I asked myself: "Why am I forcing myself to serve web-apps using Python? Especially when the Python application would just be a wrapper for the GoLang application (Ollama)?". After seeing how many successful services use GoLang (Tailscale, Paypal, Netflix), I figured I'd lean into this micro-service approach instead of trying to do everything in Python.

#### About The Architecture
Each service is designed to run on its own server instance, though for now they're all running locally. The idea is that as the app grows, we can scale each part independently:
- Go service handles the traffic, routing, and streaming
- Python manages the database stuff
- TypeScript/React gives us a nice frontend

#### Micro-Service, Stateless, API, MVC, and all those buzzwords...
Blogs, articles, tutorials, and the like all throw around these buzzwords; but what do they really mean? If you know the answer, there's no need to read this. If you don't and feel overwhelmed, start with APIs. If APIs don't make too much sense, learn Model-View-Controllers. If you're still lost, learn about OOP or build a simple app with multiple components in a language you're comfortable with. As you build more, you'll start to see how these buzzwords are more or less the same concept applied to different applications or components. For example, in this chat app, the Python service follows MVC patterns to handle data, while exposing an API that the Go service can call - showing how these concepts naturally flow together in practice.

This project is my attempt to put these concepts into practice, even if it's just a simple chat app for now.


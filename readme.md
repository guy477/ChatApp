# Micro-service Driven Chat App

## Overview

This project is collection of micro-services designed to serve open-source LLMs to a UI and maintain histories in databases. Each service handles a specific part of the chat system, making it easier to maintain and scale independently.

## Built With

- GoLang (1.18+)
- Python (3.11)
- HTML/CSS/JS

## Deployable Services

- [chat_service_go](chat_service_go/readme.md): Handles user requests, talking to other services, and streaming responses
- [chat_service_python](chat_service_python/readme.md): Manages the database and chat history
- [chat_service_html](chat_service_html/readme.md): The frontend where users interact

## Why?

### At A High Level...
I just discovered how great GoLang is for servicing apps that you intend to scale. The way it handles concurrent requests and its performance benefits made me rethink how I structure my applications.

### From A Place Of Honesty...
This evening, everything clicked as I was digging through the back-end of Ollama. I asked myself: "Why am I forcing myself to serve web-apps using Python? Especially when the Python application would just be a wrapper for the GoLang application (Ollama)?". After seeing how many successful services use GoLang (Tailscale, Paypal, Netflix), I figured I'd lean into this micro-service approach instead of trying to do everything in Python.

### About The Architecture...
Each service is designed to run on its own server instance, though for now they're all running locally. The idea is that as the app grows, we can scale each part independently:
- Go service handles the traffic, routing, and streaming
- Python manages the database stuff
- HTML gives us a simple, classic frontend

### My Understanding Of Software Architecture Patterns...
Modern software development relies heavily on established patterns and principles that promote modularity, scalability, and maintainability. While terms like "microservices," "stateless," and "MVC" might seem like buzzwords, they represent fundamental architectural concepts that solve specific problems:

- **APIs (Application Programming Interfaces)** define standardized ways for software components to communicate, whether between microservices, frontend-backend, or third-party integrations.

- **MVC (Model-View-Controller)** separates concerns in applications:
  - Models handle data and business logic
  - Views present information to users
  - Controllers coordinate between models and views

- **Microservices** break down complex applications into independent, specialized services that communicate via APIs. This enables:
  - Independent scaling and deployment
  - Technology stack flexibility per service
  - Isolated failure domains

In this project, these patterns work together naturally. For example:
- The Python service implements MVC internally to manage chat data (separation of concerns)
- It exposes this functionality through a REST API (interface)
- The Go service consumes this API as part of the larger microservice architecture

The common thread is interfaces - well-defined boundaries between components that enable modularity and flexibility. Whether you're designing class methods, service APIs, or entire system architectures, thinking in terms of clean interfaces leads to more maintainable and scalable software.

This project is my attempt to put these concepts into practice, even if it's just a simple chat app for now.


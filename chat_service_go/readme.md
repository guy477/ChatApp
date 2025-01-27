# Setup

## Install GoLang on Linux (1.18+; common CPU architectures)
```sh
sudo apt-get install golang-go
```

## Run Service
```sh
go run main.go
```

## Other Requirements
- Install [Ollama](https://ollama.com/) on a reachable device. 
    - If Ollama is hosted on the same machine as the Go service, no changes are needed.
    - If Ollama is hosted on a different machine, be sure to configure the remote machine with a reverse proxy to the Ollama API and then update the `OllamaURL` in `config.go` to point to the remote machine.


# Future

- [ ] Move Ollama to a remote machine
- [ ] Wholistically integrate Ollama API for inference
    - [ ] Enable embeddings
    - [ ] Enable image
- [ ] Add load-balancing for Ollama systems.
- [ ] Add auth endpoints
# Pure Golang AI Backend
---
This repo is meant to be a learning exercise for to handle making API requests in the pure Go standard library. It was fun.

## Instructions
Start the repo with
```bash
go run .
```

Open another terminal and send a post request to `chat/`

```bash
curl -X POST -d "hi" http://localhost:8080/chat
```


## Features
- Pure Go standard library, super fast
- Output parsing from Gemini by loading JSON into internal structs

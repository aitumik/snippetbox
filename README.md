# Snippetbox
Learning Go Web app development with a project.

## Setup
Install the latest version of Go

Mac
```
brew install go
```

Ubuntu
```
sudo apt install go
```

Clone the repository
> Note that to clone the repository you need to have git installed
```
git clone https://github.com/aitumik/snippetbox
```

Change directory
```
cd snippetbox
```

Run the application
> Note you need to enable setopts on your terminal
```
go run cmd/web/!(*_test).go
```

Or build the app
```
go build  cmd/web/* -o snippetbox
```

Run the application
```
./snippetbox
```


## Docker

Make sure you have installed docker
```
cd snippetbox
```

Build the image
```
docker build -t snippetbox
```

Run the image
```
docker run -p 4000:4000 snippetbox
```
Open your browser and visit https://localhost:4000

> Notice the https in the above URL

### Todo
- [ ] Oauth2 for authentication
- [ ] Redis for caching (we mostly do reads,writes are minimum)
- [ ] Kibana for analytics and visualization of data

### In Progress
- [ ] Add search capabilities by integrating with elasticsearch
- [ ] Switch to postgres database

### Done ✓
- [x] Cache templates for faster rendering
- [x] Request logging middleware implemented
- [x] Panic recovery middleware
- [x] Add middlewares
- [x] Dockerize the application

# Tools
* Postgres
* Elasticsearch

# Coverage & Benchmark tests

# Contributing
Create a pull request to be able to contribute to this project





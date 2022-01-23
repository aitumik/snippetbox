# Snippetbox
Learning Go Web app development with a project
> This application is capable of searching large number of texts since its uses lucene based elasticsearch

## Setup
#### Install Go

Install go for Mac OS(Catalina is what I used)
```
brew install go
```

Install go for Ubuntu(Tested with ubuntu 20.4)
```
sudo apt install go
```

Clone the repository
```
git clone https://github.com/aitumik/snippetbox
```

Change directory
```
cd snippetbox
```

Run the application
```
go run cmd/web/*
```

Or build the app
```
go build  .
```

Run the application
```
./snippetbox
```

### Todo
- [ ] Add search capabilities by integrating with elasticsearch
- [ ] Oauth2 for authentication
- [ ] Redis for caching (we mostly do reads,writes are minimum)

### In Progress
- [ ] Add middlewares
- [ ] Dockerize the application

### Done ✓
- [x] Cache templates for faster rendering
- [x] Request logging middleware implemented
- [x] Panic recovery middleware

## Tools
* SQLLite
* Go
* Javascript
* HTML and CSS 

# Contributing
If you wish to contribute to this project please feel free to create a PR and tag me (https://github.com/aitumik)
on the pull request. I will try my best to review them and merge where the standards have been met.

> NOTE  Remember to write tests before you create a PR otherwise your PR will be discarded


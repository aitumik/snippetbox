# Snippetbox
Learning Go Web app development with a project
> This application is capable of searching large number of texts since its uses lucene based elasticsearch

## Setup
Install Go

Install go for Mac OS(Catalina is what I used)
`brew install go`

Install go for Ubuntu(Tested with ubuntu 20.4)
`sudo apt install go`

### Todo
- [ ] Add search capabilities by integrating with elasticsearch
- [ ] Oauth2 for authentication
- [ ] Redis for caching (we mostly do reads,writes are minimum)

### In Progress
- [ ] Add middlewares
- [ ] Dockerize the application

### Done ✓
- [x] Cache templates for faster rendering

Running the app
```
cd snippetbox
go build  .
./snippetbox
```

## Tools
* SQLLite
* Go
* Javascript
* HTML and CSS 

# Contributing
If you wish to contribute to this project please feel free to create a PR and tag me (https://github.com/aitumik)
on the pull request. I will try my best to review them and merge where the standards have been met.

> NOTE  Remember to write tests before you create a PR otherwise your PR will be discarded


# Snippetbox
Learning Go Web app development with a project.

## Docker Setup

> Note that to clone the repository you need to have git installed
```
git clone https://github.com/aitumik/snippetbox
```

Make sure you have installed docker
```
cd snippetbox
```

Build and Run
```
docker-compose up -d
```

[Snippetbox](http://localhost:5000)
![Snippetbox Screenshot](/screenshots/image.png "screenshot of the homepage").
[Adminer](http://localhost:8080)
![Adminer Screenshot](/screenshots/adminer.png "screenshot of adminer").


### Todo
- [ ] Oauth2 for authentication
- [ ] Redis for caching (we mostly do reads,writes are minimum)
- [ ] Kibana for analytics and visualization of data
- [ ] Setup CI/CD
- [ ] Add coverage tests

### In Progress
- [ ] Add search capabilities by integrating with elasticsearch

### Done ✓
- [x] Cache templates for faster rendering
- [x] Request logging middleware implemented
- [x] Panic recovery middleware
- [x] Add middlewares
- [x] Dockerize the application

# Tools
* Postgres
* Elasticsearch

# Coverage



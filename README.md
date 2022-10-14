# Snippetbox
Snippetbox is a golang application for curating snippets

## Docker Setup

> Note that to clone the repository you need to have git installed
```
git clone https://github.com/aitumik/snippetbox
```

Make sure you have installed docker
```
cd snippetbox
```

Here is the `docker-compose.yaml`
```yaml
# Application
  snippetbox:
    container_name: snippetbox
    build:
      context: .
      dockerfile: Dockerfile
    ports:
      - "5000:5000"
    depends_on:
      - postgres
    env_file: .env
    volumes:
      - ./:/snippetbox
# Database
  postgres:
    env_file: .env
    image: "postgres"
    hostname: "postgres"
    ports:
        - "5432:5432"
    volumes:
        - pgdata:/var/lib/postgresql/data
# Adminer
  admner:
    container_name: admner
    image: dockette/adminer
    restart: always
    ports:
      - "8085:80"
    depends_on:
      - postgres
# Volumes
volumes:
  pgdata:
    driver: local
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


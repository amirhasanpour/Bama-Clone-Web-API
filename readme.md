# Car Sales Management Web API (Inspired by [bama.ir](https://bama.ir))
A scalable, observable, and modular RESTful API for managing a car sales platform. Built with Go using the Gin framework, the architecture follows clean coding practices and is containerized using Docker Compose.

---

## System Design Diagram

<p align="center"><img src='/files/images/system_diagram.png' alt='System Design Diagram' /></p>

---

## Features

- **JWT-based authentication and authorization for securing protected routes**
- **Car sales management including creating, updating, deleting, and browsing listings**
- **Clean architecture using Gin framework (router → handler → service → contracts → infra)**
- **Input validation using validator for secure and strict endpoint validation**
- **Environment and configuration management via viper with support for .env and YAML**
- **PostgreSQL as the main relational database engine**
- **PgAdmin as a visual tool for inspecting and managing the PostgreSQL database**
- **GORM as the ORM layer for interacting with PostgreSQL using models and struct-based queries**
- **Docker Compose for orchestrating all services like DB, Redis, Elasticsearch, Prometheus, and more**
- **Redis caching to reduce database load and enhance performance for hot data**
- **Prometheus for real-time metrics collection and monitoring**
- **Grafana dashboards for visualizing performance and system metrics**
- **Centralized logging pipeline using Elasticsearch, Filebeat, and Kibana**
- **Structured logging using both zap and zerolog for performant, JSON-formatted logs**
- **Auto-generated API documentation with Swagger UI for easy testing and development**

---

## Used Tools:

1. [Gin framework](https://github.com/gin-gonic/gin)
2. [JWT authentication and authorization](https://github.com/golang-jwt/jwt)
3. [Redis](https://github.com/redis/redis)
4. [Elasticsearch](https://github.com/elastic/elasticsearch)
5. [Beat](https://github.com/elastic/beats)
6. [Kibana](https://github.com/elastic/kibana)
7. [Postgresql](https://github.com/postgres/postgres)
8. [PgAdmin](https://github.com/pgadmin-org/pgadmin4)
9. [Prometheus](https://github.com/prometheus/prometheus)
10. [Grafana](https://github.com/grafana/grafana)
11. [Validator](https://github.com/go-playground/validator)
12. [Viper](https://github.com/spf13/viper)
13. [Zap](https://github.com/uber-go/zap)
14. [Zerolog](https://github.com/rs/zerolog)
15. [Gorm](https://github.com/go-gorm/gorm)
16. [Swagger](https://github.com/swaggo/swag)
17. [Docker Compose](https://github.com/docker/compose)

---

## How to run?

### 1- Start Dependencies on Docker

cd to the docker directory and run this command:

```bash
docker compose -f "docker/docker-compose.yaml" up -d setup elasticsearch kibana filebeat postgres pgadmin redis prometheus node-exporter alertmanager grafana
```

### 2- Install Swagger CLI

on the root level of project run this command:

```bash
go install github.com/swaggo/swag/cmd/swag@latest
```

### 3- Install Go Dependencies

on the root level of project run this command:

```bash
go mod download
```

### 4- Run the Application

cd to cmd directory:

```bash
cd src/cmd
```

and then run:

```bash
go run main.go
```

### 5- Visit the Application

#### Web API run on  [http://localhost:10000](http://localhost:10000)

#### Swagger on  [http://localhost:10000/swagger/index.html#/](http://localhost:10000/swagger/index.html#/)

```bash
Token Url: http://localhost:10000/api/v1/users/login-by-username
Username: admin
Password: 12345678
```

#### Kibana on  [http://localhost:5680](http://localhost:5680)

```bash
Username: elastic
Password: changeme
```

#### Prometheus on  [http://localhost:30090](http://localhost:30090)

#### Grafana on  [http://localhost:3000](http://localhost:3000)

```bash
Username: admin
Password: foobar
```

#### PgAdmin on  [http://localhost:8888](http://localhost:8888)

```bash
Username: amirhossinhp10@gmail.com
Password: admin
```

#### Postgres Server info:

```bash
Host: postgres_container
Port: 5432
Username: postgres
Password: admin
```

### 6- Stop all Services

```bash
docker compose -f "docker/docker-compose.yaml" down
```

---

## Project preview

### Swagger

<p align="center"><img src='/files/images/swagger.png' alt='swagger preview' /></p>

### Kibana

<p align="center"><img src='/files/images/kibana.png' alt='kibana preview' /></p>

### Grafana

<p align="center"><img src='/files/images/grafana.png' alt='grafana preview' /></p>

### PgAdmin

<p align="center"><img src='/files/images/pgadmin.png' alt='pgadmin preview' /></p>
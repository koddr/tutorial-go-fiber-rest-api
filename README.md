# 📖 Tutorial: Build a RESTful API on Go

Fiber, PostgreSQL, JWT and Swagger docs in isolated Docker containers.

## Quick start

1. Rename `.env.example` to `.env` and fill it with your environment values.
2. Install [Docker](https://www.docker.com/get-started) and [migrate](https://github.com/golang-migrate/migrate) tool for applying migrations.
3. Run project by this command:

```bash
make docker.run

# Process:
#   - Generate API docs by Swagger
#   - Create a new Docker network for containers
#   - Build and run Docker containers (Fiber, PostgreSQL)
#   - Apply database migrations (using github.com/golang-migrate/migrate)
```

4. Go to API Docs page (Swagger): [localhost:5000/swagger/index.html](http://localhost:5000/swagger/index.html)

![Screenshot](https://user-images.githubusercontent.com/11155743/111972767-dd16e400-8b0e-11eb-8ba1-98c648f56a5a.png)

## ⚠️ License

MIT &copy; [Vic Shóstak](https://github.com/koddr) & [True web artisans](https://1wa.co/).

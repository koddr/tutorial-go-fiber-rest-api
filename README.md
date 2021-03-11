# 📖 Tutorial: Build a RESTful API on Go

Fiber, PostgreSQL, JWT auth and automatically generated Swagger docs in isolated Docker containers.

1. Rename `.env.example` to `.env` and fill it with your environment values.
2. Run project by this command:

```bash
make docker.run

# Process:
#   - Generate API docs
#   - Create a new Docker network
#   - Build and run Docker containers (Fiber, PostgreSQL)
#   - Apply database migrations (by github.com/golang-migrate/migrate)
```

3. Go to API Docs page (Swagger): [localhost:5000/swagger/index.html](http://localhost:5000/swagger/index.html).

## ⚠️ License

MIT &copy; [Vic Shóstak](https://github.com/koddr) & [True web artisans](https://1wa.co/).

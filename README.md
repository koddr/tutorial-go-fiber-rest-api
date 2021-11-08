# 📖 Tutorial: Build a RESTful API on Go

Fiber, PostgreSQL, JWT and Swagger docs in isolated Docker containers.

👉 The full article is published on **March 22, 2021**, on Dev.to: https://dev.to/koddr/build-a-restful-api-on-go-fiber-postgresql-jwt-and-swagger-docs-in-isolated-docker-containers-475j

![fiber_cover_gh](https://user-images.githubusercontent.com/11155743/112001218-cf258b00-8b2f-11eb-9c6d-d6c38a09af86.jpg)

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

4. Go to your API Docs page: [127.0.0.1:5000/swagger/index.html](http://127.0.0.1:5000/swagger/index.html)

![Screenshot](https://user-images.githubusercontent.com/11155743/111976684-f15ce000-8b12-11eb-871a-8d32465900fe.png)

## P.S.

If you want more articles like this on this blog, then post a comment below and subscribe to me. Thanks! 😘

And, of course, you can support me by donating at [LiberaPay](https://liberapay.com/koddr/donate). _Each donation will be used to write new articles and develop non-profit open-source projects for the community._

[![Support author at LiberaPay](https://dev-to-uploads.s3.amazonaws.com/uploads/articles/zq8442cqyjq2i1jdeay8.png)](https://liberapay.com/koddr/donate)

## ⚠️ License

MIT &copy; [Vic Shóstak](https://github.com/koddr) & [True web artisans](https://1wa.co/).

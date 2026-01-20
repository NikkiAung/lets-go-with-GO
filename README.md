## Project Structure

go-fundamentals/  
 ├── main.go # Entry point, server setup  
 ├── go.mod # Module: github.com/NikkiAung/go-fundmentals  
 └── internal/  
 ├── app/app.go # Application struct, logger, health check  
 ├── api/post_handler.go # HTTP handlers for posts  
 └── routes/routes.go # Route definitions (chi router)

## Application Flow

main.go  
 ↓  
 app.NewApplication() → creates Logger + PostHandler  
 ↓  
 routes.SetUpRoutes(app) → chi router with routes  
 ↓  
 http.Server.ListenAndServe()  
 ├── GET /health → HealthCheck()  
 ├── GET /posts/{id} → HandleGetPostByID()
└── POST /posts → HandleCreatePost()

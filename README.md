gofiber-app/
│
├── cmd/
│ └── server/
│ └── main.go # Entry point of the application
│
├── config/
│ └── config.go # Configuration and environment variables
│
├── internal/
│ ├── app/
│ │ ├── models/
│ │ │ └── task.go # Model definitions
│ │ ├── handlers/
│ │ │ └── task.go # Request handlers (controllers)
│ │ ├── services/
│ │ │ └── task.go # Business logic and services
│ │ ├── repositories/
│ │ │ └── task.go # Database interactions (data access layer)
│ │ ├── middleware/
│ │ │ └── auth.go # Middleware for authentication, logging, etc.
│ │ └── routes/
│ │ └── routes.go # API route definitions
│ │
│ ├── db/
│ │ └── db.go # Database setup and connections
│ │
│ ├── utils/
│ │ └── utils.go # Utility functions
│ │
│ └── config/
│ └── config.go # Configuration settings
│
├── pkg/
│ └── utils/
│ └── helpers.go # Shared utility functions or packages
│
├── .env # Environment variables
├── .air.toml # Air configuration for live reloading
├── go.mod # Go module file
├── go.sum # Go dependencies file
└── README.md # Project documentation

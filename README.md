# URL-shortener
___
This project is a REST API application that allows you to shorten links to a specified size (from 5 to 100 characters).

# The dependencies that I used
___
1. `Zap` (Logger)
2. `Gin` (Router)
3. `Viper` (config)
4. `Godotenv` (.Env)
5. `Gorm` (DB)
6. `Postgres` (DB)

# Endpoint's
___
1. POST `/` - request to create an abbreviated link. Body Request: `{Root: "", Size: 0}`
2. GET `/:alias` - shortened link with further redirection

set positional-arguments
set fallback

# List all just tasks
default:
    @just --list

recreate_db:
    go run ./cli/main.go scraping
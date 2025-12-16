# skorgator
CLI Blog Aggregator 

A command-line blog aggregator written in Go that fetches and manages RSS feeds.

> Note for Boot.dev users:
> The original course binary name is `gator`. In this project the binary is called
> `skorgator`, but its behavior and commands are equivalent to the `gator` CLI
> described in the course.

## Prerequisites

### Install Go

**Linux/macOS:**
```bash
# Download and install Go (version 1.21 or later)
wget https://go.dev/dl/go1.21.5.linux-amd64.tar.gz
sudo rm -rf /usr/local/go && sudo tar -C /usr/local -xzf go1.21.5.linux-amd64.tar.gz
export PATH=$PATH:/usr/local/go/bin
```

**Or use your package manager:**
```bash
# Ubuntu/Debian
sudo apt update
sudo apt install golang-go

# macOS (Homebrew)
brew install go

# Arch Linux
sudo pacman -S go
```

Verify installation:
```bash
go version
```

### Install PostgreSQL

**Ubuntu/Debian:**
```bash
sudo apt update
sudo apt install postgresql postgresql-contrib
sudo systemctl start postgresql
sudo systemctl enable postgresql
```

**macOS (Homebrew):**
```bash
brew install postgresql
brew services start postgresql
```

**Arch Linux:**
```bash
sudo pacman -S postgresql
sudo -u postgres initdb -D /var/lib/postgres/data
sudo systemctl start postgresql
sudo systemctl enable postgresql
```

Create a database:
```bash
sudo -u postgres psql
CREATE DATABASE skorgator;
CREATE USER skorgator_user WITH PASSWORD 'your_password';
GRANT ALL PRIVILEGES ON DATABASE skorgator TO skorgator_user;
\q
```

## Setup

### Option 1: Install with go install (Recommended)

```bash
go install github.com/skorgum/skorgator@latest
```

This will install the `skorgator` binary to your `$GOPATH/bin` directory (usually `~/go/bin`). Make sure this directory is in your PATH.

### Option 2: Build from source

### 1. Clone the repository
```bash
git clone https://github.com/yourusername/skorgator.git
cd skorgator
```

### 2. Install dependencies
```bash
go mod download
```

### 3. Configure the application

Create a `gatorconfig.json` file in the project root:

```json
{
  "db_url": "postgres://skorgator_user:your_password@localhost:5432/skorgator?sslmode=disable",
  "current_user_name": ""
}
```

Replace `your_password` with the password you set when creating the PostgreSQL user.

### 4. Run database migrations

If you're using migrations (e.g., with goose or similar), run:
```bash
# Example with goose
goose -dir sql/schema postgres "your_connection_string" up
```

### 5. Build the application
```bash
go build -o skorgator
```

## Configuration

Before using skorgator, create a `gatorconfig.json` file in the project root (or your working directory):

```json
{
  "db_url": "postgres://skorgator_user:your_password@localhost:5432/skorgator?sslmode=disable",
  "current_user_name": ""
}
```

Replace `your_password` with the password you set when creating the PostgreSQL user.

## Usage

### Login
Set the current user:
```bash
skorgator login <username>
# Or if built from source:
./skorgator login <username>
```

### Register a new user
```bash
skorgator register <username>
```

### Add a feed
```bash
skorgator addfeed <feed_url> <feed_name>
```

### List all feeds
```bash
skorgator feeds
```

### Follow a feed
```bash
skorgator follow <feed_url>
```

### List feeds you're following
```bash
skorgator following
```

### Unfollow a feed
```bash
skorgator unfollow <feed_url>
```

### Browse posts
```bash
skorgator browse [limit]
```

### Start aggregation
Fetch new posts from all feeds:
```bash
skorgator agg <time_between_requests>
```

### Reset database
Delete all users:
```bash
skorgator reset
```

### List all users
```bash
skorgator users
```

## Configuration File

The `gatorconfig.json` file stores:
- `db_url`: PostgreSQL connection string
- `current_user_name`: The currently logged-in user

This file is automatically updated when you use the `login` command.

## Development

Run without building:
```bash
go run . <command> [args]
```

Run tests:
```bash
go test ./...
```

## License

See [LICENSE](LICENSE) file for details.

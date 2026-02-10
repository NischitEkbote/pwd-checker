## 🐳 Setup & Run (Using Docker)

### 1. Clone the Repository

```bash
git clone https://github.com/NischitEkbote/pwd-checker.git
cd pwd-checker
```

This downloads the project and navigates into the project directory.

### 2. Build the Docker Image

```bash
docker build -t pwd-checker .
```

Reads the Dockerfile in the current directory

Builds the Go binary inside a Docker container

Creates a Docker image named pwd-checker

### 3. Run the Password Checker

```bash
docker run --rm pwd-checker --pass "Hello123!"
```

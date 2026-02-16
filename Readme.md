# Damsharaz.io

A real-time backend system built with Go.

---

## Getting Started

### 1 Install Go

Make sure Go (1.22+) is installed:

```bash
go version
```

---

### 2️ Build the Project

From the root directory:

```bash
cd scripts
make
```

The binary will be generated inside:

```
/bin/damsharaz
```

---

## Running Tests

After making changes:

```bash
cd tests
go test
```

Or from the project root (recommended):

```bash
go test ./...
```

---

## Project Structure

```
├── Readme.md
├── bin
│   └── damsharaz            # Compiled binary
├── cmd
│   └── main.go              # Application entry point
├── dump.rdb                 # Redis dump file
├── go.mod
├── go.sum
├── internal
│   ├── config               # Configuration logic
│   ├── db                   # Database/Redis setup
│   ├── modules              # Feature modules
│   ├── pkg                  # Shared utilities
│   ├── server               # HTTP/WebSocket server
│   └── types                # Custom types/interfaces
├── logs.log                 # Application logs
├── scripts
│   └── Makefile             # Build scripts
└── tests
    └── utils_test.go        # Test files
```

---

## Development Workflow

1. Make changes in `internal/`
2. Add/update tests in `tests/`
3. Run:

```bash
go test ./...
```

4. Rebuild if needed:

```bash
cd scripts
make
```

---

## Contributing

Contributions are welcome.

### How to Contribute

1. Fork the repository
2. Create a feature branch:

```bash
git checkout -b feature/your-feature-name
```

3. Make your changes
4. Add tests if applicable
5. Run all tests:

```bash
go test ./...
```

6. Commit clearly:

```bash
git commit -m "Add: short description of change"
```

7. Push and open a Pull Request

---

### Contribution Guidelines

* Follow existing project structure
* Keep logic inside `internal/`
* Write tests for new features
* Avoid breaking existing APIs
* Ensure no unused imports or lint issues

---

## Requirements

* Go 1.22+
* Redis (if running full system)

---


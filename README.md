# SOMAS 2020

Main repository for SOMAS2020 Coursework.

- [Setup & Rules](./docs/SETUP.md)
- [Infra Info](./docs/INFRA.md)

## Running code
```bash
# Approach 1
go run . # Linux and macOS: Use `sudo go run .` if you encounter any "Permission denied" errors.

# Approach 2
go build # build step
./SOMAS2020 # SOMAS2020.exe if you're on Windows. Use `sudo` on Linux and macOS as Approach 1 if required.
```

### Output
After running, the `output` directory will contain the output of the program.
- `output.json`: JSON file containing the game's historic states and configuration.
- `log.txt`: logs of the run

### Visualisation Website
See [`website/README.md`](website/README.md)

## Testing
```bash
go test ./...
```

## Structure

### [`docs`](docs)
Important documents pertaining to codebase organisation, code conventions and project management. Read before writing code.

### [`internal`](internal)
Internal SOMAS2020 packages. Most development occurs here, including client and server code.

- [`clients`](internal/clients)
Individual team code goes into the respective folders in this directory.

- [`common`](internal/common)
Common utilities, or system-wide code such as game specification etc.

- [`server`](internal/server)
Self-explanatory.

- [`logger`](internal/logger)
Logger for the application.

### [`pkg`](pkg)
More generic packages dealing with general use-cases, such as system-related or file-operation utilities.

### [`website`](website)
Source code for visualisation website.

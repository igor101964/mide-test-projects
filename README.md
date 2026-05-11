# mide_test_projects

A collection of small, self-contained example projects in multiple programming languages (C, C++, Rust, Go, Python, Bash, and Lua). This repository is intended for testing build systems, development environments, code editors, and tooling integrations (such as IDEs, language servers, and CI pipelines).

Each subproject is minimal and focuses on being easy to build and run, making this repository useful as a reference or test bed.

---

## Features

- **Multi-language examples**
  - C (single-file and multi-file with `Makefile`)
  - C++ (single-file)
  - Rust (single-file and Cargo project)
  - Go (single-file and Go module project)
  - Python (single-file script)
  - Bash (single-file script)
  - Lua (single-file script)

- **Both “single file” and “project” layouts**
  - Single-file examples that compile or run directly
  - Small but complete project layouts using:
    - `Makefile` for C
    - `Cargo.toml` for Rust
    - `go.mod` for Go

- **Simple “hello world” style programs**
  - Ideal for testing:
    - Compiler installation
    - PATH and environment configuration
    - Editor/build integration
    - Debugging and profiling tools
    - Static analysis and linters

- **Cross-platform friendly**
  - Source code is POSIX- and standards-oriented where possible
  - Build steps are kept minimal and conventional

---

## Requirements / Dependencies

You do not need everything installed to use this repository; you can selectively install tools based on which projects you want to build or run.

### General

- A POSIX-like environment for shell commands (Linux, macOS, WSL, or similar)
- Git (to clone the repository)

### Language-specific tools

- **C**
  - C compiler: `gcc` or `clang`
  - `make` for `c_project`

- **C++**
  - C++ compiler: `g++` or `clang++`

- **Rust**
  - Rust toolchain (via [rustup](https://rustup.rs/)):
    - `rustc`
    - `cargo`

- **Go**
  - Go toolchain (see [https://go.dev/dl/](https://go.dev/dl/))
  - `GOPATH` and `GOROOT` configured (depending on Go version and setup)

- **Python**
  - Python 3.x (e.g., `python3`)

- **Bash**
  - `bash` shell

- **Lua**
  - Lua interpreter (e.g., `lua` or `lua5.4` depending on distribution)

---

## Installation

### 1. Clone the repository

```bash
git clone https://github.com/igor101964/mide_test_projects.git
cd mide_test_projects
```

(Replace the URL with the actual repository location.)

### 2. Install language toolchains as needed

Examples for some platforms:

**Ubuntu / Debian (apt-based):**

```bash
sudo apt update
sudo apt install build-essential gcc g++ make \
                 python3 \
                 golang \
                 lua5.4 \
                 bash
```

Install Rust via `rustup`:

```bash
curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh
source "$HOME/.cargo/env"
```

**macOS (Homebrew):**

```bash
brew install gcc make python go lua rust
```

Adjust these commands based on your OS and package manager.

---

## Usage

Below are example commands to build and run each subproject. All commands assume you are in the repository root (`mide_test_projects/`).

### C (single file)

```bash
cd c_single

# Compile
gcc -o hello hello.c

# Run
./hello
```

### C (multi-file project with Makefile)

```bash
cd c_project

# Build using Makefile
make

# Run resulting binary
./c_project

# Clean build artifacts (if supported by Makefile)
make clean
```

### C++ (single file)

```bash
cd cpp_single

# Compile
g++ -o hello hello.cpp

# Run
./hello
```

### Rust (single-file binary)

```bash
cd rust_single

# Compile
rustc hello.rs

# Run
./hello
```

### Rust project (Cargo)

```bash
cd rust_project

# Build
cargo build

# Run (debug build)
cargo run

# Or run the compiled binary directly
./target/debug/rust_project
```

*(The final binary name may differ depending on `Cargo.toml` package name.)*

### Go (single file)

```bash
cd go_single

# Run directly
go run hello.go

# Or build then run
go build -o hello hello.go
./hello
```

### Go project (Go module)

```bash
cd go_project

# Run
go run .

# Or build then run
go build -o go_project .
./go_project
```

### Python (single file)

```bash
cd py_single

python3 hello.py
```

### Bash (single file)

```bash
cd bash_single

# Ensure the script is executable
chmod +x hello.sh

# Run
./hello.sh
```

### Lua (single file)

```bash
cd lua_single

lua hello.lua
```

---

## Example Output

Most of the programs are simple “hello world” style examples. Typical output will look like:

```text
Hello, world!
```

Some multi-file projects (e.g., `c_project`, `rust_project`, `go_project`) may include slightly more complex behavior (such as calling utility functions), but are still kept intentionally small and readable.

---

## License

This project is provided for testing and demonstration purposes.

Unless otherwise specified in individual files, the contents of this repository are made available under the MIT License.

You are free to:

- Use this code for personal or commercial projects
- Modify and redistribute it

Subject to the terms of the MIT License. If you add or copy code from this repository into other projects, consider including a reference or copy of the license.

*(Replace this section with your actual license text or file reference, e.g., `See LICENSE for details`, if a specific license file is added.)*

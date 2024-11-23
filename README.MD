# Go WebAssembly Boilerplate

This project demonstrates how to run a Go program in the browser using WebAssembly. It outputs "Hello, World!" to the webpage and browser console.

## Setup

1. **Clone the Repo**:
   ```
   git clone https://github.com/yourusername/go-wasm-hello-world.git
   cd go-wasm-hello-world
   ```

2. **Get `wasm_exec.js`**:  
   Copy it from your Go installation:
   - **Windows**:
     ```
     copy "%GOROOT%\misc\wasm\wasm_exec.js" .
     ```
   - **Mac/Linux**:
     ```
     cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" .
     ```

3. **Build WebAssembly**:
   ```
   GOOS=js GOARCH=wasm go build -o hello.wasm hello.go
   ```

4. **Run the Web Server**:
   ```
   go run server.go
   ```

5. **Open in Browser**:  
   Visit `http://localhost:8080`.

## Files

- `hello.go`: Go source code.
- `hello.wasm`: Compiled WebAssembly.
- `wasm_exec.js`: JavaScript runtime for Go WebAssembly.
- `index.html`: Loads WebAssembly.


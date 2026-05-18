# AGENTS.md – Development & Review Guidelines (Fillit / Tetromino Solver Project)

This document defines how development, testing, and review should be performed for this tetromino square solver project.

---

# 1. Coding Rules

## 1.1 Standard Library Only

Only the Go standard library is allowed.  
No external packages should be used.

---

## 1.2 Code Style & Structure

- All code must be formatted using:

```bash
gofmt -w .
```

- Functions must be small (ideally <30 lines) and focused on a single responsibility.
- Avoid deeply nested logic.
- Prefer readable algorithms over overly optimized code.
- Separate parsing, validation, solving, and rendering logic.

---

## 1.3 Naming Conventions

### Variables

Use camelCase:

```go
filePath
gridSize
tetrominoes
```

### Functions

Use verb-based names:

```go
readFile
parseTetromino
canPlace
solveBoard
```

### Types & Constants

Use PascalCase:

```go
Tetromino
Board
GridSize
```

---

# 2. Project Goal (IMPORTANT)

The program must:

- Receive exactly one CLI argument:

```bash
go run . sample.txt
```

- Read tetrominoes from a text file
- Validate file format and tetromino structure
- Assemble all tetrominoes into the smallest possible square
- Print the solved board using uppercase latin letters:
  - First tetromino → `A`
  - Second → `B`
  - etc.
- Print `ERROR` for:
  - Invalid file format
  - Invalid tetromino structure
  - Invalid number of blocks
  - Impossible parsing cases

---

# 3. Program Rules

## 3.1 Input Handling

The input file:

- Contains tetrominoes separated by empty lines
- Each tetromino is:
  - 4 lines tall
  - 4 characters wide
- Valid characters:
  - `#`
  - `.`

Example:

```txt
#...
#...
#...
#...

....
....
..##
..##
```

Rules:

- At least one tetromino must exist
- Every tetromino must contain exactly 4 `#`
- Tetromino blocks must be connected
- Empty lines must appear only between tetrominoes
- Invalid formatting must return:

```txt
ERROR
```

---

## 3.2 Output Rules

The output:

- Must print the smallest valid square
- Must preserve tetromino order
- Must use uppercase letters for identification
- Must match exact formatting

Example:

```txt
ABB.
ABB.
A...
A...
```

No extra spaces.  
No additional text.

---

# 4. Testing Rules

## 4.1 Test-Driven Development (TDD)

Always follow:

1. Red
2. Green
3. Refactor

Write tests before implementation whenever possible.

---

## 4.2 Unit Tests

Each core function must be independently testable.

Functions to test:

```go
readFile()
parseTetromino()
validateTetromino()
normalizeTetromino()
canPlace()
placeTetromino()
removeTetromino()
solveBoard()
```

---

## 4.3 Golden Tests

Golden tests must validate complete program output.

Example:

Input:

```txt
##..
##..
....
....
```

Expected:

```txt
AA
AA
```

---

# 5. Core Implementation Rules

## 5.1 Tetromino Validation

Each tetromino must:

- Contain exactly 4 blocks (`#`)
- Be fully connected

Connectivity must be verified using adjacency checks.

Valid connections are:

- Up
- Down
- Left
- Right

Diagonal connections do NOT count.

---

## 5.2 Tetromino Normalization

All tetrominoes should be normalized:

- Shifted to top-left origin
- Stored consistently for placement logic

This simplifies solving algorithms.

Example:

Before normalization:

```txt
....
..##
..##
....
```

After normalization:

```txt
##..
##..
....
....
```

---

## 5.3 Solving Algorithm

Recommended approach:

- Backtracking
- Recursive placement
- Incremental square growth

Workflow:

1. Start from minimum possible square size
2. Try placing tetrominoes recursively
3. Backtrack on failure
4. Increase square size if no solution exists

---

## 5.4 Minimum Square Size

Initial board size should be:

```text
ceil(sqrt(number_of_tetrominoes * 4))
```

The solver should increase size only when necessary.

---

# 6. Edge Cases (CRITICAL)

You MUST handle:

- Empty file
- Missing lines
- Invalid characters
- Wrong tetromino dimensions
- Less or more than 4 blocks
- Disconnected tetrominoes
- Invalid separators
- Multiple consecutive empty lines
- Non-solvable placements at current board size

Behavior for invalid input:

```txt
ERROR
```

ONLY.

No panic.  
No stack traces.

---

# 7. Project Structure

```txt
.
├── main.go
├── internal/
│   ├── parser.go
│   ├── validator.go
│   ├── solver.go
│   ├── board.go
│   ├── tetromino.go
├── tests/
│   ├── unit/
│   ├── testdata/
```

Suggested separation:

- Parsing → file reading & formatting
- Validation → tetromino correctness
- Solver → backtracking algorithm
- Board → placement/removal logic

---

# 8. Using AI in Development

## 8.1 Allowed

Use AI only when:

- You already attempted solving the problem
- You want algorithm explanations
- You need help debugging
- You want alternative implementations

---

## 8.2 NOT Allowed

- Do NOT copy complete solver solutions
- Do NOT skip understanding recursion/backtracking
- Do NOT rely on AI without testing
- Do NOT use AI-generated code blindly

---

## 8.3 AI Usage Validation

You are using AI correctly if:

- You understand the solving algorithm
- You can explain recursion and backtracking
- You can debug placement logic yourself
- You understand WHY a tetromino is valid or invalid

---

# 9. Code Review Rules

Before committing:

- Code formatted with `gofmt`
- No unused variables
- No duplicated logic
- All tests pass
- Edge cases covered
- Output format EXACT
- Invalid inputs correctly print `ERROR`

---

# 10. Tooling & CI

Before every commit run:

```bash
gofmt -w .
go vet ./...
go test ./...
```

If any command fails → DO NOT commit.
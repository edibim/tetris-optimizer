# Product Requirements Document

## Project Title
Tetromino Smallest-Square Solver

## Document Purpose
This PRD defines the product goals, functional requirements, technical constraints, implementation expectations, and acceptance criteria for the tetromino solver project. It also reflects the intended working style for the project: building the solution with strong engineering habits so the primary developer learns to design and implement similar programs independently.

## Product Summary
Build a Go command-line program that reads a text file containing one or more tetrominoes, validates the input strictly, and prints the smallest possible square arrangement that places every tetromino in order. Invalid input must produce exactly `ERROR`.

## Problem Statement
The program must solve two problems correctly:

1. Parse and validate a constrained text format that represents tetromino pieces.
2. Compute the smallest square board that can contain all valid tetrominoes while preserving input order in the output lettering.

The solution must be correct, deterministic, readable, and implemented using only the Go standard library.

## Goals
- Accept exactly one CLI argument: the input file path.
- Parse tetromino definitions from the file.
- Reject malformed files and invalid tetromino shapes.
- Solve the placement problem using the smallest square possible.
- Print a board using uppercase Latin letters in input order.
- Keep the code modular, testable, and understandable.
- Support step-by-step development with strong learning value.

## Non-Goals
- No graphical interface.
- No external dependencies.
- No advanced optimization beyond what is needed for a clean and correct backtracking solution.
- No reordering of tetrominoes to simplify solving.
- No alternative output modes or debug output in the final program.

## Target User
The primary user is a student developer implementing the project to learn:
- file parsing
- validation
- recursion and backtracking
- board manipulation
- Go project structure
- test-driven development habits

## CLI Contract
The program must be executed as:

```bash
go run . sample.txt
```

Behavior:
- If exactly one valid argument is provided and the file content is valid, print the solved board.
- If the argument count is invalid, the file cannot be parsed correctly, or any tetromino is invalid, print exactly:

```text
ERROR
```

## Input Format Requirements
The input file must contain one or more tetrominoes.

Each tetromino:
- is represented by exactly 4 lines
- each line is exactly 4 characters wide
- uses only `#` and `.`
- contains exactly 4 `#` cells
- must be fully connected by orthogonal adjacency only

Tetrominoes:
- are separated by exactly one empty line
- must not contain extra empty lines
- must not contain missing lines
- must not contain invalid characters

An empty file is invalid.

## Tetromino Validation Rules
Each tetromino must satisfy all of the following:
- exactly 4 occupied blocks
- orthogonally connected shape
- valid 4x4 representation

Connectivity rules:
- allowed adjacency: up, down, left, right
- diagonal contact does not count as connected

Recommended representation:
- store occupied cells as coordinates
- normalize coordinates so the shape is shifted to the top-left origin

## Output Requirements
The program output must:
- represent the smallest valid square arrangement
- preserve tetromino order
- use uppercase letters starting from `A`
- use `.` for empty cells
- contain no extra spaces
- contain no explanatory text

Example:

```text
ABB.
ABB.
A...
A...
```

## Solving Requirements
The recommended solving approach is recursive backtracking.

Expected strategy:
1. Compute the minimum candidate board size with `ceil(sqrt(pieceCount * 4))`.
2. Attempt to place tetrominoes in order on the board.
3. Backtrack when a placement blocks future placements.
4. Increase board size only when no solution exists for the current size.

The solution should prioritize correctness and readability over premature optimization.

## Functional Requirements

### FR-1 File Reading
- Read the input file from the provided path.
- Return an error when the file cannot be read.

### FR-2 Input Parsing
- Split the file into tetromino blocks.
- Enforce strict file formatting rules.
- Reject malformed separators and incomplete tetromino definitions.

### FR-3 Tetromino Validation
- Validate dimensions, allowed characters, block count, and connectivity.
- Normalize valid tetrominoes into a consistent internal form.

### FR-4 Board Representation
- Represent the board in a way that supports:
  - checking whether a piece can be placed
  - placing a piece
  - removing a piece
  - rendering the final state

### FR-5 Solver
- Solve using backtracking.
- Preserve tetromino order.
- Return the smallest square solution.

### FR-6 Rendering
- Convert the solved board into exact output text.
- Print only the final board or `ERROR`.

### FR-7 Error Handling
- Invalid input must print exactly `ERROR`.
- The program must not panic or print stack traces.

## Technical Constraints
- Language: Go
- Dependencies: standard library only
- Formatting: `gofmt -w .`
- Quality gates before commit:
  - `gofmt -w .`
  - `go vet ./...`
  - `go test ./...`

## Recommended Project Structure
Suggested structure:

```text
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

Implementation note:
- The exact layout may evolve, but separation of concerns should remain clear.

## Recommended Internal Design

### Core Types
Recommended types include:
- `Tetromino`
- `Point`
- `Board`

### Core Functions
The following functions should be independently testable:
- `readFile()`
- `parseTetromino()`
- `validateTetromino()`
- `normalizeTetromino()`
- `canPlace()`
- `placeTetromino()`
- `removeTetromino()`
- `solveBoard()`

### Design Principles
- Keep functions small and focused.
- Prefer explicit logic over clever shortcuts.
- Avoid deep nesting when possible.
- Make invalid states hard to represent.
- Keep solver logic separate from parsing and printing.

## Testing Strategy
The project should follow a Red-Green-Refactor cycle whenever practical.

### Unit Tests
Create unit tests for:
- file parsing edge cases
- tetromino validation
- normalization
- board placement rules
- recursive solver behavior

### Golden Tests
Create full-program tests that verify exact output for:
- simple valid inputs
- multiple-piece valid inputs
- invalid formatting
- invalid shapes
- known examples from the project information

### Critical Edge Cases
Tests must cover:
- empty file
- wrong number of CLI arguments
- missing tetromino lines
- extra separator lines
- invalid characters
- fewer than 4 blocks
- more than 4 blocks
- disconnected blocks
- impossible placement on the current board size

## Acceptance Criteria
The project is considered complete when:
- the program compiles successfully
- the program accepts exactly one CLI argument
- valid inputs produce the smallest valid square
- invalid inputs print exactly `ERROR`
- tetromino order is preserved in lettering
- only standard library packages are used
- code is formatted and passes vet/test checks
- the developer can explain:
  - how parsing works
  - why a tetromino is valid or invalid
  - how normalization helps
  - how recursive backtracking finds the solution

## Delivery Milestones

### Milestone 1: Project Setup
- initialize the Go module
- create basic project structure
- add first tests and sample test data

### Milestone 2: Parsing and Validation
- implement file reading
- parse tetromino blocks
- validate formatting and connectivity
- normalize shapes

### Milestone 3: Board Operations
- create board representation
- implement `canPlace`, `place`, and `remove`
- test placement behavior thoroughly

### Milestone 4: Solver
- implement recursive backtracking
- start from the minimum candidate square
- grow board size only when necessary

### Milestone 5: Final Integration
- connect CLI, parser, validator, solver, and renderer
- add golden tests
- run formatting and verification commands

## Risks and Mitigations

### Risk: Parsing bugs
Mitigation:
- write strict parser tests before broader integration

### Risk: Incorrect tetromino connectivity checks
Mitigation:
- test several valid and invalid shapes explicitly

### Risk: Hard-to-debug recursion failures
Mitigation:
- keep solver state simple
- test board operations independently before full solver integration

### Risk: Overly complex implementation
Mitigation:
- prefer clear data structures
- keep responsibilities separated
- refactor aggressively when functions grow too much

## Mentorship Workflow
This project should be developed in a learning-oriented way:
- define one small target at a time
- prefer implementing the next smallest correct step
- review and explain each decision
- avoid copying full solutions without understanding
- use tests to drive confidence
- pause regularly to verify understanding before moving to the next layer

## Definition of Done
The project is done when the program meets all functional requirements, passes the required checks, and the primary developer can confidently explain and reproduce the solution approach without relying on opaque generated code.

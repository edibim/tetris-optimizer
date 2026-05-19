# Tetromino Smallest-Square Solver

This project is a Go command-line program that reads tetrominoes from a text file, validates the input strictly, and assembles all pieces into the smallest possible square using recursive backtracking.

The current CLI is developer-friendly: it prints `ERROR:` followed by a short explanation so invalid files are easier to debug while you build and review the project.

## Features

- Accepts exactly one CLI argument: the path to the input file
- Validates file structure and tetromino shape correctness
- Normalizes tetromino coordinates before solving
- Solves the smallest square arrangement with backtracking
- Preserves tetromino order and labels pieces with `A`, `B`, `C`, ...
- Uses only the Go standard library

## Run

```bash
go run . sample.txt
```

If the input is valid, the program prints the solved board.

If the input is invalid, the program prints a descriptive error, for example:

```text
ERROR: invalid character in tetromino file.
Details: tetromino A line 1 col 1 contains "`"
```

## Example

Input:

```text
#...
#...
#...
#...

....
....
..##
..##
```

Possible output:

```text
ABB.
ABB.
A...
A...
```

## Project Structure

```text
.
├── main.go
├── internal/
│   ├── board.go
│   ├── errors.go
│   ├── parser.go
│   ├── solver.go
│   ├── tetromino.go
│   └── validator.go
├── tests/
│   ├── testdata/
│   └── unit/
├── AI/
│   ├── AGENTS.md
│   ├── PRD.md
│   └── summary.txt
├── go.mod
└── README.md
```

## Where To Find What

- [main.go](/home/student/tetris-optimizer/main.go): CLI entrypoint and program flow
- [internal/parser.go](/home/student/tetris-optimizer/internal/parser.go): file parsing and tetromino block extraction
- [internal/validator.go](/home/student/tetris-optimizer/internal/validator.go): block count validation, connectivity validation, and normalization
- [internal/tetromino.go](/home/student/tetris-optimizer/internal/tetromino.go): core tetromino and point types
- [internal/board.go](/home/student/tetris-optimizer/internal/board.go): board allocation, placement checks, placement, removal, and rendering
- [internal/solver.go](/home/student/tetris-optimizer/internal/solver.go): smallest-square backtracking solver
- [tests/unit](/home/student/tetris-optimizer/tests/unit): unit tests for parsing, validation, board logic, and solving
- [tests/testdata](/home/student/tetris-optimizer/tests/testdata): sample inputs used by tests
- [AI/PRD.md](/home/student/tetris-optimizer/AI/PRD.md): senior-level project requirements and architecture notes

## Execution Flow

1. `main.go` validates the CLI argument count.
2. `ReadFile()` loads the raw file contents.
3. `ParseFile()` validates the text format and extracts tetromino coordinates.
4. `ValidateTetromino()` checks that each tetromino has exactly 4 blocks and is orthogonally connected.
5. `NormalizeTetromino()` shifts each shape to the top-left origin.
6. `SolveBoard()` starts from the minimum possible square and tries placements recursively.
7. The solved board is rendered and printed.

## Internal Error Design

The CLI now surfaces short, human-readable error messages so bad inputs are easier to inspect.

Internally, the code still uses more detailed sentinel errors so the implementation stays easier to debug and review:

- `ErrInvalidArgCount`
- `ErrReadFile`
- `ErrEmptyInput`
- `ErrUnexpectedBlankLine`
- `ErrTrailingSeparator`
- `ErrIncompleteBlock`
- `ErrInvalidRowWidth`
- `ErrInvalidCharacter`
- `ErrInvalidBlockCount`
- `ErrDisconnectedShape`
- `ErrBoardTooSmall`
- `ErrCellOccupied`

This keeps the internal logic maintainable while making CLI failures easier to understand.

## Testing

Run the full verification suite with:

```bash
GOCACHE=/tmp/go-build-cache go test ./...
GOCACHE=/tmp/go-build-cache go vet ./...
GOCACHE=/tmp/go-build-cache go build ./...
```

`/tmp/go-build-cache` is used here because some sandboxed environments expose a read-only default Go build cache directory.

## Audit Readiness

The repository is audit-ready because it includes:

- strict input validation
- clear separation of concerns
- deterministic normalization
- recursive backtracking solver
- unit tests for parser, validator, board, and solver behavior
- end-to-end tests for valid and invalid files
- a subject-style golden test using the official sample layout
- English-only repository documentation and comments

## Audit Talking Points

If you need to explain the project in an audit, focus on these points:

1. The parser is responsible only for text format and block extraction.
2. The validator is responsible for shape correctness and normalization.
3. The board layer is responsible for safe placement rules and final rendering.
4. The solver uses recursive backtracking and grows the square only when necessary.
5. The final CLI maps internal errors into clearer user-facing messages, which helps debug invalid tetromino files faster.

## Useful Test Files

- [tests/testdata/valid_single_o.txt](/home/student/tetris-optimizer/tests/testdata/valid_single_o.txt): simplest valid tetromino
- [tests/testdata/invalid_disconnected.txt](/home/student/tetris-optimizer/tests/testdata/invalid_disconnected.txt): invalid disconnected tetromino
- [tests/testdata/sample_subject.txt](/home/student/tetris-optimizer/tests/testdata/sample_subject.txt): larger subject-style sample used for golden testing

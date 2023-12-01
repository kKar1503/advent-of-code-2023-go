# Advent of Code 2023 (Go)

This is my attempt for Advent of Code 2023 written in Go.

## Running the solutions

Each day's attempt can be run with the following command:

```bash
go run cmd/day-{day_num}/main.go {part}
```

Example for day 1 part 1:

```bash
go run cmd/day-1/main.go part1
```

## Adding more solutions

For each solutions can be added by going to each day's directory,
for eg. `cmd/day-1/solutions/` and add a new file.

Within the file, export a function for the solution.

In the `main.go` file, append a new entry to the map for the solution,
and running the file with the argument can run the respective solution.

# 1, 2, 3... Go!

Stopwatch CLI using Go.

## WORK IN PROGRESS -- How to Use -- WORK IN PROGRESS

First of all, install the package:

```bash
go install github.com/ed-henrique/123-go
```

### Stopwatch Mode

Execute the command `123-go` and press `Enter` when done. The elapsed time will be shown.

### Timer Mode

Execute the command `123-go` with the `--timer` flag, specifying the amount of milliseconds for the time to count down from.

<details>
<summary>Example</summary>

```bash
123-go --timer 10000
```

</details>

```bash
go run main.go
```

## Development Environment Setup

1. Clone the repo

```bash
git clone git@github.com/ed-henrique/123-go.git
```

2. Run the program

```bash
go run main.go
```

## Roadmap

- [x] Time from milliseconds to days
- [x] Act as both a timer and a stopwatch

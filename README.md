# 1, 2, 3... Go!

Stopwatch CLI using Go.

## WORK IN PROGRESS -- How to Use -- WORK IN PROGRESS

1. Install the [latest release](https://github.com/ed-henrique/123-go/releases/latest) for your OS

### Stopwatch Mode

Run the binary `123-go` and press `Enter` when done. The elapsed time will be shown.

### Timer Mode

Run the binary `123-go` with the `--timer` flag, specifying the amount of milliseconds for the time to count down from.

<details>
<summary>Example</summary>

```bash
123-go --timer 10000
```

</details>

## Development Environment Setup

1. Clone the repo

```bash
git clone git@github.com/ed-henrique/123-go.git
```

2. Install the dependencies

```bash
go get ./...
```

3. Run the program

```bash
go run cmd/123-go/main.go
```

## Roadmap

- [ ] Build beautiful TUI
- [x] Time from milliseconds to days
- [x] Act as both a timer and a stopwatch

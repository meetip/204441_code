# Examples

This folder contains class examples organized into separate folders. Each example is a standalone Go program with `package main` and a `main` function.

Run an example with:

```bash
go run ./examples/<example-folder>
```

Examples included in this sample patch (renumbered by topic):
- `01_sequential` - sequential function calls
- `02_goroutines` - goroutine basics
- `03_longrun_task` - long-running background task (no synchronization)
- `03_waitgroup` - waitgroup example
- `04_pingpong_basic` - tiny ping/pong without synchronization
- `05_pingpong_channel` - ping/pong with a goroutine and channel
- `06_pingpong_chat` - ping/pong chat using two goroutines
- `06_race_condition` - race condition demo (unsynchronized counter; try `go run -race`)
 - `07_race_condition_fixed` - fixed version using a `sync.Mutex` (final count == 1000)



Guidelines:
- Put each runnable example in its own folder to avoid duplicate `main` functions.
- For shared utilities, create a `pkg/` package and import it.

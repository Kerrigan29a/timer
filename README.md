# Timer

[![GoDoc](https://godoc.org/github.com/Kerrigan29a/timer?status.svg)](https://godoc.org/github.com/Kerrigan29a/timer)

Simple timer with acoustic alarm

# Build
```bash
make
```

# Usage
The basic usage is to pass a duration compatible with [ParseDuration](https://golang.org/pkg/time/#ParseDuration)

> A duration string is a possibly signed sequence of decimal numbers, each with optional fraction and a unit suffix, such as "300ms", "-1.5h" or "2h45m". Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h".

```bash
.\timer 25m
```

It's possible to pass mutiple durations.
```bash
.\timer 25m 5m
```

If you want to repeat this secuence of durations just use the repeat option.
```bash
.\timer 25m 5m -r
```

It's also possible to mute the acoustic signal.
```bash
.\timer 25m -m
```

# Documentation

Documentation is available at [godoc](https://godoc.org/github.com/Kerrigan29a/timer)

# Contributions
Thanks to [juskiddink](https://freesound.org/people/juskiddink/) for his [gong.wav](https://freesound.org/people/juskiddink/sounds/86773/) sound

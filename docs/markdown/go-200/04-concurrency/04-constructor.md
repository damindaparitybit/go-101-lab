<!-- .slide: class="with-code-bg-dark" -->

# Concurrence

## Son constructeur associé

```go
// NewStatistics creates a new statistics structure and launches its worker routine
func NewStatistics(loggingPeriod time.Duration) *Statistics {
    sw := Statistics{
        statistics: make(chan bool, statisticsChannelSize),
        counter: 0,
        start: time.Now(),
        loggingPeriod: loggingPeriod,
    }
    go sw.run()
    return &sw
}
```

Notes:
SFR

constructeur avec pointeur

lance la go routine

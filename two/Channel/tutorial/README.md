# outlne
1. simple goroutine use `var wg sync.WaitGroup` and `wg.Wait()` to make sure every concurrency wouuld receive result
2. `<-chan` can be a blocker for processing
3. pre-defined a `buffered chan` help out concurrency without wait for another chan
4. use func ping & func pong to demostrate the effective of `func( chan )`
5. `close` chan when you don't need it
6. with `select` execute `case by case`
7. chan can also be a generator
8. parallel compute with `mutex`

# refer
https://michaelchen.tech/golang-prog/concurrency/

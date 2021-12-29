case :
Someone ask me about locking process in golang (especially in goroutine)

suggestion :
Try create example that com

result :
There are 3 functions called deductWalletX which allows you to run with concurrent process.
- deductWalletA is thread-safe
- deductWalletB isn't thread-safe,
- deductWalletC isn't thread-safe,

play :
you can run this multiple times then see the different result
`concurrency-race-condition-demo $ go run main.go `

or

using Go Race Detector
`concurrency-race-condition-demo $ go run -race main.go`
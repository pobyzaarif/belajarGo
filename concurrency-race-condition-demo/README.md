#### case :
Someone ask me about locking process in golang (especially in goroutine)

#### suggestion :
Try to create example

#### result :
There are 3 functions called deductWalletX which allows you to run with concurrent process.
- deductWalletA is thread-safe
- deductWalletB isn't thread-safe,
- `go deductWalletC` + `fmt.Println` potentially race condition,

#### play :
you can run this multiple times then see the different result
`concurrency-race-condition-demo $ go run main.go`

or

using Go Race Detector
`concurrency-race-condition-demo $ go run -race main.go`

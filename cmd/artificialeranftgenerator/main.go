// cmd/artificialeranftgenerator/main.go
package main

import (
"flag"
"log"
"os"

"artificialeranftgenerator/internal/artificialeranftgenerator"
)

func main() {
verbose := flag.Bool("verbose", false, "Enable verbose logging")
flag.Parse()

app := artificialeranftgenerator.NewApp(*verbose)
if err := app.Run(); err != nil {
log.Fatal(err)
}
}

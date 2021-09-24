# go-thaibaht
go-thaibaht is a simple library for convert number to thai baht as text

> Hit the project with a star if you find it useful ⭐

## Installation
```bash
go get -u github.com/artykaikub/go-thaibaht
```

## Example
```go
package main

import (
	"log"
	"github.com/artykaikub/go-thaibaht"
)

func main() {
	text := gothaibaht.ToText(123456789.99)
	fmt.Print(text)
	// output: หนึ่งร้อยยี่สิบสามล้านสี่แสนห้าหมื่นหกพันเจ็ดร้อยแปดสิบเก้าบาทเก้าสิบเก้าสตางค์
}
```




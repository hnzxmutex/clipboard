package clipboard_test

import (
	"fmt"

	"github.com/hnzxmutex/clipboard"
)

func Example() {
	clipboard.WriteAll("日本語")
	text, _ := clipboard.ReadAll()
	fmt.Println(text)

	// Output:
	// 日本語
}

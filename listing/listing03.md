Что выведет программа? Объяснить вывод программы. Объяснить внутреннее устройство интерфейсов и их отличие от пустых интерфейсов.

```go
package main

import (
	"fmt"
	"os"
)

func Foo() error {
	var err *os.PathError = nil
	return err
}

func main() {
	err := Foo()
	fmt.Println(err)
	fmt.Println(err == nil)
}
```

Ответ:
Выведет nil  false.
Интерфейс в go состоит из из двух полей. Первое хранит информацию о типе, а второе о конкретном значении, реализующее интерфейс. Интерфейс == nil только в том случае, когда первое поле равно nil. В данном примере первое поле содержит *os.PathError, а его значение указано nil, поэтому, проверка не проходит

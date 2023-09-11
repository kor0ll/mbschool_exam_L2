Что выведет программа? Объяснить вывод программы.

```go
package main

type customError struct {
	msg string
}

func (e *customError) Error() string {
	return e.msg
}

func test() *customError {
	{
		// do something
	}
	return nil
}

func main() {
	var err error
	err = test()
	if err != nil {
		println("error")
		return
	}
	println("ok")
}
```

Ответ:
Выведет "error".
В данном листинге, как и в листинге 3, err всегда будет != nil, так как интерфейс содержит *customError, значением которого указано nil, а интерфейс == nil только в том случае, когда он не содержит ничего в поле о типе

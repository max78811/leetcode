package main

import (
	"fmt"
	"math/rand"
	"time"
)

// func init() {
//    rand.Seed(time.Now().UnixNano())
//}

// Есть функция, работающая неопределённо долго и возвращающая число.
// Её тело нельзя изменять (представим, что внутри сетевой запрос).
func unpredictableFunc() int64 {
	rnd := rand.Int63n(5000)
	time.Sleep(time.Duration(rnd) * time.Millisecond)
	return rnd
}

// Нужно изменить функцию обёртку, которая будет работать с заданным таймаутом
// Если "длинная" функция отработала за это время - отлично, возвращаем результат
// Если нет - возвращаем ошибку. Результат работы в этом случае нам не важен.
//
// Дополнительно нужно измерить, сколько выполнялась эта функция (просто выведет)
// Сигнатуру функцию обёртки менять можно.
func predictableFunc(timeout time.Duration) (int64, error) {
	t := time.NewTimer(timeout)
	defer t.Stop()

	resultCh := make(chan int64, 1)

	go func() {
		resultCh <- unpredictableFunc()
		close(resultCh)
	}()

	select {
	case <-t.C: // ctx.Done()
		return 0, fmt.Errorf("timeout") // ctx.Err
	case result := <-resultCh:
		return result, nil
	}
}

func main() {
	fmt.Println("started")
	fmt.Println(predictableFunc(5 * time.Second))
}

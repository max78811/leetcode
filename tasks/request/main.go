package request

import (
	"context"
	"fmt"
	"log"
	"net/http"
)

//1. Поочередно выполнит http запросы по предложенному списку ссылок
//в случае получения http-кода ответа на запрос "200 OK" печатаем на экране "адрес url - ok"
//в случае получения http-кода ответа на запрос отличного от "200 OK" либо в случае ошибки печатаем на экране "адрес url - not ok"
//
//2. Модифицируйте программу таким образом, чтобы использовались каналы для коммуникации основного потока с горутинами.
//Пример:
//Запросы по списку выполняются в горутинах.
//Печать результатов на экран происходит в основном потоке
//
//3. Модифицируйте программу таким образом, чтобы нигде не использовалась длина слайса урлов. Считайте, что урлы приходят из внешнего источника.
//Сколько их будет заранее - неизвестно. Предложите идиоматичный вариант, как ваша программа будет узнавать об окончании списка и передавать сигнал
//об окончании действий далее.
//
//4. Модифицируйте программу таким образом, что бы при получении 2 первых ответов с "200 OK" остальные запросы штатно прерывались.
//При этом необходимо напечатать на экране сообщение о завершении запроса.
//
//
//package main

func httpRequest(ctx context.Context, url string) (http.StatusCode, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatalf("%v", err)
	}

	req = req.WithContext(ctx)

	client := http.DefaultClient
	res, err := client.Do(req)
	if err != nil {
		log.Fatalf("%v", err)
	}

	return res.Status, nil
}

type Responce struct {
	url        string
	statusCode http.StatusCode
	err        error
}

func UrlsWriter() chan string {
	urlsCh := make(chan string)

	var urls = []string{
		"http://ozon.ru",
		"https://ozon.ru",
		"http://google.com",
		"http://somesite.com",
		"http://non-existent.domain.tld",
		"https://ya.ru",
		"http://ya.ru",
		"http://ёёёё",
	}

	go func() {
		for _, url := range urls {
			urlsCh <- url
		}
		close(urlsCh)
	}
	return urlsCh
}

const SuccessReqCount = 2

func main() {
	urlsCh := UrlsWriter()
	resultCh := make(chan Responce)
	wg := &sync.Waitgroup{}
	ctx, cancel := context.WithCancel(context.Background())

	for url := range urlsCh {
		wg.Add(1)
		go func() {
			code, err := httpRequest(ctx, url)

			urlResponce := Responce{
				url:        url,
				statusCode: code,
				err:        err,
			}

			resultCh <- urlResponce
			wg.Done()
		}
	}

	go func() {
		wg.Wait()
		close(resultCh)
	}

	okCount := 0
	for result := range resultCh {
		if result.err != nil || result.code != http.StatusOk {
			okCount++
			if okCount >= SuccessReqCount {
				cancel()
			}
			fmt.Println(fmt.Sprintf("адрес %s - not ok", url))
			continue
		}
		fmt.Println(fmt.Sprintf("адрес %s - ok", url))
	}
}

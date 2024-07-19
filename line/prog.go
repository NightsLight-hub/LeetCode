/*
Package main
@Time : 2024/7/10 19:42
@Author : sunxy
@File : prog
@description:
*/
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
)

type ItemCounts struct {
	ItemId int `json:"itemId"`
	Count  int `json:"count"`
}

type ItemPrice struct {
	Id    int `json:"id"`
	Price int `json:"price"`
}

var lock sync.Mutex
var itemSalesMap = make(map[int]int)
var totalPrice atomic.Int64

func main() {
	startTime := time.Now()
	// 生成6月每天的日期字符串，yyyyMMdd格式
	dayStr := generateDays()

	var wg sync.WaitGroup
	errCh := make(chan error, len(dayStr))
	for _, day := range dayStr {
		wg.Add(1)
		go func(day string) {
			defer wg.Done()
			err := fetchItemsSalesCount(day)
			if err != nil {
				errCh <- err
			}
		}(day)
	}

	wg.Wait()
	close(errCh)
	for err := range errCh {
		fmt.Printf("error occurred: %v\n", err)
		return
	}

	errCh = make(chan error, len(dayStr))
	for itemId, count := range itemSalesMap {
		wg.Add(1)
		go func(itemId int, count int) {
			defer wg.Done()
			err := CalculatePrice(itemId, count)
			if err != nil {
				errCh <- err
			}
		}(itemId, count)
	}
	wg.Wait()
	close(errCh)
	for err := range errCh {
		fmt.Printf("error occurred: %v\n", err)
		return
	}

	fmt.Printf("totalPrice: %d\n", totalPrice.Load())
	fmt.Printf("cost time %d ms", time.Since(startTime).Milliseconds())
}

func generateDays() []string {
	dayStr := make([]string, 30)
	for i := 0; i <= 29; i++ {
		if i < 9 {
			dayStr[i] = "2023060" + strconv.Itoa(i+1)
		} else {
			dayStr[i] = "202306" + strconv.Itoa(i+1)
		}
	}
	return dayStr
}

func fetchItemsSalesCount(day string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:11021/line/codetest2/sales/"+day, nil)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to execute request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	var itemCounts []ItemCounts
	err = json.Unmarshal(body, &itemCounts)
	if err != nil {
		return fmt.Errorf("failed to unmarshal response body: %w", err)
	}

	// 也可以用sync.map，但是优化不明显
	lock.Lock()
	defer lock.Unlock()
	for _, itemCount := range itemCounts {
		itemSalesMap[itemCount.ItemId] += itemCount.Count
	}
	return nil
}

func CalculatePrice(itemId int, count int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:11021/line/codetest2/item/"+strconv.Itoa(itemId), nil)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to execute request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	var itemPrice ItemPrice
	err = json.Unmarshal(body, &itemPrice)
	if err != nil {
		return fmt.Errorf("failed to unmarshal response body: %w", err)
	}

	totalPrice.Add(int64(itemPrice.Price * count))

	return nil
}

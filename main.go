package main

import (
	"errors"
	"fmt"
	"go_basic/cmd/miniorder"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	passed := 0
	failed := 0

	check := func(name string, ok bool) {
		if ok {
			fmt.Printf("  ✅ %s\n", name)
			passed++
		} else {
			fmt.Printf("  ❌ %s\n", name)
			failed++
		}
	}

	// === 1. 정상 주문 → 재고 차감, 결제 성공, 영수증 출력 ===
	fmt.Println("=== 1. 정상 주문 ===")
	inv := miniorder.NewInventory()
	inv.RegisterProduct("P001")
	inv.RegisterProduct("P002")
	inv.AddStock("P001", 10)
	inv.AddStock("P002", 5)

	order := &miniorder.Order{
		ID: "ORD-001",
		Items: []miniorder.OrderItem{
			{ProductID: "P001", Quantity: 2, Subtotal: 4000000},
			{ProductID: "P002", Quantity: 1, Subtotal: 500000},
		},
		Status:    "pending",
		CreatedAt: time.Now(),
	}

	card := &miniorder.CreditCard{
		CardNumber: "1234-5678",
		Balance:    10000000,
	}
	receipt := &miniorder.ReceiptWriter{}

	err := miniorder.ProcessOrder(inv, order, card, receipt)
	check("ProcessOrder 성공", err == nil)
	check("재고 차감 P001 (10→8)", inv.HasStock("P001", 8) && !inv.HasStock("P001", 9))
	check("재고 차감 P002 (5→4)", inv.HasStock("P002", 4) && !inv.HasStock("P002", 5))
	check("결제 차감 (10,000,000→5,500,000)", card.Balance == 5500000)
	check("영수증 출력됨", receipt.String() != "")
	check("주문 상태 completed", order.Status == "completed")

	// === 2. 재고 부족 → ErrOutOfStock, 재고 변동 없음 ===
	fmt.Println("\n=== 2. 재고 부족 ===")
	order2 := &miniorder.Order{
		ID: "ORD-002",
		Items: []miniorder.OrderItem{
			{ProductID: "P001", Quantity: 100, Subtotal: 200000000},
		},
		Status:    "pending",
		CreatedAt: time.Now(),
	}
	receipt2 := &miniorder.ReceiptWriter{}

	err2 := miniorder.ProcessOrder(inv, order2, card, receipt2)
	check("ProcessOrder 실패", err2 != nil)
	check("ErrOutOfStock 반환", errors.Is(err2, miniorder.ErrOutOfStock))
	check("재고 변동 없음 P001 (여전히 8)", inv.HasStock("P001", 8) && !inv.HasStock("P001", 9))
	check("주문 상태 pending 유지", order2.Status == "pending")

	// === 3. 결제 실패 → 재고 복원, PaymentError wrapping ===
	fmt.Println("\n=== 3. 결제 실패 ===")
	poorCard := &miniorder.CreditCard{
		CardNumber: "0000-0000",
		Balance:    1,
	}
	order3 := &miniorder.Order{
		ID: "ORD-003",
		Items: []miniorder.OrderItem{
			{ProductID: "P001", Quantity: 1, Subtotal: 2000000},
		},
		Status:    "pending",
		CreatedAt: time.Now(),
	}
	receipt3 := &miniorder.ReceiptWriter{}

	err3 := miniorder.ProcessOrder(inv, order3, poorCard, receipt3)
	check("ProcessOrder 실패", err3 != nil)
	var pe *miniorder.PaymentError
	check("PaymentError 타입 확인", errors.As(err3, &pe))
	check("에러 wrapping (ErrInsufficientBalance)", errors.Is(err3, miniorder.ErrInsufficientBalance))
	check("재고 복원 P001 (여전히 8)", inv.HasStock("P001", 8) && !inv.HasStock("P001", 9))
	check("주문 상태 pending 유지", order3.Status == "pending")

	// === 4. 동시성 검증 (10 고루틴) → race condition 없음 ===
	fmt.Println("\n=== 4. 동시성 검증 (10 고루틴) ===")
	inv2 := miniorder.NewInventory()
	inv2.RegisterProduct("ITEM-A")
	inv2.AddStock("ITEM-A", 100)

	bank := &miniorder.BankTransfer{
		AccountNumber: "110-1234",
		Balance:       10000000,
	}

	var (
		wg           sync.WaitGroup
		successCount int32
		failCount    int32
	)

	for i := range 10 {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			o := &miniorder.Order{
				ID: fmt.Sprintf("CONCURRENT-%d", n),
				Items: []miniorder.OrderItem{
					{ProductID: "ITEM-A", Quantity: 5, Subtotal: 50000},
				},
				Status:    "pending",
				CreatedAt: time.Now(),
			}
			r := &miniorder.ReceiptWriter{}
			if err := miniorder.ProcessOrder(inv2, o, bank, r); err != nil {
				atomic.AddInt32(&failCount, 1)
				fmt.Printf("  [고루틴 %02d] 실패: %v\n", n, err)
			} else {
				atomic.AddInt32(&successCount, 1)
				fmt.Printf("  [고루틴 %02d] 성공\n", n)
			}
		}(i)
	}
	wg.Wait()

	check("전체 10건 성공", atomic.LoadInt32(&successCount) == 10)
	check("실패 0건", atomic.LoadInt32(&failCount) == 0)
	check("재고 정합성 (100→50)", inv2.HasStock("ITEM-A", 50) && !inv2.HasStock("ITEM-A", 51))
	check("결제 정합성 (10,000,000→9,500,000)", bank.Balance == 9500000)

	// === 최종 결과 ===
	fmt.Printf("\n=============================\n")
	fmt.Printf("검증 완료: ✅ %d 통과 / ❌ %d 실패\n", passed, failed)
	fmt.Printf("=============================\n")
}

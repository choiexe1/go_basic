# 🛒 미니 프로젝트 — 주문 처리 시스템

## Step 1 — 에러 정의 (`error.go`)

- [x] `ErrOutOfStock` — 재고 부족 (sentinel)
- [x] `ErrInvalidQuantity` — 수량 ≤ 0 (sentinel)
- [x] `ErrProductNotFound` — 상품 없음 (sentinel)
- [x] `PaymentError` — 결제 실패 (커스텀 에러 타입)
  - `Method string` (결제 수단명)
  - `Reason error` (원인 — wrapping용)
  - `Error() string` 구현
  - `Unwrap() error` 구현

## Step 2 — 상품 (`product.go`)

### `Category` (iota)

- [x] `Electronics`, `Food`, `Clothing`
- [x] `Stringer` 구현

### `Product`

| 필드       | 타입       | JSON 태그         |
| ---------- | ---------- | ----------------- |
| `ID`       | `string`   | `json:"id"`       |
| `Name`     | `string`   | `json:"name"`     |
| `Price`    | `int`      | `json:"price"`    |
| `Category` | `Category` | `json:"category"` |

- [x] `Stringer` 구현 — `[Electronics] 맥북 (2,000,000원)`

## Step 3 — 주문 (`order.go`)

### `OrderItem`

| 필드        | 타입     | JSON 태그           |
| ----------- | -------- | ------------------- |
| `ProductID` | `string` | `json:"product_id"` |
| `Quantity`  | `int`    | `json:"quantity"`   |
| `Subtotal`  | `int`    | `json:"subtotal"`   |

### `NewOrderItem(product *Product, quantity int) OrderItem`

- `Subtotal = product.Price × quantity` 자동 계산

### `Order`

| 필드        | 타입          | JSON 태그           |
| ----------- | ------------- | ------------------- |
| `ID`        | `string`      | `json:"id"`         |
| `Items`     | `[]OrderItem` | `json:"items"`      |
| `Status`    | `string`      | `json:"status"`     |
| `CreatedAt` | `time.Time`   | `json:"created_at"` |

- [x] `Total() int` — Items의 Subtotal 합산
- [x] `Stringer` 구현 — `주문 ORD-001 | 3건 | 총 5,000,000원 | pending`

## Step 4 — 재고 (`inventory.go`)

### `Inventory` (필드 unexported)

| 필드    | 타입                             |
| ------- | -------------------------------- |
| `stock` | `map[string]int` (상품ID → 수량) |
| `mu`    | `sync.Mutex`                     |

- [x] `NewInventory() *Inventory`
- [x] `AddStock(productID string, qty int) error`
- [x] `RemoveStock(productID string, qty int) error`
- [x] `HasStock(productID string, qty int) bool`

## Step 5 — 결제 (`payment.go`)

### `PaymentMethod` 인터페이스

- [x] `Pay(amount int) error`

### `CreditCard`

| 필드      | 타입         |
| --------- | ------------ |
| `Number`  | `string`     |
| `Balance` | `int`        |
| `mu`      | `sync.Mutex` |

- [x] `Pay(amount int) error` — Balance 부족 시 `PaymentError` 반환, 성공 시 Balance 차감

### `BankTransfer`

| 필드            | 타입         |
| --------------- | ------------ |
| `AccountNumber` | `string`     |
| `Balance`       | `int`        |
| `mu`            | `sync.Mutex` |

- [x] `Pay(amount int) error` — Balance 부족 시 `PaymentError` 반환, 성공 시 Balance 차감

## Step 6 — 영수증 (`receipt.go`)

### `ReceiptWriter`

| 필드     | 타입     |
| -------- | -------- |
| `buffer` | `[]byte` |

- [ ] `Write(p []byte) (int, error)` — io.Writer 구현
- [ ] `String() string` — 버퍼 내용 문자열 반환

## Step 7 — 주문 처리 (`process.go`)

### `ProcessOrder(inventory *Inventory, order *Order, payment PaymentMethod, receipt *ReceiptWriter) error`

1. 각 아이템별 재고 확인
2. 각 아이템별 재고 차감 (중간 실패 시 이미 차감된 것 롤백)
3. 결제 처리 (실패 시 전체 재고 복원 + `PaymentError` wrapping)
4. 성공 시 `receipt`에 영수증 기록, `order.Status = "completed"`

## Step 8 — 동시성 검증 (`main.go`)

- [ ] 10개 고루틴 동시 주문 처리
- [ ] Inventory, Payment 모두 Mutex로 보호
- [ ] `go run -race` 검증

## 검증 항목

- [ ] 정상 주문 → 재고 차감, 결제 성공, 영수증 출력
- [ ] 재고 부족 → ErrOutOfStock, 재고 변동 없음
- [ ] 결제 실패 → 재고 복원, PaymentError wrapping
- [ ] 동시 주문 10건 → race condition 없음

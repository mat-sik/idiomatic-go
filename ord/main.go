package main

import (
	"fmt"
	"unsafe"
)

func main() {
    orderInfo := OrderInfo{}
    fmt.Printf("OrderInfo (Total Size: %d)\n", unsafe.Sizeof(orderInfo))
    fmt.Printf("  %-15s Offset: %-4d Size: %d\n", "OrderCode", 
        unsafe.Offsetof(orderInfo.OrderCode), unsafe.Sizeof(orderInfo.OrderCode))
    fmt.Printf("  %-15s Offset: %-4d Size: %d\n", "Amount", 
        unsafe.Offsetof(orderInfo.Amount), unsafe.Sizeof(orderInfo.Amount))
    fmt.Printf("  %-15s Offset: %-4d Size: %d\n", "OrderNumber", 
        unsafe.Offsetof(orderInfo.OrderNumber), unsafe.Sizeof(orderInfo.OrderNumber))
    fmt.Printf("  %-15s Offset: %-4d Size: %d\n", "Items", 
        unsafe.Offsetof(orderInfo.Items), unsafe.Sizeof(orderInfo.Items))
    fmt.Printf("  %-15s Offset: %-4d Size: %d\n", "IsReady", 
        unsafe.Offsetof(orderInfo.IsReady), unsafe.Sizeof(orderInfo.IsReady))

    smallOrderInfo := SmallOrderInfo{}
    fmt.Printf("SmallOrderInfo (Total Size: %d)\n", unsafe.Sizeof(smallOrderInfo))
    fmt.Printf("  %-15s Offset: %-4d Size: %d\n", "IsReady", 
        unsafe.Offsetof(smallOrderInfo.IsReady), unsafe.Sizeof(smallOrderInfo.IsReady))
    fmt.Printf("  %-15s Offset: %-4d Size: %d\n", "OrderNumber", 
        unsafe.Offsetof(smallOrderInfo.OrderNumber), unsafe.Sizeof(smallOrderInfo.OrderNumber))
    fmt.Printf("  %-15s Offset: %-4d Size: %d\n", "OrderCode", 
        unsafe.Offsetof(smallOrderInfo.OrderCode), unsafe.Sizeof(smallOrderInfo.OrderCode))
    fmt.Printf("  %-15s Offset: %-4d Size: %d\n", "Amount", 
        unsafe.Offsetof(smallOrderInfo.Amount), unsafe.Sizeof(smallOrderInfo.Amount))
    fmt.Printf("  %-15s Offset: %-4d Size: %d\n", "Items", 
        unsafe.Offsetof(smallOrderInfo.Items), unsafe.Sizeof(smallOrderInfo.Items))
}

type OrderInfo struct {
	OrderCode   rune
	Amount      int
	OrderNumber uint16
	Items       []string
	IsReady     bool
}

type SmallOrderInfo struct {
	IsReady     bool
	OrderNumber uint16
	OrderCode   rune
	Amount      int
	Items       []string
}

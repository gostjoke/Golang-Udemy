// =======================================
// @File        : context.go
// @Author      : Tien-Wei Hsu
// @Date        : 2025-07-19 23:47
// @LastEditTime: 2025-07-19 23:47
// @Description : Description
// =======================================

package contextex

import (
	"context"
	"fmt"
)

func Contextuseway() {

	ctx := context.Background()
	fmt.Println("Context Example started")
	fmt.Println("Context: ", ctx)
	fmt.Println("Context error: ", ctx.Err())
	// print type
	fmt.Printf("Context type: %T\n", ctx)
	// cancel
	ctx, cancel := context.WithCancel(ctx)

	fmt.Println("Context with cancel created")
	fmt.Println("Context with cancel: ", ctx)
	fmt.Println("Context with cancel error: ", ctx.Err())
	fmt.Printf("Context type: %T\n", ctx)
	// fmt.Println("Context done: ", ctx.Done())
	fmt.Printf("Context type: %T", cancel)

	cancel() // cancel the context
	fmt.Println("cancel: ", ctx)
	fmt.Println("cancel error: ", ctx.Err())
	fmt.Printf("Context type after cancel: %T\n", ctx)

}

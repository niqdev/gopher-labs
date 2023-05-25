package myconcurrency

import (
	"context"
	"fmt"

	"golang.org/x/sync/errgroup"
)

// https://go101.org/article/channel-closing.html

func ColorFilterGroup() {

	colorChannel := make(chan MyColor, 1)
	//includeChannel := make(chan MyColor, 1)
	//excludeChannel := make(chan MyColor, 1)
	//countChannel := make(chan int, 1)

	group, ctx := errgroup.WithContext(context.TODO())

	// generate colors
	group.Go(func() error {
		for _, color := range allColor() {
			colorChannel <- color
		}
		close(colorChannel)
		return nil
	})

	// print + sleep
	group.Go(func() error {
		//defer close(colorChannel)
		for {
			select {
			case <-ctx.Done():
				return ctx.Err()
			case color, ok := <-colorChannel:
				if ok {
					fmt.Println(fmt.Sprintf(">>> COLOR: %v", color))
				} else {
					return nil
				}
			}
		}
	})

	if err := group.Wait(); err == nil {
		fmt.Println("finished")
	}
}

package test_errgroup

import (
	"fmt"
	"golang.org/x/sync/errgroup"
	"time"
)
type TestErrGroup struct {

}
func (s *TestErrGroup) Handle() error {
	var g errgroup.Group
	g.Go(func() error {
		time.Sleep(4 * time.Second)
		fmt.Println("协程1")
		return nil
	})

	g.Go(func() error {
		fmt.Println("协程2")
		return nil
	})

	if err := g.Wait(); err != nil {
		fmt.Println("errgroup error:", err)
		return err
	}

	time.Sleep(10 * time.Second)
	return nil
}


package gogroup

import (
	"fmt"
	"log"
	"runtime/debug"

	"golang.org/x/sync/errgroup"
)

type goGroup struct {
	*errgroup.Group
}

func (g *goGroup) Go(f func() error) {
	g.Group.Go(func() (err error) {
		defer func() {
			if e := recover(); e != nil {
				err = fmt.Errorf("panic happened: %v", e)

				if debugMode {
					log.Println(string(debug.Stack()))
				}
			}
		}()

		err = f()
		return err
	})
}

func NewGroup() *goGroup {
	errgroupGroup := errgroup.Group{}
	return &goGroup{
		Group: &errgroupGroup,
	}
}

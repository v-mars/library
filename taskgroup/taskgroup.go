package taskgroup

import (
	"context"
	"github.com/v-mars/library/logs"
	"sync/atomic"

	"golang.org/x/sync/errgroup"
)

type TaskGroup interface {
	Go(f func() error)
	Wait() error
}

type taskGroup struct {
	errGroup    *errgroup.Group
	ctx         context.Context
	execAllTask atomic.Bool
}

// NewTaskGroup 创建一个新的任务组实例
// 参数:
//
//	ctx: 上下文，用于控制任务组的生命周期
//	concurrentCount: 并发任务数量限制，控制同时执行的任务数
//
// 返回值:
//
//	TaskGroup: 新创建的任务组实例
func NewTaskGroup(ctx context.Context, concurrentCount int) TaskGroup {
	// 创建taskGroup实例
	t := &taskGroup{}
	// 使用errgroup创建带上下文的任务组
	t.errGroup, t.ctx = errgroup.WithContext(ctx)
	// 设置并发任务数量限制
	t.errGroup.SetLimit(concurrentCount)
	// 初始化execAllTask标志为false
	t.execAllTask.Store(false)

	return t
}

// NewUninterruptibleTaskGroup if one task return error, the rest task will continue
func NewUninterruptibleTaskGroup(ctx context.Context, concurrentCount int) TaskGroup {
	t := &taskGroup{}
	t.errGroup, t.ctx = errgroup.WithContext(ctx)
	t.errGroup.SetLimit(concurrentCount)
	t.execAllTask.Store(true)

	return t
}

func (t *taskGroup) Go(f func() error) {
	t.errGroup.Go(func() error {
		defer func() {
			if err := recover(); err != nil {
				logs.CtxErrorf(t.ctx, "[TaskGroup] exec panic recover:%+v", err)
			}
		}()

		if !t.execAllTask.Load() {
			select {
			case <-t.ctx.Done():
				return t.ctx.Err()
			default:
			}
		}

		return f()
	})
}

func (t *taskGroup) Wait() error {
	return t.errGroup.Wait()
}

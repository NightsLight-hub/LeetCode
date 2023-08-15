/*
@Time : 2021/7/7 7:18
@Author : sunxy
@File : trylock
@description:
*/
package trylock

type lock struct {
	ch chan struct{}
}

func newLock() *lock {
	l := new(lock)
	l.ch = make(chan struct{}, 1)
	l.ch <- struct{}{}
	return l
}

func (l *lock) tryLock() bool {
	r := false
	select {
	case <-l.ch:
		r = true
	default:

	}
	return r
}

func (l *lock) unlock() {
	l.ch <- struct{}{}
}

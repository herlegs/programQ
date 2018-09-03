package util

import (
	"sync"
	"sync/atomic"
)

func QSort(list []int) {
	qsort(list, 0, len(list) - 1)
}

func qsort(list []int, start, end int) {
	if start >= end || start < 0 || end >= len(list) {
		return
	}
	pivot := list[start+(end-start)/2]
	i, j := start, end
	for i <= j {
		for list[i] < pivot {
			i++
		}
		for list[j] > pivot {
			j--
		}
		if i <= j {
			list[i], list[j] = list[j], list[i]
			i++
			j--
		}
	}
	qsort(list, start, j)
	qsort(list, i, end)
}

func QSortM(list []int) {
	wg := &sync.WaitGroup{}
	qsortm(list, 0, len(list)-1, wg, false)
	wg.Wait()
}

const MaxThreadNum = 200

var semaphore = int32(MaxThreadNum)

func qsortm(list []int, start,end int, wg *sync.WaitGroup, isSubThread bool) {
	if isSubThread {
		defer wg.Done()
	}
	if start >= end || start < 0 || end >= len(list) {
		return
	}
	pivot := list[start+(end-start)/2]
	i, j := start, end
	for i <= j {
		for list[i] < pivot {
			i++
		}
		for list[j] > pivot {
			j--
		}
		if i <= j {
			list[i], list[j] = list[j], list[i]
			i++
			j--
		}
	}

	val := atomic.LoadInt32(&semaphore)
	if val > 0 && atomic.CompareAndSwapInt32(&semaphore, val, val-1){
		wg.Add(1)
		go func(){
			qsortm(list, start, j, wg, true)
			atomic.AddInt32(&semaphore,1)
		}()
	} else {
		qsortm(list, start, j, wg, false)
	}

	val = atomic.LoadInt32(&semaphore)
	if val > 0 && atomic.CompareAndSwapInt32(&semaphore, val, val-1){
		wg.Add(1)
		go func(){
			qsortm(list, i, end, wg, true)
			atomic.AddInt32(&semaphore,1)
		}()
	} else {
		qsortm(list, i, end, wg, false)
	}
}
package medium


/*
Design a hit counter which counts the number of hits received in the past 5 minutes.

Each function accepts a timestamp parameter (in seconds granularity) and you may assume that calls are being made to the system in chronological order (ie, the timestamp is monotonically increasing). You may assume that the earliest timestamp starts at 1.

It is possible that several hits arrive roughly at the same time.

Example:
HitCounter counter = new HitCounter();

// hit at timestamp 1.
counter.hit(1);

// hit at timestamp 2.
counter.hit(2);

// hit at timestamp 3.
counter.hit(3);

// get hits at timestamp 4, should return 3.
counter.getHits(4);

// hit at timestamp 300.
counter.hit(300);

// get hits at timestamp 300, should return 4.
counter.getHits(300);

// get hits at timestamp 301, should return 3.
counter.getHits(301);
 */
type HitCounter struct {
	tail *Node
}

type Node struct{
	Key int
	Count int
	Prev *Node
	Next *Node
}


/** Initialize your data structure here. */
func Constructor() HitCounter {
	return HitCounter{&Node{0,0,nil,nil}}
}


/** Record a hit.
        @param timestamp - The current timestamp (in seconds granularity). */
func (this *HitCounter) Hit(timestamp int)  {
	if this.tail.Key == 0 {
		this.tail.Key = timestamp
		this.tail.Count = 1
	}else if this.tail.Key == timestamp{
		this.tail.Count++
	}else{
		next := &Node{
			timestamp, 1, this.tail, nil,
		}
		this.tail.Next = next
		this.tail = next
	}
}


/** Return the number of hits in the past 5 minutes.
        @param timestamp - The current timestamp (in seconds granularity). */
func (this *HitCounter) GetHits(timestamp int) int {
	if this.tail == nil {
		return 0
	}
	start := this.tail
	count := 0
	for start != nil && (timestamp - start.Key) < 300{
		if start.Key <= timestamp{
			count += start.Count
		}
		start = start.Prev
	}
	return count
}


/**
 * Your HitCounter object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Hit(timestamp);
 * param_2 := obj.GetHits(timestamp);
 */



/**
 * Your HitCounter object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Hit(timestamp);
 * param_2 := obj.GetHits(timestamp);
 */

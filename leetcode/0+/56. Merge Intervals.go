//first
import (
	"container/heap"
	"sort"
)

func merge(intervals []Interval) []Interval {
	if len(intervals) < 2 {
		return intervals
	}

	pq := make(PriorityQueue, len(intervals))
	i := 0
	for _, value := range intervals {
		if value.Start > value.End {
			value.Start, value.End = value.End, value.Start
		}
		pq[i] = &unit{value.Start, value.End}
		i++
	}
	heap.Init(&pq)

	item := heap.Pop(&pq).(*unit)
	answer := make([]Interval, 0, 1)
	for pq.Len() > 0 {
		next := heap.Pop(&pq).(*unit)
		if item.value < next.priority {
			answer = append(answer, Interval{item.priority, item.value})
			item = next
		} else {
			if item.value < next.value {
				item.value = next.value
			}
		}
	}
	answer = append(answer, Interval{item.priority, item.value})
	return answer

}

type unit struct {
	priority int
	value    int
}

type PriorityQueue []*unit

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*unit))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

// second

func merge(intervals []Interval) []Interval {
	if len(intervals) < 2 {
		return intervals
	}
	sort.Slice(intervals, func(i, j int) bool { return intervals[i].Start < intervals[j].Start })

	answer := []Interval{}

	item := intervals[0]
	for i := 1; i < len(intervals); i++ {
		next := intervals[i]

		if item.End < next.Start {
			answer = append(answer, item)
			item = next
		} else {
			if item.End < next.End {
				item.End = next.End
			}
		}
	}

	answer = append(answer, item)
	return answer
}

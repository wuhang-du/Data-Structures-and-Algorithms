// first
import "container/heap"

type unit struct {
	value    int
	priority int
}

type PriorityQueue []*unit

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority > pq[j].priority
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

func topKFrequent(nums []int, k int) []int {
	if nums == nil || k < 1 {
		return nil
	}

	temp := make(map[int]int)

	for _, v := range nums {
		temp[v]++
	}

	pq := make(PriorityQueue, len(temp))
	i := 0
	for value, priority := range temp {
		pq[i] = &unit{value: value, priority: priority}
		i++
	}
	heap.Init(&pq)

	answer := make([]int, k)
	for i := 0; i < k; i++ {
		u := heap.Pop(&pq).(*unit)
		answer[i] = u.value
	}
	return answer
}

// second

func topKFrequent(nums []int, k int) []int {
	counts := make(map[int]int)
	for _, v := range nums {
		counts[v]++
	}

	pq := make([][]int, len(nums)+1)
	for value, count := range counts {
		pq[count] = append(pq[count], value)
	}

	answer := make([]int, 0, k)

	for i := len(pq) - 1; i > 0; i-- {
		if len(pq[i]) > 0 {
			for _, num := range pq[i] {
				answer = append(answer, num)
				if k == len(answer) {
					return answer
				}
			}
		}
	}
	return []int{}
}

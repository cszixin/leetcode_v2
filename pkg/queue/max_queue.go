package queue

type MaxQueue struct {
	data []int
	real []int
}

func Constructor() MaxQueue {
	q := new(MaxQueue)
	q.data = make([]int, 0)
	q.real = make([]int, 0)

	return *q
}

func (this *MaxQueue) Max_value() int {
	if len(this.data) == 0 {
		return -1
	}
	return this.data[0]
}

func (this *MaxQueue) Push_back(value int) {
	for i := len(this.data) - 1; i >= 0; i-- {
		if value > this.data[i] {
			this.data = this.data[:len(this.data)-1]
		}

	}
	this.data = append(this.data, value)
	this.real = append(this.real, value)

}

func (this *MaxQueue) Pop_front() int {
	if len(this.data) == 0 {
		return -1
	}

	tmp := this.real[0]
	this.real = this.real[1:]
	if tmp == this.data[0] {
		this.data = this.data[1:]
	}

	return tmp

}

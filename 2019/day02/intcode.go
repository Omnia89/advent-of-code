package main

type IntCode struct {
	ip      int
	program []int
}

func (i *IntCode) next() (exit bool) {
	if i.ip >= len(i.program) {
		return true
	}

	switch i.program[i.ip] {
	case 99:
		return true
	case 1:
		i.add()
	case 2:
		i.multi()
	}

	return false
}

func (i *IntCode) add() {
	if i.ip+3 >= len(i.program) {
		i.ip = len(i.program)
		return
	}

	a1 := i.program[i.program[i.ip+1]]
	a2 := i.program[i.program[i.ip+2]]
	o := i.program[i.ip+3]

	i.program[o] = a1 + a2

	i.ip += 4
}

func (i *IntCode) multi() {
	if i.ip+3 >= len(i.program) {
		i.ip = len(i.program)
		return
	}

	a1 := i.program[i.program[i.ip+1]]
	a2 := i.program[i.program[i.ip+2]]
	o := i.program[i.ip+3]

	i.program[o] = a1 * a2

	i.ip += 4
}

func (i *IntCode) run(output int) int {
	for {
		val := i.next()
		if val {
			break
		}
	}

	return i.program[output]
}

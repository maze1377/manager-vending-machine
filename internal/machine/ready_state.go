package machine

type ReadyState struct {
	machine *VendingMachine
}

func (r *ReadyState) insertMoney() State {
	return NewPaymentState(r.machine)
}

func (r *ReadyState) interactWithMenu() State {
	r.machine.println("Please insert a coin first")
	return r
}

func (r *ReadyState) dispenseProduct() State {
	r.machine.println("Please insert a coin first")
	return r
}

func (r *ReadyState) dispenseMoney() State {
	r.machine.println("Please insert a coin first")
	return r
}

func NewReadyState(machine *VendingMachine) State {
	machine.println("Ready")
	machine.println("-----------------------")
	return &ReadyState{machine}
}

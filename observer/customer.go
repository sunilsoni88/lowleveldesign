package observer

type Customer struct {
	id string
}

func NewCustomer(id string) Observer {
	return &Customer{id: id}
}

func (c *Customer) UpdateOnEvent(itemName string) {
	println("Sending email to customer ", c.id, " for item ", itemName)
}

func (c *Customer) GetId() string {
	return c.id
}

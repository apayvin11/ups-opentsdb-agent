package mockmodbus

type MockModbus struct {
	coils                 [0xffff]byte
	holdingRegistersBytes [0xffff * 2]byte
}

func New() *MockModbus {
	return &MockModbus{}
}

func (m *MockModbus) ReadCoils(address, quantity uint16) (results []byte, err error) {
	numBytes := quantity/8 +1
	return m.coils[address : address+numBytes], nil
}

func (m *MockModbus) ReadDiscreteInputs(address, quantity uint16) (results []byte, err error) {
	return nil, nil
}

func (m *MockModbus) WriteSingleCoil(address, value uint16) (results []byte, err error) {
	m.coils[address] = byte(value)
	return nil, nil
}

func (m *MockModbus) WriteMultipleCoils(address, quantity uint16, value []byte) (results []byte, err error) {
	return nil, nil
}

func (m *MockModbus) ReadInputRegisters(address, quantity uint16) (results []byte, err error) {
	return nil, nil
}

func (m *MockModbus) ReadHoldingRegisters(address, quantity uint16) (results []byte, err error) {
	return m.holdingRegistersBytes[address*2 : (address+quantity)*2], nil
}

func (m *MockModbus) WriteSingleRegister(address, value uint16) (results []byte, err error) {
	return nil, nil
}

func (m *MockModbus) WriteMultipleRegisters(address, quantity uint16, value []byte) (results []byte, err error) {
	copy(m.holdingRegistersBytes[address*2:], value)
	return nil, nil
}

func (m *MockModbus) ReadWriteMultipleRegisters(readAddress, readQuantity, writeAddress, writeQuantity uint16, value []byte) (results []byte, err error) {
	return nil, nil
}

func (m *MockModbus) MaskWriteRegister(address, andMask, orMask uint16) (results []byte, err error) {
	return nil, nil
}

func (m *MockModbus) ReadFIFOQueue(address uint16) (results []byte, err error) {
	return nil, nil
}

//temperature reads temperature from sensor
package temperature

type Reader interface {
	Read() float32
}

func NewReader() *DummyReader {
	return &DummyReader{}
}

//DummyReader dummy reader until sensor arrived
type DummyReader struct {
}

func (reader *DummyReader) Read() float32 {
	return 0.0
}

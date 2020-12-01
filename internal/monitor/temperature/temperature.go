//temperature reads temperature from sensor
package temperature

import (
	"github.com/d2r2/go-bsbmp"
	i2c "github.com/d2r2/go-i2c"
	"github.com/d2r2/go-logger"
	"log"
	"time"
)

type Measurement struct {
	Temp      float32
	Humidity  float32
	Timestamp time.Time
}

//Reader reads single temperature value
type Reader interface {
	Read() Measurement
	Close()
}

func newDummyReader() *dummyReader {
	return &dummyReader{}
}

//DummyReader dummy reader until sensor arrived
type dummyReader struct {
}

func (reader *dummyReader) Read() Measurement {
	return Measurement{}
}

func (reader *dummyReader) Close() {

}

//i2cReader
type i2cReader struct {
	i2c    *i2c.I2C
	sensor *bsbmp.BMP
}

func newI2cReader(bus int, addr uint8) (*i2cReader, error) {
	i2c, err := i2c.NewI2C(addr, bus)
	if err != nil {
		return nil, err
	}
	logger.ChangePackageLogLevel("i2c", logger.InfoLevel)
	logger.ChangePackageLogLevel("bsbmp", logger.InfoLevel)
	sensor, err := bsbmp.NewBMP(bsbmp.BME280, i2c)
	if err != nil {
		return nil, err
	}
	return &i2cReader{i2c: i2c, sensor: sensor}, nil
}

func (r *i2cReader) Read() Measurement {
	t, _ := r.sensor.ReadTemperatureC(bsbmp.ACCURACY_STANDARD)
	_, h, _ := r.sensor.ReadHumidityRH(bsbmp.ACCURACY_STANDARD)
	return Measurement{
		Temp:      t,
		Humidity:  h,
		Timestamp: time.Now(),
	}

}

func (r *i2cReader) Close() {
	r.i2c.Close()
}

//TimedReader reads temperature by interval
type IntervalReader interface {
	Start()
	Stop()
}

type intervalReadrImpl struct {
	ticker  *time.Ticker
	reader  Reader
	handler func(Measurement)
}

func (r *intervalReadrImpl) Start() {
	go func(c <-chan time.Time) {
		for range c {
			go func() {
				m := r.reader.Read()
				r.handler(m)
			}()
		}
		log.Print("Channel closed")
	}(r.ticker.C)
}

func (r *intervalReadrImpl) Stop() {
	r.ticker.Stop()
	r.reader.Close()
}
func NewIntervalReader(d time.Duration, bus int, address uint8, handler func(Measurement)) (*intervalReadrImpl, error) {
	r, err := newI2cReader(bus, address)
	if err != nil {
		return nil, err
	}
	//r := newDummyReader()
	return &intervalReadrImpl{
		ticker:  time.NewTicker(d),
		reader:  r,
		handler: handler,
	}, nil
}

package helper

import (
	"database/sql/driver"
	"errors"
	"strconv"
)

type Point struct {
	X, Y float64
}

// Implement the Valuer interface.
func (p Point) Value() (driver.Value, error) {
	out := []byte{'('}
	out = strconv.AppendFloat(out, p.X, 'f', -1, 64)
	out = append(out, ',')
	out = strconv.AppendFloat(out, p.Y, 'f', -1, 64)
	out = append(out, ')')
	return out, nil
}

// Implement the Scanner interface.
func (p *Point) Scan(src interface{}) (err error) {
	var data []byte
	switch src := src.(type) {
	case []byte:
		data = src
	case string:
		data = []byte(src)
	case nil:
		return nil
	default:
		return errors.New("(*Point).Scan: unsupported data type")
	}

	if len(data) == 0 {
		return nil
	}

	data = data[1 : len(data)-1] // drop the surrounding parentheses
	for i := 0; i < len(data); i++ {
		if data[i] == ',' {
			if p.X, err = strconv.ParseFloat(string(data[:i]), 64); err != nil {
				return err
			}
			if p.Y, err = strconv.ParseFloat(string(data[i+1:]), 64); err != nil {
				return err
			}
			break
		}
	}
	return nil
}

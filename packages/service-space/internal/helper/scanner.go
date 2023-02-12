package helper

import (
	"database/sql/driver"
	"errors"
	"fmt"
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
		fmt.Println("ENTERED 10", src)
		data = src
	case string:
		fmt.Println("ENTERED 11")
		data = []byte(src)
	case nil:
		fmt.Println("ENTERED 12")
		return nil
	default:
		fmt.Println("ENTERED 13", src)
		return errors.New("(*Point).Scan: unsupported data type")
	}

	if len(data) == 0 {
		return nil
	}

	fmt.Println("ENTERED 14")

	data = data[1 : len(data)-1] // drop the surrounding parentheses
	for i := 0; i < len(data); i++ {
		if data[i] == ',' {
			if p.X, err = strconv.ParseFloat(string(data[:i]), 64); err != nil {
				fmt.Println("ENTERED 15")
				return err
			}
			if p.Y, err = strconv.ParseFloat(string(data[i+1:]), 64); err != nil {
				fmt.Println("ENTERED 16")
				return err
			}
			break
		}
	}
	fmt.Println("ENTERED 17", data)
	return nil
}

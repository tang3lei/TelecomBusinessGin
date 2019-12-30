package db

import (
	"fmt"
	"strconv"
)

type QueryArg struct {
	ParamsMap map[string][]string
	argKey    []string
	argValue  []interface{}
	Limit     int
	Offset    int
	Order     string
	error     error
}

func (p *QueryArg) Equal(key string, column string) *QueryArg {
	p.equal(key, column)
	return p
}

func (p *QueryArg) MustEqual(key string, column string) *QueryArg  {
	err := p.equal(key, column)
	p.error = err
	return p
}

func (p *QueryArg) equal(key string, column string) error {
	val, ok := p.ParamsMap[key]
	if !ok {
		return fmt.Errorf("need request param %v", key)
	}
	p.argKey = append(p.argKey, fmt.Sprintf("%v = ?", column))
	p.argValue = append(p.argValue, val[0])
	return nil
}

func (p *QueryArg) Like(key string, column string) *QueryArg  {
	p.like(key, column)
	return p
}

func (p *QueryArg) MustLike(key string, column string)  *QueryArg {
	err := p.like(key, column)
	p.error = err
	return p
}

func (p *QueryArg) like(key string, column string) error {
	val, ok := p.ParamsMap[key]
	if !ok {
		return fmt.Errorf("need request param %v", key)
	}
	p.argKey = append(p.argKey, fmt.Sprintf("%v like ?", "%"+column+"%"))
	p.argValue = append(p.argValue, val[0])
	return nil
}

func (p *QueryArg) EqualInt(key string, column string) *QueryArg  {
	p.equalInt(key, column)
	return p
}

func (p *QueryArg) MustEqualInt(key string, column string) *QueryArg  {
	err := p.equalInt(key, column)
	p.error = err
	return p
}

func (p *QueryArg) equalInt(key string, column string) error {
	val, ok := p.ParamsMap[key]
	if !ok {
		return fmt.Errorf("need request param %v", key)
	}
	valInt,err := strconv.Atoi(val[0])
	if err != nil {
		return err
	}
	p.argKey = append(p.argKey, fmt.Sprintf("%v = ?", column))
	p.argValue = append(p.argValue, valInt)
	return nil
}

func (p *QueryArg) EqualFloat64(key string, column string) *QueryArg  {
	p.equalFloat64(key, column)
	return p
}

func (p *QueryArg) MustEqualFloat64(key string, column string) *QueryArg  {
	err := p.equalFloat64(key, column)
	p.error = err
	return p
}

func (p *QueryArg) equalFloat64(key string, column string) error {
	val, ok := p.ParamsMap[key]
	if !ok {
		return fmt.Errorf("need request param %v", key)
	}
	valFloat,err := strconv.ParseFloat(val[0],64)
	if err != nil {
		return err
	}
	p.argKey = append(p.argKey, fmt.Sprintf("%v = ?", column))
	p.argValue = append(p.argValue, valFloat)
	return nil
}

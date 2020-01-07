//============================================================
// 描述:
// 作者: Simon
// 日期: 2020/1/6 6:40 下午
//
//============================================================

package libs

import "database/sql/driver"

type MSTimeStamp struct {
	ts []uint8
	v  bool
}

func (s *MSTimeStamp) Scan(value interface{}) error {
	s.ts, s.v = value.([]uint8)
	if s.v {
		return nil
	}

	return nil
}

func (s MSTimeStamp) Value() (driver.Value, error) {

	return nil, nil
}

func (s *MSTimeStamp) UnmarshalJSON(data []byte) error {
	return nil
}

func (s MSTimeStamp) MarshalJSON() ([]byte, error) {

	return nil, nil
}

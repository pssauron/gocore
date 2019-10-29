//============================================================
// 描述:
// 作者: Simon
// 日期: 2019/10/30 5:02 下午
//
//============================================================

package main

import (
	"fmt"
	"time"

	"github.com/pssauron/gocore/libs"
	"github.com/pssauron/gocore/stores"
)

type Child struct {
	F string `db:"fs"`
	G int    `db:"gs"`
}

type Test struct {
	A *libs.Int   `db:"A"`
	B libs.String `db:"B"`
	Child
}

type UA struct {
	C libs.Bool `db:"C"`
	D libs.Time `db:"D"`
	*Test
}

func main() {
	a := new(libs.Int)
	a.Set(1)
	t := UA{
		C: libs.NewBool(true),
		D: libs.NewTime(time.Now()),
		Test: &Test{
			A: a,
			B: libs.String{},
			Child: Child{
				F: "sss",
				G: 0,
			},
		},
	}
	fs, err := stores.GetDBFields(t)

	if err != nil {
		fmt.Println(err)
		return
	}

	for _, item := range fs {
		fmt.Printf("%s:%s:%v", item.Name, item.Col, item.IsEmpty)
		fmt.Println()
	}
}

//============================================================
// 描述:
// 作者: Simon
// 日期: 2019/10/28 5:36 下午
//
//============================================================

package rs

import "github.com/pssauron/gocore/libs"

type PageData struct {
	Page libs.Int `json:"page"`
	Size libs.Int `json:"size"`
}

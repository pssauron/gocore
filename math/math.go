//============================================================
// 描述:
// 作者: Simon
// 日期: 2019/12/9 11:40 上午
//
//============================================================

package math

import (
	"errors"

	"github.com/dengsgo/math-engine/engine"
)

func Calculation(expr string) (float64, error) {
	toks, err := engine.Parse(expr)
	if err != nil {
		return 0, errors.New("公式解析错误")
	}
	ast := engine.NewAST(toks, expr)
	if ast.Err != nil {
		return 0, errors.New("公式解析错误")
	}
	ar := ast.ParseExpression()
	if ast.Err != nil {
		return 0, errors.New("公司解析错误")
	}
	r := engine.ExprASTResult(ar)

	return r, nil
}

package models

import (
	"github.com/gobuffalo/suite"
)

type ActionSuite struct {
	*suite.Action
}

func (as *ActionSuite) Test_Pizza() {
	as.LoadFixture("lots of pizzas")
	res := as.HTML("/pizzas").Get()
	as.Equal(200, res.Code)
	body := res.Body.String()
	as.Contains(body, "pizza #1")
	as.Contains(body, "pizza #2")
}

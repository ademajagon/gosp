package types

type Token string

type Expr interface{}

type Symbol string

type Number float64

type List []Expr

type Function func(...Expr) (Expr, error)

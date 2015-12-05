package player

type Enemy struct {
	Actor
	Hatred int
	target []*Actor
}

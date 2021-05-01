package domain

type Game struct {
	Id          string
	Name        string
	Description string
	SectionId   string
	Format      string
	Type        string
}

type Games struct {
	Games []Game
}

type GameList struct {
	Result Games
}

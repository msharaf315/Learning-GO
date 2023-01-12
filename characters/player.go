package characters

type Player struct {
	MaxHp  float64
	Hp     float64
	Attack float64
	Dead   bool
}

func NewPlayer() *Player {
	return &Player{
		100,
		100,
		15,
		false,
	}
}

func (player *Player) Heal() {
	player.Hp += player.MaxHp * 10
	if player.Hp >= player.MaxHp {
		player.Hp = player.MaxHp
	}
}

func (player *Player) Rest() {
	player.MaxHp += player.MaxHp * 10
}

func (player *Player) Forge() {
	player.Attack += player.Attack * 10
}

func (player *Player) TakeDamage(Attack float64) {
	player.Hp -= Attack
	player.Dead = player.Hp <= 0
}

package characters

type Dragon struct {
	MaxHp  float64
	Hp     float64
	Attack float64
	Dead   bool
}

func NewDragon(level int) *Dragon {

	return &Dragon{
		50 * (1 + (float64(level) / 100)),
		50 * (1 + (float64(level) / 100)),
		20 * (1 + (float64(level) / 100)),
		false,
	}
}

func (dragon *Dragon) Heal() {
	dragon.Hp += dragon.MaxHp * 10
	if dragon.Hp >= dragon.MaxHp {
		dragon.Hp = dragon.MaxHp
	}
}

func (dragon *Dragon) Rest() {
	dragon.MaxHp += dragon.MaxHp * 10
}

func (dragon *Dragon) Forge() {
	dragon.Attack += dragon.Attack * 10
}

func (dragon *Dragon) TakeDamage(Attack float64) {
	dragon.Hp -= Attack
	dragon.Dead = dragon.Hp <= 0
}

func (dragon *Dragon) BreathFire(player *Player) {
	dragon.Attack = float64(dragon.Attack) * 1.5
	player.TakeDamage(dragon.Attack)
}

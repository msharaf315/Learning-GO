package characters

type Goblin struct {
	MaxHp  float64
	Hp     float64
	Attack float64
	Dead   bool
}

func NewGoblin() *Goblin {
	return &Goblin{
		30,
		30,
		5,
		false,
	}
}

func (goblin *Goblin) Heal() {
	goblin.Hp += goblin.MaxHp * 10
	if goblin.Hp >= goblin.MaxHp {
		goblin.Hp = goblin.MaxHp
	}
}

func (goblin *Goblin) Rest() {
	goblin.MaxHp += goblin.MaxHp * 10
}

func (goblin *Goblin) Forge() {
	goblin.Attack += goblin.Attack * 10
}

func (goblin *Goblin) TakeDamage(Attack float64) {
	goblin.Hp -= Attack
	goblin.Dead = goblin.Hp <= 0
}

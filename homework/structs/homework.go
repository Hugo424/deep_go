package main

type Option func(*GamePerson)

func WithName(name string) func(*GamePerson) {
	return func(person *GamePerson) {
		copy(person.personInfo[:], name)
	}
}

func WithCoordinates(x, y, z int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.x, person.y, person.z = int32(x), int32(y), int32(z)
	}
}

func WithGold(gold int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.gold = uint32(gold)
	}
}

func WithMana(mana int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.stats = person.stats | uint32(mana)<<22
	}
}

func WithHealth(health int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.stats = person.stats | uint32(health)<<12
	}
}

func WithRespect(respect int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.stats = person.stats | uint32(respect)<<8
	}
}

func WithStrength(strength int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.stats = person.stats | uint32(strength)<<4
	}
}

func WithExperience(experience int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.stats = person.stats | uint32(experience)
	}
}

func WithLevel(level int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.personInfo[43] = person.personInfo[43] | uint8(level)
	}
}

func WithHouse() func(*GamePerson) {
	return func(person *GamePerson) {
		person.personInfo[42] = person.personInfo[42] | uint8(1)
	}
}

func WithGun() func(*GamePerson) {
	return func(person *GamePerson) {
		person.personInfo[42] = person.personInfo[42] | uint8(1)<<1
	}
}

func WithFamily() func(*GamePerson) {
	return func(person *GamePerson) {
		person.personInfo[42] = person.personInfo[42] | uint8(1)<<2
	}
}

func WithType(personType int) func(*GamePerson) {
	return func(person *GamePerson) {
		person.personInfo[42] = person.personInfo[42] | uint8(personType)<<3
	}
}

const (
	BuilderGamePersonType = iota
	BlacksmithGamePersonType
	WarriorGamePersonType
)

type GamePerson struct {
	personInfo [44]byte //[0-41] - name; [42] - house, gun, family, type; [43] - level
	x, y, z    int32
	gold       uint32
	stats      uint32 //4bit - experience; 4 - strength; 4 - respect; 10 - health; 10 - mana
}

func NewGamePerson(options ...Option) GamePerson {
	gamePerson := GamePerson{}

	for _, option := range options {
		option(&gamePerson)
	}
	return gamePerson
}

func (p *GamePerson) Name() string {
	return string(p.personInfo[:42])
}

func (p *GamePerson) X() int {
	return int(p.x)
}

func (p *GamePerson) Y() int {
	return int(p.y)
}

func (p *GamePerson) Z() int {
	return int(p.z)
}

func (p *GamePerson) Gold() int {
	return int(p.gold)
}

func (p *GamePerson) Mana() int {
	return int(p.stats >> 22 & 0x3FF)
}

func (p *GamePerson) Health() int {
	return int(p.stats >> 12 & 0x3FF)
}

func (p *GamePerson) Respect() int {
	return int(p.stats >> 8 & 0xF)
}

func (p *GamePerson) Strength() int {
	return int(p.stats >> 4 & 0xF)
}

func (p *GamePerson) Experience() int {
	return int(p.stats) & 0xF
}

func (p *GamePerson) Level() int {
	return int(p.personInfo[43])
}

func (p *GamePerson) HasHouse() bool {
	return (p.personInfo[42] & 1) == 1
}

func (p *GamePerson) HasGun() bool {
	return (p.personInfo[42] >> 1 & 1) == 1
}

func (p *GamePerson) HasFamilty() bool {
	return (p.personInfo[42] >> 2 & 1) == 1
}

func (p *GamePerson) Type() int {
	return int(p.personInfo[42] >> 3)
}

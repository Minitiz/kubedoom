package monsters

type IMonster interface {
	Add(Monster)
	Delete(Monster)
}

type Monster string

type Monsters []Monster

func (m *Monsters) Add(monsterName Monster) {
	*m = append(*m, monsterName)
	// m = &new
}

func (m *Monsters) Delete(monsterName Monster) {
	tmp := *m
	for index, monster := range tmp {
		if monster == monsterName {
			tmp = append(tmp[:index], tmp[index+1:]...)
		}
	}
	*m = tmp
}

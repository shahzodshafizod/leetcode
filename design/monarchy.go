package design

/*
Problem:
Given the following interface, implement its methods.
interface Monarchy {
	void birth(String child, String parent);
	void death(String name);
	List<String> getOrderOfSuccession();
}

Monarchy Family Tree
	Jake (monarch):
		1. Catherine
			1. Jane
				1. Farah
			2. Mark
		2. Tom
		3. Celine
			1. Peter

new Monarchy('Jake')

birth('Catherine', 'Jake')
birth('Tom', 'Jake')
birth('Celine', 'Jake')
birth('Jane', 'Catherine')
birth('Peter', 'Celine')
birth('Farah', 'Jane')
birth('Mark', 'Catherine')

getOrderOfSuccession()
	: ['Jake', 'Catherine', 'Jane', 'Farah', 'Mark', 'Tom', 'Celine', 'Peter']

death('Jake')
death('Jane')

getOrderOfSuccession()
	: ['Catherine', 'Farah', 'Mark', 'Tom', 'Celine', 'Peter']

Step 1: Verify the constraints
	- Can we implement helper classes/objects?
		: Yes, you can use any features you see fit.
	- What if there are many same names?
		: Assume all names are unique.
Step 2: Write out some test cases

*/

type Monarchy interface {
	Birth(child string, parent string)
	Death(name string)
	GetOrderOfSuccession() []string
}

type member struct {
	name     string
	alive    bool
	children []*member
}

type monarchy struct {
	king    *member
	members map[string]*member
}

func NewMonarchy(kingName string) Monarchy {
	king := &member{
		name:     kingName,
		alive:    true,
		children: make([]*member, 0),
	}

	return &monarchy{
		king:    king,
		members: map[string]*member{kingName: king},
	}
}

func (m *monarchy) Birth(child string, parent string) {
	theParent := m.members[parent]
	if theParent == nil {
		return
	}

	newChild := &member{
		name:     child,
		alive:    true,
		children: make([]*member, 0),
	}
	theParent.children = append(theParent.children, newChild)
	m.members[child] = newChild
}

func (m *monarchy) Death(name string) {
	if theMember := m.members[name]; theMember != nil {
		theMember.alive = false
	}
}

func (m *monarchy) GetOrderOfSuccession() []string {
	return m._getOrderOfSuccession(m.king)
}

func (m *monarchy) _getOrderOfSuccession(current *member) []string {
	values := make([]string, 0)
	if current.alive {
		values = append(values, current.name)
	}

	for _, child := range current.children {
		values = append(values, m._getOrderOfSuccession(child)...)
	}

	return values
}

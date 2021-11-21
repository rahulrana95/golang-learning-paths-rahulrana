package concepts

type LangRating struct {
	language string
	rating   int
}

type Skill struct {
	name      string
	languages []string
}

type User struct {
	name           string
	firstName      string
	lastName       string
	currentCompany string
	skill          Skill
}

type Users struct {
}

func (users *Users) AddUser(name string) {

}

func AddLanguages(lang []string) {

}

// Map Examples

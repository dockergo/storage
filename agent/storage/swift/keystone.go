//Site: https://github.com/flyaways
//Auth: Flyaway
//Date 27/09/2016 13:22 Beijing

package swift

//request token
type Domain struct {
	Name string `json:"name"`
}

type User struct {
	Domain   Domain `json:"domain"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type Password struct {
	User User `json:"user"`
}

type Identity struct {
	Methods  []string `json:"methods"`
	Password Password `json:"password"`
}

type Project struct {
	Domain Domain `json:"domain"`
	Name   string `json:"name"`
}

type Scope struct {
	Project Project `json:"project"`
}

type Auth struct {
	Identity Identity `json:"identity"`
	Scope    Scope    `json:"scope"`
}
type AuthR struct {
	Auth Auth `json:"auth"`
}

//response token
type ProjectT struct {
	Id string `json:"id"`
}

type Token struct {
	ProjectT ProjectT `json:"project"`
}

type TokenR struct {
	Token Token `json:"token"`
}

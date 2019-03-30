package data

type Id struct {
	Scheme string `jute:"scheme"`
	Id     string `jute:"id"`
}

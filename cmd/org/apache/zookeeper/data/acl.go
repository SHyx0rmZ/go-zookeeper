package data

type ACL struct {
	Perms uint32 `jute:"perms"`
	Id    Id     `jute:"id"`
}

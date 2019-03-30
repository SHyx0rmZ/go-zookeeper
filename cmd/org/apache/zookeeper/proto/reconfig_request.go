package proto

type ReconfigRequest struct {
	JoiningServers string `jute:"joiningServers"`
	LeavingServers string `jute:"leavingServers"`
	NewMembers     string `jute:"newMembers"`
	CurConfigId    uint64 `jute:"curConfigId"`
}

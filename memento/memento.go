package memento

import "fmt"

type Memento interface {
	Tag() string
	Restore()
}

type rpgArchive struct {
	tag           string
	roleState     []string
	scenarioState string
	rgp           *RolePlayGame
}

func newRpgArchive(tag string, roleState []string, scenarioState string, rgp *RolePlayGame) *rpgArchive {
	return &rpgArchive{tag: tag, roleState: roleState, scenarioState: scenarioState, rgp: rgp}
}

func (r *rpgArchive) Tag() string {
	return r.tag
}

func (r *rpgArchive) Restore() {
	r.rgp.SetRoleState(r.roleState)
	r.rgp.SetScenarioState(r.scenarioState)
}

type RPGArchiveManager struct {
	archives map[string]Memento
}

func NewRPGArchiveManager() *RPGArchiveManager {
	return &RPGArchiveManager{archives: make(map[string]Memento)}
}

func (r *RPGArchiveManager) Reload(tag string) {
	if archive, ok := r.archives[tag]; ok {
		fmt.Printf("重新加载%s;\n", tag)
		archive.Restore()
	}
}

func (r *RPGArchiveManager) Put(memento Memento) {
	r.archives[memento.Tag()] = memento
}

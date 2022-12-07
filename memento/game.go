package memento

import "fmt"

type Originator interface {
	Save(tag string) Memento
}

type RolePlayGame struct {
	name          string   // 角色名字
	roleState     []string // 角色状态
	scenarioState string   // 游戏场景
}

func NewRoleSystem(name string, roleName string) *RolePlayGame {
	return &RolePlayGame{name: name, roleState: []string{roleName, "血量100"}, scenarioState: "关卡1"}
}

func (r *RolePlayGame) Save(tag string) Memento {
	return newRpgArchive(tag, r.roleState, r.scenarioState, r)
}

func (r *RolePlayGame) SetRoleState(roleState []string) {
	r.roleState = roleState
}

func (r *RolePlayGame) SetScenarioState(scenarioState string) {
	r.scenarioState = scenarioState
}

// 游戏信息
func (r *RolePlayGame) String() string {
	return fmt.Sprintf("在%s游戏中，玩家使用%s，%s,%s", r.name, r.roleState[0], r.roleState[1], r.scenarioState)
}

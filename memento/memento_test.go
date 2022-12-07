package memento

import (
	"fmt"
	"testing"
)

/*
1. 需要一个对象统一管理存档，所以有 archiveManger
2. 存档对象需要进行抽象，所以有 Memento 接口，而rpgarchiv 实现了这个接口,所有可以处理 存档和恢复两个操作
3. 存档的对象是游戏，所以有游戏对象
4. 游戏--》存档之间有信息转化，所有游戏对象实现了Originator,这个接口就是把游戏当前状态保存为存档对象


*/

func TestMemento(t *testing.T) {
	rpgManager := NewRPGArchiveManager()
	rpg := NewRoleSystem("dota2", "Lion")
	fmt.Println(rpg)
	rpgManager.Put(rpg.Save("第一关存档"))

	rpg.SetRoleState([]string{"Lion", "死亡"})
	rpg.SetScenarioState("第一关闯关失败")
	fmt.Println(rpg)

	rpgManager.Reload("第一关存档")
	fmt.Println(rpg)
}

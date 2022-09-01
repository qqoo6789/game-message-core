package xlsxTable

func TableModels() []interface{} {
	return []interface{}{
		BuffTable{},
		ChatTable{},
		DropTable{},
		GameValueTable{},
		ItemTable{},
		MonsterTable{},
		RewardTable{},
		RoleLvTable{},
		SlotLvTable{},
		TaskTable{},
		TaskListTable{},
	}
}

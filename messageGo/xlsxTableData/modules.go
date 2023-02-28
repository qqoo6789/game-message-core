package xlsxTable

func TableModels() []interface{} {
	return []interface{}{
		ChatTableRow{},
		GameValueTable{},
		ItemTable{},
		MonsterTableRow{},
		RewardTableRow{},
		RoleTableRow{},
		SlotLvTableRow{},
		TaskTableRow{},
		TaskListTableRow{},
		SceneAreaRow{},
		TalentTreeRow{},
		EntityAttributeRow{},
		ItemEatableRow{},
	}
}

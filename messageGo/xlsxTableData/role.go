package xlsxTable

import (
	"encoding/json"
	"time"
)

type AttributeData struct {
	Id    uint  `json:"id"`
	Value int32 `json:"value"` // 属性值
}

type AttributeDataList struct {
	List []AttributeData `json:"list"`
}

type RoleTableRow struct {
	Id                    uint32    `gorm:"primaryKey" json:"id"`
	RoleName              string    `json:"roleName"`              // 角色名
	RoleSex               string    `json:"roleSex"`               // 角色性别
	AssetId               uint32    `json:"assetId"`               // 角色资源id
	InitialAttributesJson string    `json:"initialAttributesJson"` // 初始属性
	CreatedAt             time.Time `json:"createdAt"`             // 过期判断条件

	InitialAttributes *AttributeDataList `gorm:"-" json:"-"`
}

func (p *RoleTableRow) GetInitialAttributes() (list []AttributeData, err error) {
	if len(p.InitialAttributesJson) < 2 {
		return nil, nil
	}

	attrList := &AttributeDataList{}
	err = json.Unmarshal([]byte(p.InitialAttributesJson), attrList)
	if err != nil {
		return nil, err
	}
	p.InitialAttributes = attrList
	return p.InitialAttributes.List, nil
}

func (p *RoleTableRow) SetInitialAttributes(data *AttributeDataList) error {
	bs, err := json.Marshal(data)
	if err != nil {
		return err
	}
	p.InitialAttributesJson = string(bs)
	p.InitialAttributes = data
	return nil
}

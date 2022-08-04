package dnd

type baseitem struct {
	Name           string                  `json:"name,omitempty"`
	Source         string                  `json:"source,omitempty"`
	Page           int                     `json:"page,omitempty"`
	Type           string                  `json:"type,omitempty"`
	Rarity         string                  `json:"rarity,omitempty"`
	Weight         baseitem_weight         `json:"weight,omitempty"`
	Weaponcategory string                  `json:"weaponCategory,omitempty"`
	Age            string                  `json:"age,omitempty"`
	Property       []string                `json:"property,omitempty"`
	Range          string                  `json:"range,omitempty"`
	Reload         int                     `json:"reload,omitempty"`
	Dmg1           string                  `json:"dmg1,omitempty"`
	Dmgtype        string                  `json:"dmgType,omitempty"`
	Firearm        bool                    `json:"firearm,omitempty"`
	Weapon         bool                    `json:"weapon,omitempty"`
	Ammotype       string                  `json:"ammoType,omitempty"`
	Eng_name       string                  `json:"ENG_name,omitempty"`
	Srd            bool                    `json:"srd,omitempty"`
	Value          baseitem_value          `json:"value,omitempty"`
	Packcontents   []baseitem_packContents `json:"packContents,omitempty"`
	Dmg2           string                  `json:"dmg2,omitempty"`
	Axe            bool                    `json:"axe,omitempty"`
	Ac             int                     `json:"ac,omitempty"`
	Armor          bool                    `json:"armor,omitempty"`
	Entries        []baseitem_entries      `json:"entries,omitempty"`
	Strength       string                  `json:"strength,omitempty"`
	Stealth        bool                    `json:"stealth,omitempty"`
	Club           bool                    `json:"club,omitempty"`
	Scftype        string                  `json:"scfType,omitempty"`
	Dagger         bool                    `json:"dagger,omitempty"`
	Othersources   []baseitem_otherSources `json:"otherSources,omitempty"`
	Sword          bool                    `json:"sword,omitempty"`
	Hasfluff       bool                    `json:"hasFluff,omitempty"`
	Crossbow       bool                    `json:"crossbow,omitempty"`
	Spear          bool                    `json:"spear,omitempty"`
	Hammer         bool                    `json:"hammer,omitempty"`
	Bow            bool                    `json:"bow,omitempty"`
	Mace           bool                    `json:"mace,omitempty"`
	Net            bool                    `json:"net,omitempty"`
	Staff          bool                    `json:"staff,omitempty"`
}
type baseitem_weight struct {
	Int_raw     int     `json:"int_raw,omitempty"`
	Float32_raw float32 `json:"float32_raw,omitempty"`
}
type baseitem_value struct {
	Int_raw     int     `json:"int_raw,omitempty"`
	Float32_raw float32 `json:"float32_raw,omitempty"`
}
type baseitem_packContents struct {
	Item     string `json:"item,omitempty"`
	Quantity int    `json:"quantity,omitempty"`
}
type baseitem_entries struct {
	Type       string   `json:"type,omitempty"`
	Name       string   `json:"name,omitempty"`
	Entries    []string `json:"entries,omitempty"`
	String_raw string   `json:"string_raw,omitempty"`
}
type baseitem_otherSources struct {
	Source string `json:"source,omitempty"`
	Page   int    `json:"page,omitempty"`
}

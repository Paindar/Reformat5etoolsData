package dnd

type spell struct {
	Name                string                      `json:"name,omitempty"`
	Source              string                      `json:"source,omitempty"`
	Page                int                         `json:"page,omitempty"`
	Level               int                         `json:"level,omitempty"`
	School              string                      `json:"school,omitempty"`
	Time                []spell_time                `json:"time,omitempty"`
	Range               spell_range                 `json:"range,omitempty"`
	Components          spell_components            `json:"components,omitempty"`
	Duration            []spell_duration            `json:"duration,omitempty"`
	Entries             []spell_entries             `json:"entries,omitempty"`
	Entrieshigherlevel  []spell_entriesHigherLevel  `json:"entriesHigherLevel,omitempty"`
	Damageresist        []string                    `json:"damageResist,omitempty"`
	Damageimmune        []string                    `json:"damageImmune,omitempty"`
	Damagevulnerable    []string                    `json:"damageVulnerable,omitempty"`
	Savingthrow         []string                    `json:"savingThrow,omitempty"`
	Misctags            []string                    `json:"miscTags,omitempty"`
	Areatags            []string                    `json:"areaTags,omitempty"`
	Classes             spell_classes               `json:"classes,omitempty"`
	Scalingleveldice    []spell_scalingLevelDice    `json:"scalingLevelDice,omitempty"`
	Damageinflict       []string                    `json:"damageInflict,omitempty"`
	Meta                spell_meta                  `json:"meta,omitempty"`
	Abilitycheck        []string                    `json:"abilityCheck,omitempty"`
	Eng_name            string                      `json:"ENG_name,omitempty"`
	Conditioninflict    []string                    `json:"conditionInflict,omitempty"`
	Spellattack         []string                    `json:"spellAttack,omitempty"`
	Hasfluffimages      bool                        `json:"hasFluffImages,omitempty"`
	Othersources        []spell_otherSources        `json:"otherSources,omitempty"`
	Hasfluff            bool                        `json:"hasFluff,omitempty"`
	Srd                 spell_srd                   `json:"srd,omitempty"`
	Backgrounds         []spell_backgrounds         `json:"backgrounds,omitempty"`
	Races               []spell_races               `json:"races,omitempty"`
	Eldritchinvocations []spell_eldritchInvocations `json:"eldritchInvocations,omitempty"`
}
type spell_time struct {
	Number    int    `json:"number,omitempty"`
	Unit      string `json:"unit,omitempty"`
	Condition string `json:"condition,omitempty"`
}
type spell_range struct {
	Type     string               `json:"type,omitempty"`
	Distance spell_range_distance `json:"distance,omitempty"`
}
type spell_components struct {
	V bool               `json:"v,omitempty"`
	S bool               `json:"s,omitempty"`
	M spell_components_m `json:"m,omitempty"`
	R bool               `json:"r,omitempty"`
}
type spell_duration struct {
	Type          string                  `json:"type,omitempty"`
	Duration      spell_duration_duration `json:"duration,omitempty"`
	Concentration bool                    `json:"concentration,omitempty"`
	Condition     string                  `json:"condition,omitempty"`
	Ends          []string                `json:"ends,omitempty"`
}
type spell_entries struct {
	Type       string                 `json:"type,omitempty"`
	Items      []string               `json:"items,omitempty"`
	String_raw string                 `json:"string_raw,omitempty"`
	Collabels  []string               `json:"colLabels,omitempty"`
	Colstyles  []string               `json:"colStyles,omitempty"`
	Rows       [][]spell_entries_rows `json:"rows,omitempty"`
	Name       string                 `json:"name,omitempty"`
	Entries    []string               `json:"entries,omitempty"`
	Caption    string                 `json:"caption,omitempty"`
	By         string                 `json:"by,omitempty"`
}
type spell_entriesHigherLevel struct {
	Type     string   `json:"type,omitempty"`
	Name     string   `json:"name,omitempty"`
	Entries  []string `json:"entries,omitempty"`
	Eng_name string   `json:"ENG_name,omitempty"`
}
type spell_classes struct {
	Fromclasslist        []spell_classes_fromClassList        `json:"fromClassList,omitempty"`
	Fromsubclass         []spell_classes_fromSubclass         `json:"fromSubclass,omitempty"`
	Fromclasslistvariant []spell_classes_fromClassListVariant `json:"fromClassListVariant,omitempty"`
}
type spell_scalingLevelDice struct {
	Label   string                         `json:"label,omitempty"`
	Scaling spell_scalingLevelDice_scaling `json:"scaling,omitempty"`
}
type spell_meta struct {
	Technomagic bool `json:"technomagic,omitempty"`
	Ritual      bool `json:"ritual,omitempty"`
}
type spell_otherSources struct {
	Source string `json:"source,omitempty"`
	Page   int    `json:"page,omitempty"`
}
type spell_srd struct {
	Bool_raw   bool   `json:"bool_raw,omitempty"`
	String_raw string `json:"string_raw,omitempty"`
}
type spell_backgrounds struct {
	Name   string `json:"name,omitempty"`
	Source string `json:"source,omitempty"`
}
type spell_races struct {
	Name       string `json:"name,omitempty"`
	Source     string `json:"source,omitempty"`
	Basename   string `json:"baseName,omitempty"`
	Basesource string `json:"baseSource,omitempty"`
}
type spell_eldritchInvocations struct {
	Name   string `json:"name,omitempty"`
	Source string `json:"source,omitempty"`
}
type spell_range_distance struct {
	Type   string `json:"type,omitempty"`
	Amount int    `json:"amount,omitempty"`
}
type spell_components_m struct {
	Text       string `json:"text,omitempty"`
	Cost       int    `json:"cost,omitempty"`
	String_raw string `json:"string_raw,omitempty"`
	Consume    bool   `json:"consume,omitempty"`
}
type spell_duration_duration struct {
	Type   string `json:"type,omitempty"`
	Amount int    `json:"amount,omitempty"`
	Upto   bool   `json:"upTo,omitempty"`
}
type spell_entries_rows struct {
	Type       string                  `json:"type,omitempty"`
	Roll       spell_entries_rows_roll `json:"roll,omitempty"`
	String_raw string                  `json:"string_raw,omitempty"`
}
type spell_classes_fromClassList struct {
	Name            string `json:"name,omitempty"`
	Source          string `json:"source,omitempty"`
	Definedinsource string `json:"definedInSource,omitempty"`
}
type spell_classes_fromSubclass struct {
	Class    spell_classes_fromSubclass_class    `json:"class,omitempty"`
	Subclass spell_classes_fromSubclass_subclass `json:"subclass,omitempty"`
}
type spell_classes_fromClassListVariant struct {
	Name            string `json:"name,omitempty"`
	Source          string `json:"source,omitempty"`
	Definedinsource string `json:"definedInSource,omitempty"`
}
type spell_scalingLevelDice_scaling struct {
	_1  string `json:"_1,omitempty"`
	_5  string `json:"_5,omitempty"`
	_11 string `json:"_11,omitempty"`
	_17 string `json:"_17,omitempty"`
}
type spell_entries_rows_roll struct {
	Exact int  `json:"exact,omitempty"`
	Min   int  `json:"min,omitempty"`
	Max   int  `json:"max,omitempty"`
	Pad   bool `json:"pad,omitempty"`
}
type spell_classes_fromSubclass_class struct {
	Name   string `json:"name,omitempty"`
	Source string `json:"source,omitempty"`
}
type spell_classes_fromSubclass_subclass struct {
	Name        string `json:"name,omitempty"`
	Source      string `json:"source,omitempty"`
	Subsubclass string `json:"subSubclass,omitempty"`
}

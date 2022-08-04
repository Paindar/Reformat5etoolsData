package dnd

import (
	"encoding/json"
	"fmt"
	"strings"
)

func School2String(s string) string {
	var SPELL_CATALOG = map[string]string{
		"V": "塑能",
		"C": "咒法",
		"T": "变化",
		"D": "预言",
		"A": "防护",
		"N": "死灵",
		"I": "预言",
	}
	if val, ok := SPELL_CATALOG[s]; ok {
		return val
	}
	return s
}

func Time2String(s string) string {
	switch s {
	case "round":
		return "轮"
	case "minute":
		return "分钟"
	case "hour":
		return "小时"
	default:
		return s
	}
}

func Class2String(s string) string {
	s = strings.ToLower(s)
	var CLASS_CATALOG = map[string]string{
		"cleric":    "牧师",
		"artificer": "奇械师",
		"barbarian": "野蛮人",
		"bard":      "吟游诗人",
		"druid":     "德鲁伊",
		"fighter":   "战士",
		"monk":      "武僧",
		"paladin":   "圣武士",
		"ranger":    "游侠",
		"rogue":     "游荡者",
		"sorcerer":  "术士",
		"warlock":   "契术师",
		"wizard":    "法师",
	}
	if val, ok := CLASS_CATALOG[s]; ok {
		return val
	}
	return s
}

func AddToBuilderIfNotNull(builder *strings.Builder, str string, prefix string) {
	if str != "" {
		builder.WriteString(prefix)
		builder.WriteString(str)
		builder.WriteRune('\n')
	}
}

func (s *spell) String() string {
	builder := new(strings.Builder)
	builder.WriteString(fmt.Sprintf("%s	%s\n%d环 %s系", s.Name, s.Eng_name, s.Level, School2String(s.School)))
	tmp := s.Meta.String()
	if tmp != "" {
		builder.WriteRune('(')
		builder.WriteString(tmp)
		builder.WriteRune(')')
	}
	builder.WriteRune('\n')

	dataLen := len(s.Time)
	join_str := make([]string, dataLen)
	for i := 0; i < dataLen; i++ {
		join_str[i] = fmt.Sprint(s.Time[i])
	}

	dataLen = len(s.Duration)
	join_duration := make([]string, dataLen)
	for i := 0; i < dataLen; i++ {
		join_duration[i] = fmt.Sprint(s.Duration[i])
	}

	builder.WriteString(fmt.Sprintf("施法时间：%s\n射程：%s\n成分：%s\n持续时间：%s\n",
		strings.Join(join_str, " 或 "),
		s.Range,
		s.Components,
		strings.Join(join_duration, " 或 ")),
	)
	builder.WriteString("-----------------------------\n")
	dataLen = len(s.Entries)
	join_duration = make([]string, dataLen)
	for i := 0; i < dataLen; i++ {
		join_duration[i] = fmt.Sprint(s.Entries[i])
	}
	builder.WriteString(strings.Join(join_duration, "\n"))
	builder.WriteRune('\n')

	dataLen = len(s.Entrieshigherlevel)
	join_str = make([]string, dataLen)
	for i := 0; i < dataLen; i++ {
		join_str[i] = fmt.Sprint(s.Entrieshigherlevel[i])
	}

	builder.WriteString(strings.Join(join_str, "\n"))
	builder.WriteString("\n\n")
	AddToBuilderIfNotNull(builder, s.Classes.String(), "")

	dataLen = len(s.Races)
	join_str = make([]string, dataLen)
	for i := 0; i < dataLen; i++ {
		join_str[i] = fmt.Sprint(s.Races[i])
	}
	AddToBuilderIfNotNull(builder, strings.Join(join_str, ", "), "种族：")
	builder.WriteRune('\n')

	dataLen = len(s.Backgrounds)
	join_str = make([]string, dataLen)
	for i := 0; i < dataLen; i++ {
		join_str[i] = fmt.Sprint(s.Backgrounds[i])
	}
	AddToBuilderIfNotNull(builder, strings.Join(join_str, ", "), "背景：")
	builder.WriteRune('\n')

	dataLen = len(s.Eldritchinvocations)
	join_str = make([]string, dataLen)
	for i := 0; i < dataLen; i++ {
		join_str[i] = fmt.Sprint(s.Eldritchinvocations[i])
	}
	AddToBuilderIfNotNull(builder, strings.Join(join_str, ", "), "魔能祁唤：")
	builder.WriteRune('\n')

	dataLen = len(s.Othersources)
	join_str = make([]string, dataLen)
	for i := 0; i < dataLen; i++ {
		join_str[i] = fmt.Sprint(s.Othersources[i])
	}
	AddToBuilderIfNotNull(builder, s.Source, "资源：")
	builder.WriteRune(' ')
	AddToBuilderIfNotNull(builder, strings.Join(join_str, ", "), "同时记载于 ")
	return builder.String()
}

func (m spell_meta) String() string {
	if m.Ritual {
		if m.Technomagic {
			return "仪式, technomagic"
		} else {
			return "仪式"
		}
	} else {
		if m.Technomagic {
			return "technomagic"
		} else {
			return ""
		}
	}
}

func (t spell_time) String() string {
	var TEMPLATE = map[string]string{
		"bonus":    "附赠动作",
		"action":   "动作",
		"reaction": "反应",
		"minute":   "分钟",
		"hour":     "小时",
	}
	unit := TEMPLATE[t.Unit]
	if unit == "" {
		unit = t.Unit
	}
	if t.Condition != "" {
		return fmt.Sprintf("%d %s(%s)", t.Number, unit, t.Condition)
	} else {
		return fmt.Sprintf("%d %s", t.Number, unit)
	}
}

func (t spell_range) String() string {
	switch t.Type {
	case "point":
		return t.Distance.String()
	case "cone":
		return fmt.Sprintf("自身 (%s 锥形)", t.Distance.String())
	case "line":
		return fmt.Sprintf("自身 (%s 直线)", t.Distance.String())
	default:
		return fmt.Sprintf("%s %s", t.Distance.String(), t.Type)
	}
}

func (s spell_range_distance) String() string {
	return fmt.Sprintf("%d %s", s.Amount, s.Type)
}

func (c spell_components) String() (ret string) {
	if c.V {
		ret += "V "
	}
	if c.S {
		ret += "S "
	}
	if c.M.String_raw != "" {
		ret += c.M.String_raw
	} else if c.M.Text != "" {
		ret += c.M.Text
	}
	return
}

func (c *spell_components_m) UnmarshalJSON(data []byte) error {
	var dict map[string]interface{}
	err := json.Unmarshal(data, &dict)
	if err != nil { //maybe str
		return json.Unmarshal(data, &(c.String_raw))
	} else {
		c.Consume, _ = dict["consume"].(bool)
		c.Cost, _ = dict["cost"].(int)
		c.Text, _ = dict["text"].(string)
	}
	return nil
}

func (d spell_duration) String() string {
	switch d.Type {
	case "instant":
		return "即刻"
	case "time":
		var ret string
		if d.Duration.Upto {
			ret = "至多 "
		}
		ret += fmt.Sprintf("%d %s ", d.Duration.Amount, Time2String(d.Duration.Type))
		if d.Concentration {
			ret += "(专注)"
		}
		return ret
	case "permanent":
		return "直到 " + strings.Join(d.Ends, " 或 ")
	default:
		return fmt.Sprintf("%s %v %v %s %v", d.Type, d.Duration, d.Concentration, d.Condition, d.Ends)
	}
}

func (e *spell_entries) UnmarshalJSON(data []byte) error {
	var dict map[string]interface{}
	err := json.Unmarshal(data, &dict)
	if err != nil { //maybe str
		return json.Unmarshal(data, &(e.String_raw))
	} else {
		e.Type, _ = dict["type"].(string)
		e.Items, _ = dict["items"].([]string)
		e.Collabels, _ = dict["collables"].([]string)
		e.Colstyles, _ = dict["colstyles"].([]string)
		e.Rows, _ = dict["rows"].([][]spell_entries_rows)
		e.Name, _ = dict["name"].(string)
		e.Entries, _ = dict["entries"].([]string)
		e.Caption, _ = dict["caption"].(string)
		e.By, _ = dict["by"].(string)
	}
	return nil
}

func (e spell_entries) String() string {
	if e.String_raw != "" {
		return e.String_raw
	}
	switch e.Type {
	case "entries":
		return fmt.Sprintf("%s. %s", e.Name, strings.Join(e.Entries, " "))
	case "list":
		return strings.Join(e.Items, "\n")
	case "table":
		col := make([]string, len(e.Rows))

		for i := 0; i < len(e.Rows); i++ {
			cells := make([]string, len(e.Rows[i]))
			for j := 0; j < len(e.Rows[i]); j++ {
				cells[i] = fmt.Sprint(e.Rows[i][j])
			}
			col[i] = strings.Join(cells, "\t")
		}
		return fmt.Sprintf("%s\n%s\n%s", strings.Join(col, "\t"), e.Caption, strings.Join(col, "\n"))
	case "quote":
		return fmt.Sprintf("“%s” ——%s", strings.Join(e.Entries, " "), e.By)
	default:
		return e.Type
	}
}

func (e *spell_entries_rows) UnmarshalJSON(data []byte) error {
	var dict map[string]interface{}
	err := json.Unmarshal(data, &dict)
	if err != nil { //maybe str
		return json.Unmarshal(data, &(e.String_raw))
	} else {
		e.Type, _ = dict["type"].(string)
		e.Roll, _ = dict["roll"].(spell_entries_rows_roll)
	}
	return nil

}

func (t spell_entries_rows) String() string {
	if t.String_raw != "" {
		return t.String_raw
	}

	switch t.Type {
	case "cell":
		if t.Roll.Exact != 0 {
			return fmt.Sprint(t.Roll.Exact)
		} else {
			return fmt.Sprintf("%d-%d", t.Roll.Min, t.Roll.Max)
		}
	default:
		return fmt.Sprintf("%s %v", t.Type, t.Roll)
	}
}

func (t spell_entriesHigherLevel) String() string {
	switch t.Type {
	case "entries":
		return fmt.Sprintf("%s. %s %s", t.Name, t.Eng_name, strings.Join(t.Entries, " "))
	case "":
		return ""
	default:
		return t.Type
	}
}

func (t spell_classes) String() string {
	dataLen := len(t.Fromclasslist)
	classlists := make([]string, dataLen)
	for i := 0; i < dataLen; i++ {
		classlists[i] = fmt.Sprint(t.Fromclasslist[i])
	}

	dataLen = len(t.Fromsubclass)
	subclasslists := make([]string, dataLen)
	for i := 0; i < dataLen; i++ {
		subclasslists[i] = fmt.Sprint(t.Fromsubclass[i])
	}

	dataLen = len(t.Fromclasslistvariant)
	classlistvariantlists := make([]string, dataLen)
	for i := 0; i < dataLen; i++ {
		classlistvariantlists[i] = fmt.Sprint(t.Fromclasslistvariant[i])
	}
	return fmt.Sprintf("职业：%s\n子职业：%s\n可选/变体职业：%s",
		strings.Join(classlists, ", "),
		strings.Join(subclasslists, ", "),
		strings.Join(classlistvariantlists, ", "),
	)
}

func (t spell_classes_fromClassList) String() string {
	return Class2String(t.Name)
}

func (t spell_classes_fromSubclass) String() string {
	return fmt.Sprintf("%s(%s)", t.Subclass.Name, Class2String(t.Class.Name))
}

func (t spell_classes_fromClassListVariant) String() string {
	return t.Name
}

func (t spell_races) String() string {
	return t.Name
}

func (t spell_backgrounds) String() string {
	return t.Name
}

func (t spell_eldritchInvocations) String() string {
	return t.Name
}

func (t spell_otherSources) String() string {
	return fmt.Sprintf("%s, 第 %d 页", t.Source, t.Page)
}

func (e *spell_srd) UnmarshalJSON(data []byte) error {
	err := json.Unmarshal(data, &e.Bool_raw)
	if err != nil { //maybe str
		return json.Unmarshal(data, &(e.String_raw))
	}
	return nil
}

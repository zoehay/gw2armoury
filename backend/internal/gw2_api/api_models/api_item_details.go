package apimodels

type ApiItemDetails interface {
}

type ApiArmorDetails struct {
	Type                  string              `json:"type"`
	WeightClass           string              `json:"weight_class"`
	Defense               int                 `json:"defense"`
	InfusionSlots         *[]InfusionSlotType `json:"infusion_slots"`
	AttributeAdjustment   float64             `json:"attriubte_adjustment"`
	InfixUpgrade          *ApiInfixUpgrade    `json:"infix_upgrade"`
	SuffixItemId          *int                `json:"suffix_item_id"`
	SecondarySuffixItemId string              `json:"secondary_suffix_item_id"`
	StatChoices           *[]int              `json:"stat_choices"`
}

type BagDetails struct {
	BagItems string
}

type InfusionSlotType string

const (
	EnrichmentSlot InfusionSlotType = "Enrichment"
	InfusionSlot   InfusionSlotType = "Infusion"
)

type ApiInfusionSlot struct {
	Flags  []string
	ItemId *int
}

type ApiInfixUpgrade struct {
	Id         int             `json:"id"`
	Attributes []ApiAttributes `json:"attributes"`
}

type ApiAttributes struct {
	Attribute string `json:"attribute"`
	Modifier  int    `json:"modifier"`
}

type InfixAttribute string

const (
	AgonyResistance   InfixAttribute = "AgonyResistance"
	BoonDuration      InfixAttribute = "BoonDuration"
	ConditionDamage   InfixAttribute = "ConditionDamage"
	ConditionDuration InfixAttribute = "ConditionDuration"
	CritDamage        InfixAttribute = "CritDamage"
	Healing           InfixAttribute = "Healing"
	Power             InfixAttribute = "Power"
	Precision         InfixAttribute = "Precision"
	Toughness         InfixAttribute = "Toughness"
	Vitality          InfixAttribute = "Vitality"
)

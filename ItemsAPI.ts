// Items v2/items

type ItemType =
  | "Armor"
  | "Back"
  | "Bag"
  | "Consumable"
  | "Container"
  | "CraftingMaterial"
  | "Gathering"
  | "Gizmo"
  | "JadeTechModule"
  | "Key"
  | "MiniPet"
  | "PowerCore"
  | "Relic"
  | "Tool"
  | "Trait"
  | "Trinket"
  | "Trophy"
  | "UpgradeComponent"
  | "Weapon";

type Rarity =
  | "Junk"
  | "Basic"
  | "Fine"
  | "Masterwork"
  | "Rare"
  | "Exotic"
  | "Ascended"
  | "Legendary";

type ItemTypeDetails =
  | ArmourDetails
  | BackItemDetails
  | BagDetails
  | ConsumableDetails
  | ContainerDetails
  | GatheringDetails
  | GizmoDetails;

type ItemUpgradesToOrFrom = {
  upgrade: string;
  item_id: number;
};

interface Item {
  id: number;
  chat_link: string;
  name: string;
  icon: string;
  description?: string;
  type: ItemType;
  rarity: Rarity;
  level: number;
  vendor_value: number;
  default_skin?: string;
  flags?: string[];
  game_types?: string[];
  restrictions?: string[];
  upgrades_into?: ItemUpgradesToOrFrom[];
  upgrades_from?: ItemUpgradesToOrFrom[];
  details?: ItemTypeDetails;
}

interface ArmourDetails {
  type: string;
  weight_class: string;
  defense: number;
  infusion_slots: InfusionSlots[];
  //If no infusion slot is present, the value of the infusion_slots property is an empty array. For each present upgrade slot, a flag object as defined above is present.
  attribute_adjustment: number;
  infix_upgrade?: InfixUpgrade;
  suffix_item_id?: number;
  secondary_suffix_item?: string;
  stat_choices?: number[];
}

interface BackItemDetails {
  infusion_slots: InfusionSlots[];
  attribute_adjustment: number;
  infix_upgrade?: InfixUpgrade;
  suffix_item_id?: number;
  secondary_suffix_item_id?: string;
  stat_choices?: number[];
}

interface BagDetails {
  size: number;
  no_sell_or_sort: boolean;
}

interface ConsumableDetails {
  type: string;
  description?: string;
  duration_ms?: number;
  unlock_type?: string;
  color_id?: number;
  recipe_id?: number;
  extra_recipe_ids?: number[];
  guild_upgrade_id?: number;
  apply_count?: number;
  name?: string;
  icon?: string;
  skins?: number[];
}

interface ContainerDetails {
  type: string;
}

interface GatheringDetails {
  type: string;
}

interface GizmoDetails {
  type: string;
  guild_upgrade_id?: number;
  vendor_ids?: number[];
}

interface MiniatureDetails {
  minipet_id: number;
}

interface SavageKitDetails {
  type: string;
  charges: number;
}

interface TrinketDetails {
  type: string;
  infusion_slots: InfusionSlots[];
  attribute_adjustment: number;
  infix_upgrade?: InfixUpgrade;
  suffix_item_id?: number;
  secondary_suffix_item_id?: string;
  stat_choices: number[];
}

interface UpgradeComponentDetails {
  type: string;
  flags: string[];
  infusion_upgrade_flags: string[];
  suffix: string;
  infix_upgrade: InfixUpgrade;
  bonuses?: string[];
}

interface WeaponDetails {
  type: string;
  damage_type: string;
  min_power: number;
  max_power: number;
  defense: number;
  infusion_slots: InfusionSlots[];
  attribute_adjustment: number;
  infix_upgrade?: InfixUpgrade;
  suffix_item_id?: number;
  secondary_suffix_item_id?: string;
  stat_choices?: number[];
}

// Subobjects
interface InfixUpgrade {
  id: number;
  attributes: BonusAttribute[];
  buff?: Buff;
}

interface InfusionSlots {
  flags: string[];
  item_id?: number;
}

// Sub Subobjects
interface BonusAttribute {
  attribute: string;
  modifier: number;
}

interface Buff {
  skill_id: number;
  description?: string;
}

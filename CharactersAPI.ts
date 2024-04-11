// vt/characters/

interface Character {
  name: string;
  race: string;
  gender: string;
  profession: string;
  level: number;
  equipment: Equipment[];
  bags: Bag[];
}

interface Equipment {
  id: number;
  slot?: string;
  infusions?: number[];
  upgrades?: number[];
  skin?: number;
  stats?: ItemStats;
  binding?: string;
  location?: string;
  tabs?: number[];
  charges?: number;
  bound_to?: string;
  dyes?: number[];
}

interface Bag {
  id: number;
  size: number;
  inventory: BagSlot[];
}

type BagSlot = BagItem | null;

interface BagItem {
  id: number;
  count: number;
  charges?: number;
  infusions?: number[];
  upgrades?: number[];
  skin?: number;
  stats?: ItemStats;
  dyes?: number[];
  binding?: string;
  bound_to?: string;
}

interface ItemStats {
  id: number;
  attributes: Attributes;
}

interface Attributes {
  [attributeName: string]: number;
}

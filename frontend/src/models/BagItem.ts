interface BagItem {
  AccountID: string;
  CharacterName: string;
  BagItemID: number;
  Icon: string;
  Count: number;
  Charges?: number;
  Infusions?: number[];
  Upgrades?: number[];
  Skin?: number;
  Stats?: { [key: string]: unknown };
  Dyes?: number[];
  Binding?: string;
  BoundTo?: string;
}

export default BagItem;

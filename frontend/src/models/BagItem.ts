interface BagItem {
  character_name: string; 
  id: number;
  icon: string;
  count: number;
  charges?: number;
  infusions?: number[];
  upgrades?: number[];
  skin?: number;
  stats?: { [key: string]: unknown };
  dyes?: number[];
  binding?: string;
  boundTo?: string;
}

export default BagItem;

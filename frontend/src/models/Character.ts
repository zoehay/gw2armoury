import { BagItem, APIBagItem, APIBagItemToBagItem } from "./BagItem";

export interface Character {
  name: string;
  equipment?: BagItem[];
  inventory?: BagItem[];
}

export interface APICharacter {
  name: string;
  equipment?: APIBagItem[];
  inventory?: APIBagItem[];
}

export function APICharacterToCharacter(apiCharacter: APICharacter): Character {
  return {
    name: apiCharacter.name,
    equipment: apiCharacter.equipment?.map(APIBagItemToBagItem),
    inventory: apiCharacter.inventory?.map(APIBagItemToBagItem),
  };
}

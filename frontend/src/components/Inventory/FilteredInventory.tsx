import { useState } from "react";
import { BagItem } from "../../models/BagItem";
import InventoryGroup from "./InventoryGroup";
import inventory from "./inventory.module.css";

interface FilteredInventoryProps {
  bagItems?: BagItem[];
}

type CharacterRecord = Record<string, BagItem[]>;

const FilteredInventory: React.FC<FilteredInventoryProps> = ({ bagItems }) => {
  let [searchTerm, setSearchTerm] = useState<string>("");

  const handleChange = (e: { target: { value: string } }) => {
    setSearchTerm(e.target.value);
  };

  let characterRecord: CharacterRecord = {}; // <character name, bag contents>

  bagItems?.forEach((bagItem) => {
    if (
      searchTerm == "" ||
      bagItem.name?.toLowerCase().includes(searchTerm.toLowerCase()) ||
      bagItem.count.toString().includes(searchTerm)
    ) {
      let name = bagItem.characterName;
      if (characterRecord[name]) {
        characterRecord[name].push(bagItem);
      } else {
        characterRecord[name] = [bagItem];
      }
    }
  });

  let characters: JSX.Element[] = [];

  if (characterRecord) {
    for (let key in characterRecord) {
      characters.push(
        <InventoryGroup
          characterName={key}
          contents={characterRecord[key]}
          key={key}
        ></InventoryGroup>
      );
    }
  }

  return (
    <>
      <input type="text" onChange={handleChange} />
      <div className={inventory.inventoryGroups}>{characters!}</div>
    </>
  );
};

export default FilteredInventory;

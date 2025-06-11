import React from "react";
import InventoryGroup from "./InventoryGroup";
import inventory from "./inventory.module.css";
import { BagItem } from "../../models/BagItem";
import { Character } from "../../models/Character";

interface AccountInventoryProps {
  sharedInventory?: BagItem[] | undefined;
  characters?: Character[];
}

const InventoryContainer: React.FC<AccountInventoryProps> = ({
  sharedInventory,
  characters,
}) => {
  // let [searchTerm, setSearchTerm] = useState<string>("");

  // const handleChange = (e: { target: { value: string } }) => {
  //   setTimeout(async () => {
  //     setSearchTerm(e.target.value);
  //   }, 1000);
  // };

  return (
    <>
      {/* <input type="text" onChange={handleChange} /> */}
      {sharedInventory && (
        <InventoryGroup
          characterName="Shared Inventory"
          characterInventory={sharedInventory}
        ></InventoryGroup>
      )}
      <div className={inventory.inventoryGroups}>
        {characters &&
          characters.map((character) => {
            return (
              <InventoryGroup
                key={character.name}
                characterName={character.name}
                characterInventory={character.inventory}
                equipment={character.equipment}
              ></InventoryGroup>
            );
          })}
      </div>
    </>
  );
};

export default InventoryContainer;

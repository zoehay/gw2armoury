import React from "react";
import InventoryGroup from "./InventoryGroup";
import inventory from "./inventory.module.css";
import { AccountInventory } from "../../models/AccountInventory";

interface AccountInventoryProps {
  accountInventory: AccountInventory;
}

const InventoryContainer: React.FC<AccountInventoryProps> = ({
  accountInventory,
}) => {
  // let [searchTerm, setSearchTerm] = useState<string>("");

  // const handleChange = (e: { target: { value: string } }) => {
  //   setTimeout(async () => {
  //     setSearchTerm(e.target.value);
  //   }, 1000);
  // };

  let characters: JSX.Element[] = [];

  if (accountInventory.characters) {
    for (let character of accountInventory.characters) {
      characters.push(
        <InventoryGroup
          characterName={character.name}
          contents={character.inventory}
        ></InventoryGroup>
      );
    }
  }

  return (
    <>
      {/* <input type="text" onChange={handleChange} /> */}
      {accountInventory.sharedInventory && (
        <InventoryGroup
          characterName="Shared Inventory"
          contents={accountInventory.sharedInventory}
        ></InventoryGroup>
      )}
      {accountInventory.characters && (
        <div className={inventory.inventoryGroups}>{characters!}</div>
      )}
    </>
  );
};

export default InventoryContainer;

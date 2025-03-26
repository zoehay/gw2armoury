import { useContext, useEffect, useState } from "react";
import { BagItem } from "../../models/BagItem";
import { ClientContext } from "../../util/ClientContext";
import content from "../content.module.css";
import CharacterInventory from "./CharacterInventory";
import inventory from "./inventory.module.css";

type CharacterInventories = Map<string, BagItem[]>;

const Inventory = () => {
  let context = useContext(ClientContext);
  let client = context;

  let [characterInventories, setCharacterInventories] =
    useState<CharacterInventories>(new Map());

  async function fetchData() {
    let items: BagItem[] = await client.getBagItems();
    setCharacterInventories(sortBagItems(items));
  }

  function sortBagItems(items: BagItem[]) {
    let inventories: CharacterInventories = new Map();
    for (let item of items) {
      let characterName = item.characterName;
      if (inventories.get(characterName)) {
        inventories.get(characterName)?.push(item);
      } else {
        inventories.set(characterName, [item]);
      }
    }
    return inventories;
  }

  useEffect(() => {
    fetchData();
  }, []);

  let characters: JSX.Element[] = [];

  if (characterInventories) {
    characterInventories.forEach((value: BagItem[], index: string) => {
      characters.push(
        <CharacterInventory
          characterName={index}
          contents={value}
        ></CharacterInventory>
      );
    });
  }

  return (
    <div className={content.page}>
      <div className={inventory.inventory}>{characters!}</div>
    </div>
  );
};

export default Inventory;

import React from "react";
import inventory from "./inventory.module.css";
import { BagItem } from "../../models/BagItem";
import { InventoryTile } from "./InventoryTile";

interface CharacterProps {
  characterName: string;
  contents?: BagItem[];
}

const CharacterInventory: React.FC<CharacterProps> = ({
  characterName,
  contents,
}) => {
  let itemTiles;
  if (contents) {
    itemTiles = contents.map((item, index) => (
      <InventoryTile bagItem={item} key={index} />
    ));
  }

  return (
    <div className={inventory.characterInventory}>
      <div className={inventory.characterName}>{characterName}</div>
      <div className={inventory.contents}>{itemTiles}</div>
    </div>
  );
};

export default CharacterInventory;

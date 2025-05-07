import React from "react";
import inventory from "./inventory.module.css";
import { BagItem } from "../../models/BagItem";
import { InventoryTile } from "./InventoryTile";

interface InventoryGroupProps {
  characterName: string;
  contents?: BagItem[];
}

const InventoryGroup: React.FC<InventoryGroupProps> = ({
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
    <div className={inventory.group}>
      <div className={inventory.name}>{characterName}</div>
      <div className={inventory.contents}>{itemTiles}</div>
    </div>
  );
};

export default InventoryGroup;

import BagItem from "../models/BagItem";
import React from "react";
import inventory from "./inventory.module.css";

export interface InventoryTileProps {
  bagItem: BagItem;
}

export const InventoryTile: React.FC<InventoryTileProps> = ({ bagItem }) => {
  return (
    <div className={inventory.tile}>
      <img className={inventory.icon} src={bagItem.icon} alt={bagItem.name} />
    </div>
  );
};

import BagItem from "../models/BagItem";
import React from "react";

export interface InventoryTileProps {
  bagItem: BagItem;
}

export const InventoryTile: React.FC<InventoryTileProps> = ({ bagItem }) => {
  return (
    <div>
      <p>AccountID ${bagItem.AccountID}</p>
    </div>
  );
};

import BagItem from "../models/BagItem";
import React, { useState } from "react";
import inventory from "./inventory.module.css";

export interface InventoryTileProps {
  bagItem: BagItem;
}

export const InventoryTile: React.FC<InventoryTileProps> = ({ bagItem }) => {
  let [displayDetails, setDisplayDetails] = useState(false);

  const handleMouseEnter = () => {
    setDisplayDetails(true);
  };

  const handleMouseLeave = () => {
    setDisplayDetails(false);
  };

  const handleTapToggle = () => {
    setDisplayDetails(!displayDetails);
  };

  return (
    <div
      className={inventory.tile}
      onMouseEnter={handleMouseEnter}
      onClick={handleTapToggle}
      onMouseLeave={handleMouseLeave}
      // onClickAway={handleTap}
    >
      {bagItem.count > 1 && (
        <div className={inventory.count}>{bagItem.count}</div>
      )}
      <img className={inventory.icon} src={bagItem.icon} alt={bagItem.name} />
      {displayDetails && (
        <div className={inventory.tooltip}>
          {bagItem.name}
          {bagItem.boundTo}
        </div>
      )}
    </div>
  );
};

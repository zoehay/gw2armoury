import React, { ReactNode } from "react";
import inventory from "./inventory.module.css";

interface InventoryGroupProps {
  children: ReactNode;
}

const InventoryGroup: React.FC<InventoryGroupProps> = ({ children }) => {
  return <div className={inventory.group}>{children}</div>;
};

export default InventoryGroup;

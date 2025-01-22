import React, { ReactNode } from "react";

interface InventoryGroupProps {
  children: ReactNode;
}

const InventoryGroup: React.FC<InventoryGroupProps> = ({ children }) => {
  return <div>{children}</div>;
};

export default InventoryGroup;

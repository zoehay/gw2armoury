import React from "react";
import { Account } from "../../models/Account";
import { KeyTile } from "./KeyTile";
import managekeys from "./managekeys.module.css";

interface KeyGroupProps {
  accounts?: Account[];
}

export const KeyGroup: React.FC<KeyGroupProps> = ({ accounts }) => {
  let keyTiles;
  if (accounts) {
    keyTiles = accounts.map((account, index) => (
      <KeyTile account={account} key={index} />
    ));
  }

  return (
    <div className={managekeys.keytiles}>
      <div>{keyTiles}</div>
    </div>
  );
};

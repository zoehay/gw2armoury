import React from "react";
import { Account } from "../../models/Account";
import managekeys from "./managekeys.module.css";

interface KeyTileProps {
  account: Account;
}

export const KeyTile: React.FC<KeyTileProps> = ({ account }) => {
  return (
    <div className={managekeys.keytile}>
      <p>{account.accountID}</p>
      <p>{account.accountName}</p>
      <p>{account.apiKey}</p>
    </div>
  );
};

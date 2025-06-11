import { useContext, useEffect, useState } from "react";
import { ClientContext } from "../../util/ClientContext";
import content from "../content.module.css";
import NoKeyPage from "../ErrorPage/NoKeyPage";
import { AccountInventory } from "../../models/AccountInventory";
import InventoryContainer from "./InventoryContainer";

const Inventory = () => {
  let context = useContext(ClientContext);
  let client = context;

  let [accountInventory, setAccountInventory] = useState<
    AccountInventory | undefined
  >();

  async function fetchData() {
    let inventory: AccountInventory = await client.getAccountInventory();
    setAccountInventory(inventory);
  }

  useEffect(() => {
    fetchData();
  }, []);

  return (
    <div className={content.page}>
      {accountInventory ? (
        <InventoryContainer
          accountInventory={accountInventory}
        ></InventoryContainer>
      ) : (
        <NoKeyPage />
      )}
    </div>
  );
};

export default Inventory;

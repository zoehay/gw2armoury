import { useContext, useEffect, useState } from "react";
import { BagItem } from "../../models/BagItem";
import { ClientContext } from "../../util/ClientContext";
import content from "../content.module.css";
import FilteredInventory from "./FilteredInventory";
import NoKeyPage from "../ErrorPage/NoKeyPage";

const Inventory = () => {
  let context = useContext(ClientContext);
  let client = context;

  let [bagItems, setBagItems] = useState<BagItem[] | undefined>([]);

  async function fetchData() {
    let items: BagItem[] = await client.getBagItems();
    setBagItems(items);
  }

  useEffect(() => {
    fetchData();
  }, []);

  return (
    <div className={content.page}>
      {bagItems ? (
        <FilteredInventory bagItems={bagItems}></FilteredInventory>
      ) : (
        <NoKeyPage />
      )}
    </div>
  );
};

export default Inventory;

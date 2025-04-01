import { useContext, useEffect, useState } from "react";
import { BagItem } from "../../models/BagItem";
import { ClientContext } from "../../util/ClientContext";
import content from "../content.module.css";
import FilteredInventory from "./FilteredInventory";

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
      <FilteredInventory bagItems={bagItems}></FilteredInventory>
    </div>
  );
};

export default Inventory;

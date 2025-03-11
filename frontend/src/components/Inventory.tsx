import { InventoryTile } from "./InventoryTile";
import InventoryGroup from "./InventoryGroup";
import { useContext, useEffect, useState } from "react";
import BagItem from "../models/BagItem";
import { ClientContext } from "../util/ClientContext";
import content from "./content.module.css";

const Inventory = () => {
  let [bagItems, setBagItems] = useState<BagItem[] | null>(null);
  let context = useContext(ClientContext);
  let client = context;
  let itemTiles: JSX.Element[];

  async function fetchData() {
    let items: BagItem[] = await client.getBagItems();
    setBagItems(items);
  }

  useEffect(() => {
    fetchData();
  }, []);

  if (bagItems) {
    itemTiles = bagItems.map((item, index) => (
      <InventoryTile bagItem={item} key={index} />
    ));
  }

  return (
    <div className={content.page}>
      <InventoryGroup>{itemTiles!}</InventoryGroup>
    </div>
  );
};

export default Inventory;

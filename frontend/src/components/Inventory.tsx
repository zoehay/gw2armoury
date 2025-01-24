import { InventoryTile } from "./InventoryTile";
import InventoryGroup from "./InventoryGroup";
import { useContext, useEffect, useState } from "react";
// import { ClientContext } from "../util/ClientContext";
import BagItem from "../models/BagItem";
import { ClientContext } from "../util/ClientContext";

// let item1: BagItem = {
//     AccountID: "accountidstring",
//     CharacterName: "Roman Meows",
//     BagItemID: 43772,
//     Icon: "",
//     Count: 7,
//     Binding: "Account"
// }


// let item2: BagItem = {
//   AccountID: "accountidstring",
//   CharacterName: "Roman Meows",
//   BagItemID: 89140,
//   Icon: "",
//   Count: 250,
// }

const Inventory = () => {
  let [bagItems, setBagItems] = useState<BagItem[] | null>(null);
  let context = useContext(ClientContext);
  let client = context;
  let itemTiles: JSX.Element[]
  
  async function fetchData() {
    let items: BagItem[] = await client.getBagItems();
    // let items = [item1, item2]
    setBagItems(items);
  }

  useEffect(() => {
    fetchData()
  }, [])

if (bagItems) {
  itemTiles= bagItems.map((item, index) => (
    <InventoryTile bagItem={item} key={index} />
  ));
}

  return (
    <div>
      <h2>Inventory</h2>
      <InventoryGroup>{itemTiles!}</InventoryGroup>
    </div>
  );
};

export default Inventory;

import { InventoryTile } from "./InventoryTile";
import InventoryGroup from "./InventoryGroup";
import { useContext } from "react";
// import { ClientContext } from "../util/ClientContext";
import BagItem from "../models/BagItem";

let item1: BagItem = {
    AccountID: "accountidstring",
    CharacterName: "Roman Meows",
    BagItemID: 43772,
    Icon: "",
    Count: 7,
    Binding: "Account"
}


let item2: BagItem = {
  AccountID: "accountidstring",
  CharacterName: "Roman Meows",
  BagItemID: 89140,
  Icon: "",
  Count: 250,
}



const Inventory = () => {
  // const context = useContext(ClientContext);
  // const client = context;
  // let items: BagItem[] = client.getBagItems();
  // console.log(typeof client);
  let items = [item1, item2]
  let itemTiles: JSX.Element[] = items.map((item, index) => (
    <InventoryTile bagItem={item} key={index} />
  ));

  return (
    <div>
      <h2>Inventory</h2>
      <InventoryGroup>{itemTiles}</InventoryGroup>
    </div>
  );
};

export default Inventory;

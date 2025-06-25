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
  let [searchTerm, setSearchTerm] = useState<string>("");

  const fetchData = async () => {
    let inventory: AccountInventory = await client.getAccountInventory();
    setAccountInventory(inventory);
  };

  useEffect(() => {
    fetchData();
  }, []);

  const handleChange = (e: React.FormEvent<HTMLInputElement>) => {
    const input = e.currentTarget.value;
    setSearchTerm(input);
  };

  const handleSubmit = async (e: React.SyntheticEvent) => {
    e.preventDefault();
    if (searchTerm) {
      const filteredInventory = await client.postInventorySearch(searchTerm);
      if (!filteredInventory) {
        console.log("Could not post search");
      } else {
        setAccountInventory(filteredInventory);
      }
    } else {
      fetchData();
    }
  };

  return (
    <div className={content.page}>
      {accountInventory ? (
        <>
          <form onSubmit={handleSubmit}>
            <label htmlFor="apikey-input">{`Search`}</label>
            <div>
              <input
                type="search"
                name="search-input"
                id="search"
                value={searchTerm}
                onChange={handleChange}
              />
            </div>
            <input type="submit" value="Submit" />
          </form>
          <InventoryContainer
            accountInventory={accountInventory}
          ></InventoryContainer>
        </>
      ) : (
        <NoKeyPage />
      )}
    </div>
  );
};

export default Inventory;

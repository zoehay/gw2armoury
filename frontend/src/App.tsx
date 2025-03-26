import "./App.css";
import ManageKeys from "./components/ManageKeys";
import { ClientProvider } from "./util/ClientContext";
import content from "./components/content.module.css";
import { Navbar } from "./components/Navbar/Navbar";
import Inventory from "./components/Inventory/Inventory";

function App() {
  return (
    <>
      <ClientProvider>
        <Navbar></Navbar>
        <div className={content.main}>
          <Inventory></Inventory>
          <ManageKeys></ManageKeys>
        </div>
      </ClientProvider>
    </>
  );
}

export default App;

import "./App.css";
import Inventory from "./components/Inventory";
import ManageKeys from "./components/ManageKeys";
import { ClientProvider } from "./util/ClientContext";
import content from "./components/content.module.css";

function App() {
  return (
    <>
      <ClientProvider>
        <div className={content.main}>
          <Inventory></Inventory>
          <ManageKeys></ManageKeys>
        </div>
      </ClientProvider>
    </>
  );
}

export default App;

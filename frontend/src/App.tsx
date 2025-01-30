import './App.css'
import Inventory from './components/Inventory'
import ManageKeys from './components/ManageKeys'
import { ClientProvider } from './util/ClientContext'

function App() {

  return (
    <>
    <ClientProvider>
    <div>Hello World</div>
    <Inventory></Inventory>
    <ManageKeys></ManageKeys>
    </ClientProvider>
    </>
  )
}

export default App

import './App.css'
import Inventory from './components/Inventory'
import { ClientProvider } from './util/ClientContext'

function App() {

  return (
    <>
    <ClientProvider>
    <div>Hello World</div>
    <Inventory></Inventory>
    </ClientProvider>
    </>
  )
}

export default App

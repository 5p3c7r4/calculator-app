import Calculator from "./components/Calculator";
import { Toaster } from 'sonner'
import "./App.css";
import "./components/Calculator.css";

function App() {
  return (
    <>
      <Calculator />
      <Toaster position="top-right" />
    </>
  );
}

export default App;

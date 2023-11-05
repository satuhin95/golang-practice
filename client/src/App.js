
import './App.css';
import { Route, Routes } from "react-router-dom";
import Home from "./pages/Home";
import ClientInfo from "./pages/ClientInfo";
import Header from './components/Header';

function App() {
  return (
    <>
    <Header/>
    <Routes>
      <Route path="/" element={<Home/>} />
      <Route path="/clientinfo/:id" element={<ClientInfo/>} />
    </Routes>
    </>
  );
}

export default App;

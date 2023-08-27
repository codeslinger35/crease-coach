import './App.css';
import Search from './Pages/Search';
import GoalieDetail from './Pages/GoalieDetail';
import { Route, Routes } from 'react-router-dom';

function App() {
  return (
    <>
      <Routes>
        <Route path='/' element={<Search />}></Route>
        <Route path='/goalie/:id' element={<GoalieDetail />}></Route>
      </Routes>
    </>
  );
}

export default App;

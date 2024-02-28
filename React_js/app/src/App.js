import './App.css';
import Border from './Border';
import Test from './Test';



function App() {
  console.log("app")
  return (
    <div className="App">
      <Border>
        <Test/>
      </Border>
    </div>
  );
}



const MyComp = () => {
  return(
    <>
      <h1>Saishwar</h1>
    </>
  )
}


export default App;

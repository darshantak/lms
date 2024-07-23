import ReactPlayer from 'react-player';
import './App.css';
function App() {
  return (
    <div className="App">
      <ReactPlayer 
        url = "https://www.youtube.com/watch?v=Mx92lTYxrJQ"
        volume = "0.4"
        
      />
    </div>
  );
}

export default App;

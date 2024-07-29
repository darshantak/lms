import { MantineProvider } from '@mantine/core';
import LoginPage from './LoginPage';
function App() {
  return (
    <div className="App">
      <MantineProvider>
        <LoginPage/>
      </MantineProvider>
    </div>
  );
}

export default App;

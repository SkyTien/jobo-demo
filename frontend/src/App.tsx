import './App.css';
import PatientList from './components/PatientList';

const App = () => {
  return (
    <div style={{
      display: 'flex',
      alignItems: 'center',
      justifyContent: 'center',
      paddingTop: 30
  }}>
    <PatientList/>
    </div>
  );
}

export default App;

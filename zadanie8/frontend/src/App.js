import { useState } from 'react';

function App() {
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');
  const [token, setToken] = useState(null);
  const [message, setMessage] = useState('');

  const login = async () => {
    const res = await fetch('http://localhost:8080/login', {
      method: 'POST',
      headers: {'Content-Type': 'application/json'},
      body: JSON.stringify({ username, password }),
    });

    const data = await res.json();
    if (res.ok) {
      setToken(data.token);
      setMessage('Zalogowano pomyślnie');
    } else {
      setMessage(data.error || 'Błąd logowania');
    }
  };

  const fetchProtected = async () => {
    const res = await fetch('http://localhost:8080/protected', {
      headers: {
        'Authorization': token,
      }
    });

    const data = await res.json();
    setMessage(data.message || 'Błąd');
  };

  return (
    <div className="p-4">
      <h1>Logowanie</h1>
      <input value={username} onChange={e => setUsername(e.target.value)} placeholder="Login" />
      <input type="password" value={password} onChange={e => setPassword(e.target.value)} placeholder="Hasło" />
      <button onClick={login}>Zaloguj</button>
      {token && <button onClick={fetchProtected}>Pobierz dane</button>}
      <p>{message}</p>
    </div>
  );
}

export default App;

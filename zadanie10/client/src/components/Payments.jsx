import { useState } from "react";
import axios from "axios";

export default function Payments() {
  const [productId, setProductId] = useState("");
  const [name, setName] = useState("");

  const handleSubmit = e => {
    e.preventDefault();
    axios.post("https://server10test2-hmdmhwd5h7hubjhw.polandcentral-01.azurewebsites.net/payment", {
      productId,
      name,
    }).then(() => alert("Płatność wysłana!"));
  };

  return (
    <div>
      <h2>Płatność</h2>
      <form onSubmit={handleSubmit}>
        <input
          type="text"
          placeholder="ID produktu"
          value={productId}
          onChange={e => setProductId(e.target.value)}
        />
        <input
          type="text"
          placeholder="Imię i nazwisko"
          value={name}
          onChange={e => setName(e.target.value)}
        />
        <button type="submit">Zapłać</button>
      </form>
    </div>
  );
}
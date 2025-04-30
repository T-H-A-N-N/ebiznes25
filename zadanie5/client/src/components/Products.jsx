import { useEffect, useState } from "react";
import axios from "axios";

export default function Products() {
  const [products, setProducts] = useState([]);

  useEffect(() => {
    axios.get("http://localhost:8080/products")
      .then(res => setProducts(res.data))
      .catch(err => console.error(err));
  }, []);

  return (
    <div>
      <h2>Produkty</h2>
      <ul>
        {products.map(p => (
          <li key={p.id}>{p.name} - {p.price} PLN</li>
        ))}
      </ul>
    </div>
  );
}
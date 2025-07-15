import React, { useState, useEffect } from 'react';
import './style.css';

const STORAGE_MARKS = [
  '0', '250GB', '500GB', '1TB', '2TB', '3TB', '4TB', '8TB', '12TB', '24TB', '48TB', '72TB'
];
const STORAGE_NUMS = [0, 250, 500, 1024, 2048, 3072, 4096, 8192, 12288, 24576, 49152, 73728];
const RAM_OPTIONS = ['2GB', '4GB', '8GB', '12GB', '16GB', '24GB', '32GB', '48GB', '64GB', '96GB'];
const HDD_TYPES = ['SAS', 'SATA', 'SSD'];
const LOCATION_LIST = [
  'AmsterdamAMS-01', 'DallasDAL-10', 'FrankfurtFRA-10', 'Hong KongHKG-10',
  'LondonLON-01', 'San FranciscoSFO-12', 'SingaporeSIN-11', 'Washington D.C.WDC-01'
]; // Update as needed

export default function App() {
  const [storage, setStorage] = useState([0, STORAGE_NUMS.length - 1]);
  const [ram, setRam] = useState([]);
  const [hdd, setHdd] = useState('');
  const [location, setLocation] = useState('');
  const [servers, setServers] = useState([]);
  const [loading, setLoading] = useState(false);

  const handleRamChange = (e) => {
    const value = e.target.value;
    setRam(prev =>
      prev.includes(value) ? prev.filter(r => r !== value) : [...prev, value]
    );
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    setLoading(true);
    // Prepare filter payload
    const payload = {};
    if (ram.length > 0) payload.ram = ram.join(',');
    if (hdd) payload.hdd = hdd;
    if (location) payload.location = location;
    if (storage[0] > 0 || storage[1] < STORAGE_NUMS.length - 1) {
      payload.storage = `${STORAGE_MARKS[storage[0]]}-${STORAGE_MARKS[storage[1]]}`;
    }
    try {
      const res = await fetch('/servers/filter', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(payload),
      });
      const data = await res.json();
      setServers(data.data || []);
    } catch (err) {
      setServers([]);
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="container">
      <h1>Leaseweb Server Finder</h1>
      <form onSubmit={handleSubmit} className="filter-form">
        <div className="form-group">
          <label>Storage (Range):</label>
          <div className="slider-group">
            <input
              type="range"
              min={0}
              max={STORAGE_NUMS.length - 1}
              value={storage[0]}
              onChange={e => setStorage([+e.target.value, storage[1]])}
              step={1}
              list="storagemarks"
            />
            <input
              type="range"
              min={0}
              max={STORAGE_NUMS.length - 1}
              value={storage[1]}
              onChange={e => setStorage([storage[0], +e.target.value])}
              step={1}
              list="storagemarks"
            />
            <datalist id="storagemarks">
              {STORAGE_MARKS.map((mark, idx) => (
                <option value={idx} label={mark} key={mark} />
              ))}
            </datalist>
            <div>
              <span>{STORAGE_MARKS[storage[0]]}</span> - <span>{STORAGE_MARKS[storage[1]]}</span>
            </div>
          </div>
        </div>
        <div className="form-group">
          <label>RAM:</label>
          <div className="checkbox-group">
            {RAM_OPTIONS.map(opt => (
              <label key={opt}>
                <input
                  type="checkbox"
                  value={opt}
                  checked={ram.includes(opt)}
                  onChange={handleRamChange}
                />
                {opt}
              </label>
            ))}
          </div>
        </div>
        <div className="form-group">
          <label>Harddisk type:</label>
          <select value={hdd} onChange={e => setHdd(e.target.value)}>
            <option value="">-- Select --</option>
            {HDD_TYPES.map(type => (
              <option value={type} key={type}>{type}</option>
            ))}
          </select>
        </div>
        <div className="form-group">
          <label>Location:</label>
          <select value={location} onChange={e => setLocation(e.target.value)}>
            <option value="">-- Select --</option>
            {LOCATION_LIST.map(loc => (
              <option value={loc} key={loc}>{loc}</option>
            ))}
          </select>
        </div>
        <button type="submit" disabled={loading}>Search</button>
      </form>
      <div className="results">
        {loading && <p>Loading...</p>}
        {!loading && servers.length > 0 && (
          <table>
            <thead>
              <tr>
                <th>ID</th>
                <th>Model</th>
                <th>RAM</th>
                <th>HDD</th>
                <th>Location</th>
                <th>Price</th>
              </tr>
            </thead>
            <tbody>
              {servers.map(server => (
                <tr key={server.id}>
                  <td>{server.id}</td>
                  <td>{server.model}</td>
                  <td>{server.ram}</td>
                  <td>{server.hdd}</td>
                  <td>{server.location}</td>
                  <td>{server.price}</td>
                </tr>
              ))}
            </tbody>
          </table>
        )}
        {!loading && servers.length === 0 && <p>No servers found.</p>}
      </div>
    </div>
  );
}

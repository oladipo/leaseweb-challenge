# Leaseweb Challenge Frontend

This is a React-based frontend for the Leaseweb Challenge API. It allows users to filter and search for servers using a web form with the following criteria:

- **Storage:** Range slider (0, 250GB, 500GB, 1TB, ... 72TB)
- **RAM:** Checkboxes (2GB, 4GB, ... 96GB)
- **Harddisk type:** Dropdown (SAS, SATA, SSD)
- **Location:** Dropdown (see list)

## Getting Started

### Prerequisites
- Node.js 18+
- npm or yarn

### Installation
```bash
cd frontend
npm install
```

### Running the App (Development)
```bash
npm run dev
```
The app will be available at [http://localhost:5173](http://localhost:5173) and will proxy API requests to the backend at `http://localhost:8080`.

### Build for Production
```bash
npm run build
```

## Usage
- Fill in the form to filter servers.
- Results will be displayed in a table below the form.

## API Backend
Make sure the Leaseweb Challenge API is running locally on port 8080, or update the proxy in `vite.config.js` accordingly.

---

## Folder Structure
- `src/App.jsx`: Main app and form logic
- `src/style.css`: Styles
- `vite.config.js`: Vite config with proxy

---

## License
MIT

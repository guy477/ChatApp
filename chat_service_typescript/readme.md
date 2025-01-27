# Setup

## Install NVM

```sh
curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.39.7/install.sh | bash
nvm --version
```

```sh
nvm install 20
nvm use 20
```

## Run Frontend

```sh
cd frontend
rm -rf node_modules package-lock.json
npm install vite@latest
npm install
npm run dev
```


### (REFERENCE) Create NEW Frontend

For reference, the projects current frontend was built using the Vite template below.

```sh
npm create vite@latest "frontend" -- --template react-ts
cd frontend
npm install
```
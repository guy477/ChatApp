# Setup

## Create Environment
```sh
conda create -n chat_service python=3.11
conda activate chat_service
```

## Install Dependencies
```sh 
pip install fastapi uvicorn aiosqlite
```

## Run Service
```sh
python python_service.py
```


# Future

- [ ] Add auth layer
- [ ] Revisit database design to handle more datatypes
    - [ ] embeddings
    - [ ] images
    - [ ] audio(?)
    - [ ] etc?
- [ ] Migrate to a remote database solution
- [ ] Replace aiosqlite
    - [ ] asyncpg; aiomysql; aioodbc; etc.
    - [ ] Persistent connections
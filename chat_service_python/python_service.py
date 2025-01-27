import uvicorn
from fastapi import FastAPI
from routes.routes import router
from _utils._utils import get_db_connection, logger

app = FastAPI()

@app.on_event("startup")
async def startup_event():
    db = await get_db_connection()
    await db.execute("""
        CREATE TABLE IF NOT EXISTS users (
            user_id TEXT PRIMARY KEY,
            created_at DATETIME DEFAULT CURRENT_TIMESTAMP
        );
    """)
    await db.execute("""
        CREATE TABLE IF NOT EXISTS chats (
            chat_id TEXT PRIMARY KEY,
            user_id TEXT NOT NULL,
            created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
            FOREIGN KEY(user_id) REFERENCES users(user_id)
        );
    """)
    await db.execute("""
        CREATE TABLE IF NOT EXISTS messages (
            message_id TEXT PRIMARY KEY,
            chat_id TEXT NOT NULL,
            role TEXT NOT NULL,
            content TEXT NOT NULL,
            created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
            FOREIGN KEY(chat_id) REFERENCES chats(chat_id)
        );
    """)
    await db.commit()
    logger.info("Database tables ensured.")

@app.on_event("shutdown")
async def shutdown_event():
    db = await get_db_connection()
    await db.close()
    logger.info("Database connection closed.")

# Include API routes
app.include_router(router)

if __name__ == "__main__":
    uvicorn.run("python_service:app", host="0.0.0.0", port=8000, reload=True) 
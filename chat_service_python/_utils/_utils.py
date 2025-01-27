import aiosqlite
import logging

DB_PATH = "chat_app.db"

# Initialize logger
logger = logging.getLogger("chat_service")
logger.setLevel(logging.INFO)
handler = logging.StreamHandler()
formatter = logging.Formatter(
    '%(asctime)s - %(name)s - %(levelname)s - %(message)s'
)
handler.setFormatter(formatter)
logger.addHandler(handler)

# Asynchronous function to get the database connection
async def get_db_connection():
    if not hasattr(get_db_connection, "connection"):
        get_db_connection.connection = await aiosqlite.connect(DB_PATH)
        get_db_connection.connection.row_factory = aiosqlite.Row
    return get_db_connection.connection 
import uuid
from datetime import datetime
from _utils._utils import get_db_connection, logger

class UserModel:
    @staticmethod
    async def create_user(user_id: str):
        db = await get_db_connection()
        async with db.execute("SELECT user_id FROM users WHERE user_id = ?", (user_id,)) as cursor:
            row = await cursor.fetchone()
            if not row:
                await db.execute("INSERT INTO users(user_id, created_at) VALUES (?, ?)", (user_id, datetime.utcnow()))
                await db.commit()
                logger.info(f"Created user with ID: {user_id}")
        return {"user_id": user_id}

class ChatModel:
    @staticmethod
    async def create_chat(user_id: str):
        chat_id = str(uuid.uuid4())
        db = await get_db_connection()
        await db.execute(
            "INSERT INTO chats(chat_id, user_id, created_at) VALUES (?, ?, ?)",
            (chat_id, user_id, datetime.utcnow())
        )
        await db.commit()
        logger.info(f"Created chat with ID: {chat_id} for user: {user_id}")
        return {"chat_id": chat_id}

    @staticmethod
    async def get_messages(chat_id: str):
        db = await get_db_connection()
        async with db.execute(
            "SELECT role, content FROM messages WHERE chat_id=? ORDER BY created_at ASC",
            (chat_id,)
        ) as cursor:
            rows = await cursor.fetchall()
            messages = [{"role": row["role"], "content": row["content"]} for row in rows]
            logger.info(f"Fetched messages for chat ID: {chat_id}")
            return messages

class MessageModel:
    @staticmethod
    async def add_message(chat_id: str, role: str, content: str):
        message_id = str(uuid.uuid4())
        db = await get_db_connection()
        await db.execute(
            "INSERT INTO messages(message_id, chat_id, role, content, created_at) VALUES(?, ?, ?, ?, ?)",
            (message_id, chat_id, role, content, datetime.utcnow())
        )
        await db.commit()
        logger.info(f"Added message ID: {message_id} to chat ID: {chat_id}")
        return {"message_id": message_id} 
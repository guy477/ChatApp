from models.models import UserModel, ChatModel, MessageModel

class UserController:
    @staticmethod
    async def create_user(user_id: str):
        return await UserModel.create_user(user_id)

class ChatController:
    @staticmethod
    async def create_chat(user_id: str):
        return await ChatModel.create_chat(user_id)

    @staticmethod
    async def get_chat_messages(chat_id: str):
        return await ChatModel.get_messages(chat_id)

class MessageController:
    @staticmethod
    async def add_message(chat_id: str, role: str, content: str):
        return await MessageModel.add_message(chat_id, role, content) 
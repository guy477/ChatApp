from fastapi import APIRouter, HTTPException
from pydantic import BaseModel
from controllers.controllers import UserController, ChatController, MessageController

router = APIRouter()

class MessageIn(BaseModel):
    role: str
    content: str

@router.post("/users/{user_id}")
async def create_user_if_not_exists(user_id: str):
    result = await UserController.create_user(user_id)
    return result

@router.post("/chats/")
async def create_chat(user_id: str):
    result = await ChatController.create_chat(user_id)
    return result

@router.get("/chats/{chat_id}/messages")
async def get_chat_messages(chat_id: str):
    messages = await ChatController.get_chat_messages(chat_id)
    if messages is None:
        raise HTTPException(status_code=404, detail="Chat not found")
    return messages

@router.post("/chats/{chat_id}/messages")
async def add_message(chat_id: str, msg: MessageIn):
    result = await MessageController.add_message(chat_id, msg.role, msg.content)
    return result 
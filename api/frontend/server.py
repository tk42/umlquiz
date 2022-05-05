# -*- coding: utf-8 -*-
import os
import ast
import zmq
import datetime
import logging
from enum import Enum
from datetime import datetime

from pydantic import BaseModel
from typing import Optional #List
from fastapi import FastAPI, Response, HTTPException, Query, Form, Depends, status
from fastapi.security import OAuth2PasswordBearer

is_prod = bool(ast.literal_eval(os.environ.get("IS_PROD", "False")))

AUTH_USERNAME = os.environ.get("AUTH_USERNAME")
assert AUTH_USERNAME, "AUTH_USERNAME is empty"

AUTH_PASSWORD = os.environ.get("AUTH_PASSWORD")
assert AUTH_PASSWORD, "AUTH_PASSWORD is empty"

app = FastAPI(
    title="UMLQUIZ API",
    description="DM API FOR UMLQUIZ.COM",
    version="1.0.0",
    redoc_url=None if is_prod else "/redoc",
    docs_url=None if is_prod else "/explorer",
    openapi_url=None if is_prod else "/openapi.json"
)

oauth2_scheme = OAuth2PasswordBearer(tokenUrl="token")


class RawResponse(Response):
    media_type = "application/json"

    def render(self, content: bytes) -> bytes:
        return content



Membership = Enum(
    "Membership",
    {
        "Bronze": "bronze",
        "Silver": "silver",
        "Gold": "gold",
    },
)

DiagramType = Enum(
    "DiagramType",
    {
        "Class": "class",
        "Sequence": "sequence",
    },
)


QuizStatus = Enum(
    "QuizStatus",
    {
        "Draft": "draft",
        "Reviewing": "reviewing",
        "Public": "public",
    }
)


# UUID = quiz_id + language
class Quiz(BaseModel):
    quiz_id: Optional[str]
    language: str
    diagram_type: DiagramType
    level: str
    text: str
    diagram: str
    likes: Optional[int]
    author_id: Optional[str]
    created_at: Optional[int]
    updated_at: Optional[int]
    status: Optional[QuizStatus]


class User(BaseModel):
    username: str
    password: str
    email: str
    user_id: str
    created_at: int
    updated_at: int
    profile: str
    membership: Membership
    liked_quiz_ids: str
    quiz_history: str


context = zmq.Context()
socket = context.socket(zmq.REQ)
socket.connect("tcp://localhost:5556")

####
##  Authorize
####
from fastapi.param_functions import Depends
from fastapi.security.oauth2 import OAuth2PasswordRequestForm

@app.post('/token')
def get_token(request: OAuth2PasswordRequestForm = Depends()):
    socket.send_json(
        {
            "group": "token",
            "method": "get",
            "params": {
                "username": request.username,
                "password": request.password,
            },
        }
    )
    return RawResponse(content=socket.recv())

####
##  User
####

@app.post("/user", response_model=User)
async def post_user(
        username: str = Form(...),
        password: str = Form(...),
        email: str = Form(...),
        token: str = Depends(oauth2_scheme)
    ):
    socket.send_json(
        {
            "group": "user",
            "method": "post",
            "params": {
                "username": username,
                "password": password,
                "email": email,
            },
        }
    )
    return RawResponse(content=socket.recv())

@app.get("/user/{user_id}")
async def get_user(
        user_id: str,
        token: str = Depends(oauth2_scheme)
    ):
    socket.send_json(
        {
            "group": "user",
            "method": "get",
            "params": {
                "user_id": user_id,
            },
        }
    )
    return RawResponse(content=socket.recv())

@app.put("/user/{user_id}")
async def put_user(
        user_id: str, 
        user: User = Form(...),
        token: str = Depends(oauth2_scheme)
    ):
    d = {k: v.value if type(v) in [Membership] else v for k, v in user.dict().items()}
    socket.send_json(
        {
            "group": "user",
            "method": "put",
            "params": {
                "user_id": user_id,
                "user": d,
            },
        }
    )
    return RawResponse(content=socket.recv())


@app.delete("/user/{user_id}")
async def delete_user(
        user_id: str,
        token: str = Depends(oauth2_scheme)
    ):
    socket.send_json(
        {
            "group": "user",
            "method": "delete",
            "params": {
                "user_id": user_id,
            },
        }
    )
    return RawResponse(content=socket.recv())


####
##  quiz
####

@app.get("/quiz")
async def get_quiz(
        token: str = Depends(oauth2_scheme)
    ):
    socket.send_json(
        {
            "group": "quiz",
            "method": "get",
        }
    )
    return RawResponse(content=socket.recv())


@app.post("/quiz/{quiz_id}")
async def post_quiz(
        quiz_id: str,
        language: str = Form(...),
        quiz: Quiz = Form(...),
        token: str = Depends(oauth2_scheme)
    ):
    d = {k: v.value if type(v) in [Diagram, QuizStatus] else v for k, v in quiz.dict().items()}
    socket.send_json(
        {
            "group": "quiz",
            "method": "post",
            "params": {
                "quiz_id": quiz_id,
                "language": language,
                "quiz": d,
            },
        }
    )
    return RawResponse(content=socket.recv())


@app.put("/quiz/{quiz_id}")
async def put_quiz(
        quiz_id: str,
        quiz: Quiz = Form(...),
        token: str = Depends(oauth2_scheme)
    ):
    d = {k: v.value if type(v) in [Diagram, QuizStatus] else v for k, v in quiz.dict().items()}
    socket.send_json(
        {
            "group": "quiz",
            "method": "put",
            "params": {
                "quiz_id": quiz_id,
                "quiz": d,
            },
        }
    )
    return RawResponse(content=socket.recv())


####
##  quiz by User
####


@app.get("/{user_id}/quiz")
async def get_quizs_by_user(
        user_id: str,
        token: str = Depends(oauth2_scheme)
    ):
    socket.send_json(
        {
            "group": "quizs_by_user",
            "method": "get",
            "params": {
                "user_id": user_id,
            },
        }
    )
    return RawResponse(content=socket.recv())


@app.get("/{user_id}/quiz/{quiz_id}")
async def get_quiz_by_user(
        user_id: str,
        quiz_id: str,
        token: str = Depends(oauth2_scheme)
    ):
    socket.send_json(
        {
            "group": "quiz_by_user",
            "method": "get",
            "params": {
                "user_id": user_id,
                "quiz_id": quiz_id,
            },
        }
    )
    return RawResponse(content=socket.recv())



@app.post("/{user_id}/quiz")
async def post_quiz_by_user(
        user_id: str,
        quiz: Quiz = Form(...),
        token: str = Depends(oauth2_scheme)
    ):
    socket.send_json(
        {
            "group": "quiz_by_user",
            "method": "post",
            "params": {
                "user_id": user_id,
                "quiz": quiz,
            },
        }
    )
    return RawResponse(content=socket.recv())



@app.put("/{user_id}/quiz/{quiz_id}")
async def put_quiz_by_user(
        user_id: str,
        quiz_id: str,
        quiz: Quiz = Form(...),
        token: str = Depends(oauth2_scheme)
    ):
    socket.send_json(
        {
            "group": "quiz_by_user",
            "method": "put",
            "params": {
                "user_id": user_id,
                "quiz_id": quiz_id,
                "quiz": quiz,
            },
        }
    )
    return RawResponse(content=socket.recv())


@app.delete("/{user_id}/quiz/{quiz_id}")
async def delete_quiz_by_user(
        user_id: str,
        quiz_id: str,
        token: str = Depends(oauth2_scheme)
    ):
    socket.send_json(
        {
            "group": "quiz_by_user",
            "method": "delete",
            "params": {
                "user_id": user_id,
                "quiz_id": quiz_id,
            },
        }
    )
    return RawResponse(content=socket.recv())

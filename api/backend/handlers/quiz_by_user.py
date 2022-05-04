import uuid
from fastapi import Form
from . import session_scope, Quiz


def post(
        user_id: str,
        quiz: dict,
    ):
    now = int(time.time())
    q = Quiz(
        quiz_id = str(uuid.uuid4()),
        language = quiz["language"],
        diagram_type = quiz["diagram_type"],
        level = quiz["level"],
        text = quiz["text"],
        diagram = quiz["diagram"],
        likes = 0,
        author_id = user_id,
        created_at = now,
        updated_at = now,
        status = "draft",
    )
    with session_scope() as s:
        s.add(q)
        return q.to_dict()

def get(
        user_id: str,
        quiz_id: str,
    ):
    with session_scope() as s:
        q = s.query(Quiz).filter(Quiz.quiz_id==quiz_id and Quiz.author_id==user_id).first()
        return q.to_dict()


def put(
        user_id: str,
        quiz_id: str,
        quiz: dict,
    ):
    with session_scope() as s:
        q = s.query(Quiz).filter(Quiz.quiz_id==quiz_id and Quiz.author_id==user_id).first()
        for k, v in quiz.items():
            setattr(q, k, v)
    return quiz


def delete(
        user_id: str,
        quiz_id: str,
    ):
    with session_scope() as s:
        q = s.query(Quiz).filter(Quiz.quiz_id==quiz_id and Quiz.author_id==user_id).first()
        s.delete(q)
        return q.to_dict()

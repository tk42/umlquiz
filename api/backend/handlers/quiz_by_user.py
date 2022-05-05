import uuid
from fastapi import Form
from . import session_scope, Quiz, QuizText


def post(
        user_id: str,
        **kwargs,
    ):
    quiz_id = str(uuid.uuid4())
    now = int(time.time())
    q = Quiz(
        quiz_id = quiz_id,
        diagram_type = kwargs["diagram_type"],
        level = kwargs["level"],
        diagram = kwargs["diagram"],
        likes = 0,
        author_id = user_id,
        created_at = now,
        updated_at = now,
        status = "draft",
    )
    qt = QuizText(
        quiz_id = quiz_id,
        language = kwargs["language"],
        title = kwargs["title"],
        text = kwargs["text"],
    )
    with session_scope() as s:
        s.add(q)
        s.add(qt)
        return q.to_dict() | qt.to_dict()

def get(
        language: str,
        quiz_id: str,
    ):
    with session_scope() as s:
        q = s.query(Quiz).filter(Quiz.quiz_id==quiz_id).first()
        qt = s.query(QuizText).filter(QuizText.quiz_id==quiz_id and Quiz.language==language).first()
        return q.to_dict() | qt.to_dict()


def put(
        language: str,
        user_id: str,
        quiz_id: str,
        **kwargs
    ):
    with session_scope() as s:
        q = s.query(Quiz).filter(Quiz.quiz_id==quiz_id and Quiz.author_id==user_id).first()
        qt = s.query(QuizText).filter(QuizText.quiz_id==quiz_id and QuizText.language==language).first()
        for k, v in kwargs.items():
            if v:
                if k in ['text', 'title']:
                    setattr(qt, k, v)
                else:
                    setattr(q, k, v)
        return q.to_dict() | qt.to_dict()


def delete(
        user_id: str,
        quiz_id: str,
    ):
    with session_scope() as s:
        q = s.query(Quiz).filter(Quiz.quiz_id==quiz_id and Quiz.author_id==user_id).first()
        s.delete(q)
        return q.to_dict()

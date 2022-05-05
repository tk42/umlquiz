from fastapi import Form
from . import session_scope, ReportQuiz


def get(
        user_id: str
    ):
    reports = []
    with session_scope() as s:
        quizs = s.query(Quiz).filter(Quiz.author_id==user_id).all()
        for q in quizs:
            reports += [s.query(ReportQuiz).filter(ReportQuiz.quiz_id==q.quiz_id and ReportQuiz.language==q.language).all()]
        return [report.to_dict() for report in reports]

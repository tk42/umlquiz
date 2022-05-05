import os
import hashlib

from dotenv import load_dotenv
from sqlalchemy import create_engine
from sqlalchemy.orm import Session
from sqlalchemy.ext.declarative import declarative_base

for key in ['USERNAME', 'PASSWORD', 'HOST', 'DATABASE']:
    assert os.getenv(key), f'{key} is empty'

engine = create_engine(
    f"mysql+pymysql://{os.getenv('USERNAME')}:{os.getenv('PASSWORD')}@{os.getenv('HOST')}/{os.getenv('DATABASE')}?charset=utf8mb4",
    connect_args={
        "ssl": {
            "ssl_mode": "VERIFY_IDENTITY",
            "ssl_ca": "/etc/ssl/certs/ca-certificates.crt",
        }
    }, pool_timeout=20, pool_recycle=299
)
Base = declarative_base(bind=engine)

from contextlib import contextmanager

@contextmanager
def session_scope():
    session = Session()  # def __enter__
    try:
        yield session  # with asでsessionを渡す
        session.commit()  # 何も起こらなければcommit()
    except:
        session.rollback()  # errorが起こればrollback()
        raise
    finally:
        session.close()  # どちらにせよ最終的にはclose()


from sqlalchemy.schema import Column
from sqlalchemy.types import Integer, String, TEXT

class User(Base):
    __tablename__ = "user"  # テーブル名を指定
    user_id = Column(String(255), primary_key=True)
    username = Column(String(255))
    password = Column(String(255))
    email = Column(String(255))
    created_at = Column(Integer)
    updated_at = Column(Integer)
    membership = Column(String(8))
    profile = Column(TEXT)
    liked_quiz_ids = Column(TEXT)
    quiz_history = Column(TEXT)

    def get_liked_quiz_ids(self):
        return self.liked_quiz_ids.split(',')

    def get_quiz_history(self):
        return self.quiz_history.split(',')

    def to_dict(self):
        return {k: str(v) for k, v in self.__dict__.items() if k != "_sa_instance_state"}


class Quiz(Base):
    __tablename__ = "quiz"  # テーブル名を指定
    quiz_id = Column(String(255), primary_key=True)
    language = Column(String(8), primary_key=True)
    diagram_type = Column(String(16))
    level = Column(String(16))
    text = Column(TEXT)
    diagram = Column(TEXT)
    likes = Column(Integer)
    author_id = Column(String(255))
    created_at = Column(Integer)
    updated_at = Column(Integer)
    status = Column(String(8))

    def uuid(self):
        return hashlib.sha256((f"{self.quiz_id}:{self.language}").encode()).hexdigest()

    def to_dict(self):
        return {k: v for k, v in self.__dict__.items() if k != "_sa_instance_state"}


class ReportQuiz(Base):
    __tablename__ = "report_quiz"  # テーブル名を指定
    quiz_id = Column(String(255), primary_key=True)
    language = Column(String(8), primary_key=True)
    created_at = Column(Integer, primary_key=True)
    text = Column(TEXT)
    diagram = Column(TEXT)

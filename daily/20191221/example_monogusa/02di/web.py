# this module is generated by monogusa.web.codegen
import commands
from fastapi import (
    APIRouter,
    Depends,
    FastAPI
)
import typing as t
from pydantic import BaseModel
from monogusa.web import (
    runtime
)


def db(database_url: str=Depends(commands.database_url)) -> commands.DB:
    return commands.db(database_url)


router = APIRouter()


class HelloInput(BaseModel):
    name: str  = 'world'


@router.post("/hello", response_model=runtime.CommandOutput)
def hello(input: HelloInput, db: commands.DB=Depends(db)) -> t.Dict[str, t.Any]:
    with runtime.handle() as s:
        commands.hello(db, **input.dict())
        return s.dict()


@router.post("/byebye", response_model=runtime.CommandOutput)
def byebye(db: commands.DB=Depends(db)) -> t.Dict[str, t.Any]:
    with runtime.handle() as s:
        commands.byebye(db)
        return s.dict()


def main(app: FastAPI):
    from monogusa.web import cli
    cli.run(app)


app = FastAPI()
app.include_router(router)


if __name__ == '__main__':
    main(app=app)
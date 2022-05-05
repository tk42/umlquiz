export const API_URL = "https://umlquiz-wphmvh57gq-uc.a.run.app"
export const ENDPOINT_TOKEN = "/token"
export const ENDPOINT_USER = "/user"
export const ENDPOINT_QUIZ = "/quiz"

export const TUTORIAL_INITIAL_MARKDOWN = `Animal <|-- Duck
Animal <|-- Fish
Animal <|-- Zebra
Animal : +int age
Animal : +String gender
Animal: +isMammal()
Animal: +mate()
class Duck{
    +String beakColor
    +swim()
    +quack()
}
class Fish{
    -int sizeInFeet
    -canEat()
}
class Zebra{
    +bool is_wild
    +run()
}`;

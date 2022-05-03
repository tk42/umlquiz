import UMLEditor from '../components/umleditor'
import TextareaAutosize from '@mui/material/TextareaAutosize';
import { Typography, Button, Grid } from "@mui/material";
import {useState} from 'react';

const mdMermaid = `classDiagram
Animal <|-- Duck
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
}
`;

function HomePage() {
    const [value, setValue] = useState(mdMermaid);

    const handleChangeMarkdown = (event: React.ChangeEvent<{ value: string }>) => {
        setValue(event.target.value);
    };

    return (
        <div className="bg">
            <Typography variant={"h1"} px={2}>UML Quiz</Typography>
            <div className="glass">
                <p>
                    ある会員制のホテルは15部屋保有しています．予約は一人の会員に対して，1回に1部屋だけ受け付けます．適切なクラス図を描きましょう．
                </p>
                <Grid container spacing={2}>
                    <Grid item xs={6}>
                        <TextareaAutosize
                            value={value}
                            placeholder={"Input classdiagram"}
                            style={{ width: '100%' }}
                            onChange={handleChangeMarkdown}
                        />
                    </Grid>
                    <Grid item xs={6}>
                        <UMLEditor value={value}/>
                    </Grid>
                </Grid>
                <Button>
                    Submit
                </Button>
            </div>
        </div>
    );
}

export default HomePage;
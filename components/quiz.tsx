import { Typography, Grid, Button } from "@mui/material";
import TextField from "@mui/material/TextField";
import UMLPreviewer from "./umlpreview";
import {useState} from 'react';
import { API_URL, ENDPOINT_TOKEN } from './const';

const mdMermaid = `Animal <|-- Duck
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


export default function Quiz({caption, token} : {caption: string, token: string}) {
    // fetch(API_URL+)

    const [value, setValue] = useState(mdMermaid);

    const handleChangeMarkdown = (
      event: React.ChangeEvent<{ value: string }>
    ) => {
      setValue(event.target.value);
    };
  
    return (
        <div className="glass">
          <Typography variant={"h2"}>{caption}</Typography>
          <p>
            ある会員制のホテルは15部屋保有しています．予約は一人の会員に対して，1回に1部屋だけ受け付けます．適切なクラス図を描きましょう．
          </p>
          <Grid container spacing={2}>
            <Grid item xs={6}>
              <TextField
                value={value}
                multiline
                fullWidth
                onChange={handleChangeMarkdown}
              />
            </Grid>
            <Grid item xs={6}>
              <UMLPreviewer value={value} prefix={"classDiagram"} />
            </Grid>
          </Grid>
          <Button>Submit</Button>
        </div>
    )
}
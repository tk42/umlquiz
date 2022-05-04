import {
  useTranslation,
} from "next-export-i18n";

import Link from 'next/link';

import UMLPreviewer from "../components/umlpreview";
import { Typography, Button, Grid } from "@mui/material";
import TextField from "@mui/material/TextField";
import { useState } from "react";

const notation_visibility = `class Visibility
Visibility : +int public_member
Visibility : #int protected_member
Visibility : -int private_member
Visibility : ~int package_member`;

const notation_relationships = `classA0 --|> classB0 : Inheritance
classA1 --* classB1 : Composition
classA2 --o classB2 : Aggregation
classA3 --> classB3 : Association
classA4 ..> classB4 : Dependency
classA5 ..|> classB5 : Realization`;

const notation_sequences = `Alice->>+John: Hello John, how are you?
Alice-->>+John: John, can you hear me?
John-)-Alice: Hi Alice, I can hear you! (async)
John--)-Alice: I feel great! (async)`;

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
}
`;

function HomePage() {
  const { t } = useTranslation();
  const [lang, setLang] = useState('en');

  // const handleChangeLang = (event: SelectChangeEvent) => {
  //   setLang(event.target.value as string);
  // };

  const [value, setValue] = useState(mdMermaid);

  const handleChangeMarkdown = (
    event: React.ChangeEvent<{ value: string }>
  ) => {
    setValue(event.target.value);
  };

  return (
      <div className="bg">
        <Typography variant={"h1"}>UML Quiz</Typography>
        <p>
          <Link href={"/?lang=en"}>English</Link>/<Link href={"/?lang=fr"}>Français</Link>/<Link href={"/?lang=ja"}>日本語</Link>
        </p>
        <p>{t('intro')}</p>
        <div className="glass">
          <Typography variant={"h2"}>{t('notation')}</Typography>
          <Typography variant={"h4"}>{t('caption_classdiagram')}</Typography>
          <Grid container spacing={2}>
            <Grid item xs={6}>
              <TextField
                label="Visibility"
                multiline
                rows={5}
                fullWidth
                defaultValue={notation_visibility}
              />
            </Grid>
            <Grid item xs={6}>
              <UMLPreviewer value={notation_visibility} prefix={"classDiagram"} />
            </Grid>
            <Grid item xs={6}>
              <TextField
                label="Relationships"
                multiline
                rows={6}
                fullWidth
                defaultValue={notation_relationships}
              />
            </Grid>
            <Grid item xs={6}>
              <UMLPreviewer
                value={notation_relationships}
                prefix={"classDiagram"}
              />
            </Grid>
          </Grid>
          <Typography variant={"h4"}>{t('caption_sequencediagram')}</Typography>
          <Grid container spacing={2}>
            <Grid item xs={6}>
              <TextField
                label="Sequence"
                multiline
                rows={4}
                fullWidth
                defaultValue={notation_sequences}
              />
            </Grid>
            <Grid item xs={6}>
              <UMLPreviewer value={notation_sequences} prefix={"sequenceDiagram"} />
            </Grid>
          </Grid>
        </div>
        <div className="glass">
          <Typography variant={"h2"}>Q1</Typography>
          <p>
            ある会員制のホテルは15部屋保有しています．予約は一人の会員に対して，1回に1部屋だけ受け付けます．適切なクラス図を描きましょう．
          </p>
          <Grid container spacing={2}>
            <Grid item xs={6}>
              <TextField
                value={value}
                // placeholder={"Input classdiagram"}
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
      </div>
  );
}

export default HomePage;

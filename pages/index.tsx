import { useTranslation } from "next-export-i18n";

import Link from "next/link";

import UMLPreviewer from "../components/umlpreview";
import Notation from "../components/notation";
import { Typography, Button, Grid } from "@mui/material";
import TextField from "@mui/material/TextField";
import { useState } from "react";

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

function HomePage() {
  const { t } = useTranslation();

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
        <Link href={"/?lang=en"}>English</Link>/
        <Link href={"/?lang=fr"}>Français</Link>/
        <Link href={"/?lang=ja"}>日本語</Link>
      </p>
      <div className="glass">
        {t("intro")}
      </div>
      <Notation />
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

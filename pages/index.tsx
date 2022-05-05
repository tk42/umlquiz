import { useTranslation, useSelectedLanguage, LanguageSwitcher, useLanguageQuery } from "next-export-i18n";
import { useState, useEffect } from "react";

import Head from "next/head";
import Link from "next/link";

import Notation from "../components/notation";
import { Typography } from "@mui/material";
import ToggleButton from '@mui/material/ToggleButton';
import ToggleButtonGroup from '@mui/material/ToggleButtonGroup';
import Tutorial from "../components/tutorial";

function HomePage() {
  const { t } = useTranslation();
  // const [query] = useLanguageQuery();
  // const { lang, setLang } = useSelectedLanguage();
  // const [accessToken, setAccessToken] = useState("");
  // useEffect(() => {
  //   const fetchToken = async () => {
  //     const response = await fetch("/api/token");
  //     const data = await response.json();
  //     setAccessToken(data.access_token);
  //   };
  //   fetchToken();
  // }, []);

  return (
    <>
      <Head>
        <title>UML Quiz</title>
      </Head>
      <div className="bg">
        <Typography variant={"h1"}>UML Quiz</Typography>
        <ToggleButtonGroup
          // value={lang}
          exclusive
          // onChange={(e, value)=>{
          //   setLang(value)
          // }}
        >
          <Link href="/?lang=en">
            <ToggleButton value="en">
              English
            </ToggleButton>
          </Link>
          <Link href="/?lang=fr">
            <ToggleButton value="fr">
              Français
            </ToggleButton>
          </Link>
          <Link href="/?lang=ja">
            <ToggleButton value="ja">
              日本語
            </ToggleButton>
          </Link>
        </ToggleButtonGroup>
        <div className="glass">{t("intro")}</div>
        <Notation />
        <Tutorial />
      </div>
    </>
  );
}

export default HomePage;

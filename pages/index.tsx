import { useTranslation } from "next-export-i18n";
import { useState, useEffect } from 'react';

import Link from "next/link";
import Head from "next/head";

import Notation from "../components/notation";
import Quiz from "../components/quiz";
import { getToken } from "../components/token";
import { Typography } from "@mui/material";


function HomePage() {
  const { t } = useTranslation();
  const [accessToken, setAccessToken] = useState('');

  useEffect(()=>{
    getToken().then((token: string)=>{
      setAccessToken(token)
    })
  }, [])


  return (
    <>
      <Head>
        <title>UML Quiz</title>
      </Head>
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
        <Quiz caption={"Q1"} token={accessToken} />
      </div>
    </>
  );
}

export default HomePage;

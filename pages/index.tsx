import { useTranslation } from "next-export-i18n";

import Link from "next/link";
import Head from "next/head";

import Notation from "../components/notation";
import Quiz from "../components/quiz";
import { getToken } from "../components/token";
import { Typography } from "@mui/material";


async function HomePage() {
  const { t } = useTranslation();

  const access_token: string = await getToken();

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
        <Quiz caption={"Q1"} token={access_token} />
      </div>
    </>
  );
}

export default HomePage;

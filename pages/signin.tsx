import { useTranslation } from "next-export-i18n";
import { TUTORIAL_INITIAL_MARKDOWN } from "../components/const";
import QuizViewer from "../components/quizviewer";
import { Typography, Grid, Button, Box, Modal, TextField } from "@mui/material";

export default function Signin() {
    const { t } = useTranslation();
    return (
        <>
            <Button>Signin</Button>
            {/* <Quiz language={lang} quiz_id={"tutorial"}/> // required access_token */}
        </>        
    );
}
import { useTranslation } from "next-export-i18n";
import { TUTORIAL_INITIAL_MARKDOWN } from "../components/const";
import QuizViewer from "../components/quizviewer";

export default function Tutorial() {
    const { t } = useTranslation();
    return (
        <>
            <QuizViewer
                caption={"Tutorial"}
                diagram_type={"classDiagram"}
                text={t("tutorial")}
                diagram={TUTORIAL_INITIAL_MARKDOWN}
                linkto={"/signin"}
            />
            {/* <Quiz language={lang} quiz_id={"tutorial"}/> // required access_token */}
        </>        
    );
}
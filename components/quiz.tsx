import QuizViewer from './quizviewer';
import { FetchQuizProps, Quiz as QuizModel } from './model'
import { ENDPOINT_QUIZ, TUTORIAL_INITIAL_MARKDOWN } from './const';
import {useEffect, useState} from 'react';

export default function Quiz(props: FetchQuizProps) {
    const [data, setData] = useState({} as QuizModel);
    useEffect(() => {
        const fetchData = async () => {
            const resp = await fetch("./api/" + props.language + "/" + props.user_id + ENDPOINT_QUIZ);
            const data = await resp.json();
            setData(data);
        }
        fetchData();
    }, [props]);
    
    return (
        <QuizViewer
            caption={data.title}
            diagram_type={data.diagram_type}
            text={data.text}
            diagram={TUTORIAL_INITIAL_MARKDOWN}
            linkto={props.linkto}
        />
    )
}
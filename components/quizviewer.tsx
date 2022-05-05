import { useTranslation } from "next-export-i18n";
import { Typography, Grid, Button, Box, Modal, TextField } from "@mui/material";
import UMLPreviewer from "./umlpreview";
import { useState } from "react";
import ReactDiffViewer from 'react-diff-viewer';
import { TUTORIAL_INITIAL_MARKDOWN } from "./const";

const style = {
  position: 'absolute' as 'absolute',
  top: '50%',
  left: '50%',
  transform: 'translate(-50%, -50%)',
  width: '80%',
  bgcolor: 'background.paper',
  border: '2px solid #000',
  boxShadow: 24,
  p: 4,
};

export default function QuizViewer({
  caption,
  diagram_type,
  text,
  diagram, // answer
  linkto,
}: {
  caption: string;
  diagram_type: string;
  text: string;
  diagram: string;
  linkto: string;
}) {
  const { t } = useTranslation();

  const [open, setOpen] = useState(false);

  const [markdown, setMarkdown] = useState(TUTORIAL_INITIAL_MARKDOWN);

  const handleChangeMarkdown = (
    event: React.ChangeEvent<{ value: string }>
  ) => {
    setMarkdown(event.target.value);
  };

  return (
    <div className="glass">
      <Typography variant={"h2"}>{caption}</Typography>
      <p>
        {text}
      </p>
      <Grid container spacing={2}>
        <Grid item xs={6}>
          <TextField
            value={markdown}
            multiline
            fullWidth
            onChange={handleChangeMarkdown}
          />
        </Grid>
        <Grid item xs={6}>
          <UMLPreviewer value={markdown} prefix={diagram_type} />
        </Grid>
      </Grid>
      <Button onClick={()=>setOpen(true)}>✓ CHECK</Button>
      <Modal
        open={open}
        onClose={(e) => setOpen(false)}
      >
        <Box sx={style}>
          {
            (markdown === diagram) ? 
                <Typography sx={{ mt: 2 }}>
                  {t('correct_message')}
                </Typography>
              :
              <>
                <Typography variant="h6" component="h2">
                  {t('wrong_message')}
                </Typography>
                <ReactDiffViewer oldValue={markdown} newValue={diagram} splitView={true} />
              </>
          }
          <Button href={linkto}>{t('proceed_quiz')}</Button>
        </Box>
      </Modal>
    </div>
  );
}

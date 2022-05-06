import { useTranslation } from "next-export-i18n";
import { Typography, Grid } from "@mui/material";
import TextField from "@mui/material/TextField";
import UMLPreviewer from "./umlpreview";

const notation_annotations = `class Interface
<<interface>> Interface
class Abstract
<<abstract>> Abstract
class Enum
<<enum>> Enum`

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

const notation_sequences = `Alice->>+Bob: Hello Bob, how are you?
Alice-->>+Bob: Bob, can you hear me?
Bob-)-Alice: Hi Alice, I can hear you! (async)
Bob--)-Alice: I feel great! (async)`;


export default function Notation() {
    const { t } = useTranslation();
    return (
      <div className="glass">
        <Typography variant={"h2"}>{t("notation")}</Typography>
        <Typography variant={"h3"}>{t("caption_classdiagram")}</Typography>
        <Grid container spacing={2}>
          <Grid item xs={6}>
            <TextField
              label="Annotations"
              multiline
              rows={6}
              fullWidth
              defaultValue={notation_annotations}
            />
          </Grid>
          <Grid item xs={6}>
            <UMLPreviewer value={notation_annotations} prefix={"classDiagram"} />
          </Grid>
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
        <Typography variant={"h3"}>{t("caption_sequencediagram")}</Typography>
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
            <UMLPreviewer
              value={notation_sequences}
              prefix={"sequenceDiagram"}
            />
          </Grid>
        </Grid>
      </div>
    );
  };
  
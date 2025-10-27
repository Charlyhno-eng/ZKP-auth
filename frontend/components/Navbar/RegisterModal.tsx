import Box from "@mui/material/Box";
import Button from "@mui/material/Button";
import Modal from "@mui/material/Modal";
import TextField from "@mui/material/TextField";
import Typography from "@mui/material/Typography";

type RegisterModalProps = {
  open: boolean;
  onClose: () => void;
};

const style = {
  position: "absolute" as const,
  top: "50%",
  left: "50%",
  transform: "translate(-50%, -50%)",
  width: 400,
  bgcolor: "#242526",
  color: "#fff",
  border: "none",
  borderRadius: 2,
  boxShadow: 24,
  p: 4,
};

export default function RegisterModal({ open, onClose }: RegisterModalProps) {
  return (
    <Modal open={open} onClose={onClose}>
      <Box sx={style}>
        <Typography variant="h6" component="h2" sx={{ mb: 2 }}>
          Créer un compte
        </Typography>
        <TextField
          label="Nom d'utilisateur"
          variant="filled"
          fullWidth
          InputProps={{ sx: { background: "#35363a", color: "#fff" } }}
          InputLabelProps={{
            sx: { color: "#bbb" },
          }}
          sx={{ mb: 2 }}
        />
        <Button
          variant="contained"
          fullWidth
          color="primary"
          onClick={onClose}
          sx={{ background: "#1976d2", color: "#fff" }}
        >
          Créer
        </Button>
      </Box>
    </Modal>
  );
}
